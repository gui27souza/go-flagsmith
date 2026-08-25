package router

import (
	"context"
	"goflagsmith/internal/service/flags"
	"goflagsmith/internal/state"
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

// TODO - implement Route logic
func (e *Engine) Route(
	ctx context.Context, req DecideReq,
) (DecideRes, error) {

	// TODO - Define res Target and Reason
	// Define Target based on userID and flag,
	// by consequence, define Reason

	snapshot := e.s.Snapshot()

	return DecideRes{
		Target: "",
		Reason: "",
		Telemetry: TelemetryData{
			EvaluatedAt:   time.Now(),
			CacheHydrated: snapshot.Features,
		},
	}, nil
}
