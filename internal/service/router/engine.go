package router

import (
	"context"
	"encoding/json"
	"goflagsmith/internal/service/flags"
	"goflagsmith/internal/state"
	"strings"
	"time"
)

type BucketCalculator func(expression string) int
type Now func() time.Time

// Engine is responsible for making traffic routing decisions
// based on feature flags and dynamic canary rules.
type Engine struct {
	fr  flags.Reader
	s   *state.State
	bc  BucketCalculator
	now Now
}

func NewEngine(
	fr flags.Reader, s *state.State, bc BucketCalculator, now Now,
) *Engine {
	return &Engine{
		fr: fr, s: s, bc: bc, now: now,
	}
}

// DecideReq defines the client context data required
// to make a routing decision.
type DecideReq struct {
	UserID     string `json:"user_id"`
	Country    string `json:"country"`
	AppVersion string `json:"app_version"`
}

// TelemetryData encapsulates telemetry metadata about the routing evaluation.
type TelemetryData struct {
	EvaluatedAt   time.Time `json:"evaluated_at"`
	CacheHydrated bool      `json:"cache_hydrated"`
}

// DecideRes represents the response payload containing the final routing decision.
type DecideRes struct {
	Target    string        `json:"target"`
	Reason    string        `json:"reason"`
	Telemetry TelemetryData `json:"telemetry"`
}

// CanaryRoutingRules defines the dynamic configuration structure
// retrieved via Remote Config.
type CanaryRoutingRules struct {
	Percentage int      `json:"canary_percentage"`
	Countries  []string `json:"allowed_countries"`
}

// Route evaluates the user context and deterministically decides whether the
// traffic should be routed to the stable version ("v1") or the canary version ("v2").
func (e *Engine) Route(
	ctx context.Context, req DecideReq,
) DecideRes {

	if !e.fr.IsFeatureEnabled(ctx, "enable_v2_routing") {
		return e.deniedRouting("v2 routing not enabled")
	}

	rulesJSON, err := e.fr.GetJSONConfig(ctx, "canary_routing_rules")
	if err != nil {
		return e.deniedRouting("internal error on canary rules fetching")
	}

	// PERF - parse rules once, somewhere else
	var rules CanaryRoutingRules
	if err := json.Unmarshal([]byte(rulesJSON), &rules); err != nil {
		return e.deniedRouting("internal error on canary rules parsing")
	}

	var allowCountry bool
	for _, country := range rules.Countries {
		if strings.EqualFold(country, req.Country) {
			allowCountry = true
			break
		}
	}
	if !allowCountry {
		return e.deniedRouting("unavailable for user country")
	}

	usrBucket := e.bc(req.UserID)

	target := "v1"
	if usrBucket < rules.Percentage {
		target = "v2"
	}

	return DecideRes{
		Target:    target,
		Reason:    "canary sorting rules",
		Telemetry: e.getTelemetryData(),
	}
}

func (e *Engine) getTelemetryData() TelemetryData {
	snapshot := e.s.Snapshot()
	return TelemetryData{
		EvaluatedAt:   e.now(),
		CacheHydrated: snapshot.Features,
	}
}

func (e *Engine) deniedRouting(msg string) DecideRes {
	return DecideRes{
		Target:    "v1",
		Reason:    msg,
		Telemetry: e.getTelemetryData(),
	}
}
