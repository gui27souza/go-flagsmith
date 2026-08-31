package router_test

import (
	"goflagsmith/internal/service/router"
	"goflagsmith/internal/state"
	"goflagsmith/internal/testutil"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRoute_Success(t *testing.T) {

	fr := testutil.NewMockReader(
		true,
		`{"canary_percentage": 50, "allowed_countries": ["BR"]}`,
		nil,
	)

	s := state.NewState()
	s.SetClientsReady()
	s.SetFeaturesReady()

	mockBcktCalc := func(expression string) int {
		return 0
	}

	fixedTime := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)

	mockEngine := router.NewEngine(
		fr, s, mockBcktCalc, func() time.Time { return fixedTime },
	)

	req := router.DecideReq{
		UserID:     "user_1234",
		Country:    "br",
		AppVersion: "1.27",
	}

	res := mockEngine.Route(t.Context(), req)

	expectedRes := router.DecideRes{
		Target: "v2",
		Reason: "canary sorting rules",
		Telemetry: router.TelemetryData{
			EvaluatedAt:   res.Telemetry.EvaluatedAt,
			CacheHydrated: s.Snapshot().Features,
		},
	}

	assert.Equal(t, expectedRes, res)
}
