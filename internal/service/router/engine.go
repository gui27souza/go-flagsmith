package router

import (
	"context"
	"encoding/json"
	"goflagsmith/internal/service/flags"
	"goflagsmith/internal/state"
	"strings"
	"time"
)

type Engine struct {
	fr flags.Reader
	s  *state.State
}

func NewEngine(fr flags.Reader, s *state.State) *Engine {
	return &Engine{
		fr: fr, s: s,
	}
}

type DecideReq struct {
	UserID     string `json:"user_id"`
	Country    string `json:"country"`
	AppVersion string `json:"app_version"`
}

type TelemetryData struct {
	EvaluatedAt   time.Time `json:"evaluated_at"`
	CacheHydrated bool      `json:"cache_hydrated"`
}

type DecideRes struct {
	Target    string        `json:"target"`
	Reason    string        `json:"reason"`
	Telemetry TelemetryData `json:"telemetry"`
}

type CanaryRoutingRules struct {
	Percentage int      `json:"canary_percentage"`
	Countries  []string `json:"allowed_countries"`
}

// TODO - implement Route logic
func (e *Engine) Route(
	ctx context.Context, req DecideReq,
) (DecideRes, error) {

	if !e.fr.IsFeatureEnabled(ctx, "enable_v2_routing") {
		return e.deniedRouting("v2 routing not enabled"), nil
	}

	rulesJSON, err := e.fr.GetJSONConfig(ctx, "canary_routing_rules")
	if err != nil {
		return e.deniedRouting("internal error on canary rules fetching"), nil
	}

	var rules CanaryRoutingRules
	if err := json.Unmarshal([]byte(rulesJSON), &rules); err != nil {
		return e.deniedRouting("internal error on canary rules parsing"), nil
	}

	var allowCountry bool
	for _, country := range rules.Countries {
		if strings.EqualFold(country, req.Country) {
			allowCountry = true
			break
		}
	}
	if !allowCountry {
		return e.deniedRouting("unavailable for user country"), nil
	}

	// TODO - Define res Target

	var target string
	return DecideRes{
		Target:    target,
		Reason:    "canary sorting rules",
		Telemetry: e.getTelemetryData(),
	}, nil
}

func (e *Engine) getTelemetryData() TelemetryData {
	snapshot := e.s.Snapshot()
	return TelemetryData{
		EvaluatedAt:   time.Now(),
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
