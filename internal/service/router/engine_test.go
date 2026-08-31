package router_test

import (
	"errors"
	"goflagsmith/internal/service/router"
	"goflagsmith/internal/state"
	"goflagsmith/internal/testutil"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRoute(t *testing.T) {

	s := state.NewState()
	s.SetClientsReady()
	s.SetFeaturesReady()

	fixedTime := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	mockNow := func() time.Time { return fixedTime }

	defaultReq := router.DecideReq{
		UserID:     "user_1234",
		Country:    "br",
		AppVersion: "1.27",
	}
	defaultJSONConfig := `{"canary_percentage": 50, "allowed_countries": ["BR"]}`
	defaultBc := func(str string) int { return 0 }

	tests := []struct {
		name          string
		req           router.DecideReq
		featEnable    bool
		jsonConfig    string
		errJSONConfig error
		bc            router.BucketCalculator
		wantRes       router.DecideRes
	}{
		{
			name:          "succesfull routing to v2",
			req:           defaultReq,
			featEnable:    true,
			jsonConfig:    defaultJSONConfig,
			errJSONConfig: nil,
			bc:            defaultBc,
			wantRes: router.DecideRes{
				Target: "v2",
				Reason: "canary sorting rules",
				Telemetry: router.TelemetryData{
					EvaluatedAt:   fixedTime,
					CacheHydrated: s.Snapshot().Features,
				},
			},
		},
		{
			name:          "succesful routing to v1 - bucket calculation",
			req:           defaultReq,
			featEnable:    true,
			jsonConfig:    defaultJSONConfig,
			errJSONConfig: nil,
			bc:            func(str string) int { return 60 },
			wantRes: router.DecideRes{
				Target: "v1",
				Reason: "canary sorting rules",
				Telemetry: router.TelemetryData{
					EvaluatedAt:   fixedTime,
					CacheHydrated: s.Snapshot().Features,
				},
			},
		},
		{
			name:          "feat disabled",
			req:           defaultReq,
			featEnable:    false,
			jsonConfig:    "",
			errJSONConfig: nil,
			bc:            defaultBc,
			wantRes: router.DecideRes{
				Target: "v1",
				Reason: "v2 routing not enabled",
				Telemetry: router.TelemetryData{
					EvaluatedAt:   fixedTime,
					CacheHydrated: s.Snapshot().Features,
				},
			},
		},
		{
			name:          "error fetching canary rules",
			req:           defaultReq,
			featEnable:    true,
			jsonConfig:    "",
			errJSONConfig: errors.New("something went wrong"),
			bc:            defaultBc,
			wantRes: router.DecideRes{
				Target: "v1",
				Reason: "internal error on canary rules fetching",
				Telemetry: router.TelemetryData{
					EvaluatedAt:   fixedTime,
					CacheHydrated: s.Snapshot().Features,
				},
			},
		},
		{
			name:          "error unmarshalling canary rules json",
			req:           defaultReq,
			featEnable:    true,
			jsonConfig:    `{"broken}`,
			errJSONConfig: nil,
			bc:            defaultBc,
			wantRes: router.DecideRes{
				Target: "v1",
				Reason: "internal error on canary rules parsing",
				Telemetry: router.TelemetryData{
					EvaluatedAt:   fixedTime,
					CacheHydrated: s.Snapshot().Features,
				},
			},
		},
		{
			name: "v2 unavailable for user country",
			req: router.DecideReq{
				UserID:     "user_1234",
				Country:    "Us",
				AppVersion: "1.27",
			},
			featEnable:    true,
			jsonConfig:    defaultJSONConfig,
			errJSONConfig: nil,
			bc:            defaultBc,
			wantRes: router.DecideRes{
				Target: "v1",
				Reason: "unavailable for user country",
				Telemetry: router.TelemetryData{
					EvaluatedAt:   fixedTime,
					CacheHydrated: s.Snapshot().Features,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			fr := testutil.NewMockReader(
				tt.featEnable, tt.jsonConfig, tt.errJSONConfig,
			)

			mockEngine := router.NewEngine(
				fr, s, tt.bc, mockNow,
			)

			res := mockEngine.Route(t.Context(), tt.req)

			assert.Equal(t, tt.wantRes, res)
		})
	}
}
