package router_test

import (
	"errors"
	"goflagsmith/internal/domain"
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

	defaultReq := domain.UserContext{
		UserID:     "user_1234",
		Country:    "br",
		AppVersion: "1.27",
	}
	defaultJSONConfig := `{"canary_percentage": 50, "allowed_countries": ["BR"]}`
	defaultBc := func(str string) int { return 0 }
	defaultTelemetry := domain.TelemetryData{
		EvaluatedAt:   fixedTime,
		CacheHydrated: s.Snapshot().Features,
	}

	tests := []struct {
		name          string
		req           domain.UserContext
		featEnable    bool
		jsonConfig    string
		errJSONConfig error
		bc            router.BucketCalculator
		wantRes       domain.RouteDecision
	}{
		{
			name:          "succesfull routing to v2",
			req:           defaultReq,
			featEnable:    true,
			jsonConfig:    defaultJSONConfig,
			errJSONConfig: nil,
			bc:            defaultBc,
			wantRes: domain.RouteDecision{
				Target:    "v2",
				Reason:    "canary sorting rules",
				Telemetry: defaultTelemetry,
			},
		},
		{
			name:          "succesful routing to v1 - bucket calculation",
			req:           defaultReq,
			featEnable:    true,
			jsonConfig:    defaultJSONConfig,
			errJSONConfig: nil,
			bc:            func(str string) int { return 60 },
			wantRes: domain.RouteDecision{
				Target:    "v1",
				Reason:    "canary sorting rules",
				Telemetry: defaultTelemetry,
			},
		},
		{
			name:          "feat disabled",
			req:           defaultReq,
			featEnable:    false,
			jsonConfig:    "",
			errJSONConfig: nil,
			bc:            defaultBc,
			wantRes: domain.RouteDecision{
				Target:    "v1",
				Reason:    "v2 routing not enabled",
				Telemetry: defaultTelemetry,
			},
		},
		{
			name:          "error fetching canary rules",
			req:           defaultReq,
			featEnable:    true,
			jsonConfig:    "",
			errJSONConfig: errors.New("something went wrong"),
			bc:            defaultBc,
			wantRes: domain.RouteDecision{
				Target:    "v1",
				Reason:    "internal error on canary rules fetching",
				Telemetry: defaultTelemetry,
			},
		},
		{
			name:          "error unmarshalling canary rules json",
			req:           defaultReq,
			featEnable:    true,
			jsonConfig:    `{"broken}`,
			errJSONConfig: nil,
			bc:            defaultBc,
			wantRes: domain.RouteDecision{
				Target:    "v1",
				Reason:    "internal error on canary rules parsing",
				Telemetry: defaultTelemetry,
			},
		},
		{
			name: "v2 unavailable for user country",
			req: domain.UserContext{
				UserID:     "user_1234",
				Country:    "Us",
				AppVersion: "1.27",
			},
			featEnable:    true,
			jsonConfig:    defaultJSONConfig,
			errJSONConfig: nil,
			bc:            defaultBc,
			wantRes: domain.RouteDecision{
				Target:    "v1",
				Reason:    "unavailable for user country",
				Telemetry: defaultTelemetry,
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
