package handlers_test

import (
	"bytes"
	"encoding/json"
	"goflagsmith/internal/handlers"
	"goflagsmith/internal/service/router"
	"goflagsmith/internal/state"
	"goflagsmith/internal/testutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func setupMockRouterDecide(rh *handlers.RouteHandler) *gin.Engine {
	return testutil.MockRouter(
		http.MethodPost, "/decide", rh.Handle,
	)
}

func TestRouteHandler_Success(t *testing.T) {

	state := state.NewState()

	mockReader := testutil.NewMockReader(
		true,
		`{"canary_percentage": 50, "allowed_countries": ["BR"]}`,
		nil,
	)

	rh := handlers.NewRouteHandler(
		router.NewEngine(mockReader, state),
	)

	ginRouter := setupMockRouterDecide(rh)

	jsonBody := []byte(`{
		"user_id": "user_123",
		"country": "BR",
		"app_version": "1.27"
	}`)
	req, _ := http.NewRequest(
		http.MethodPost, "/decide", bytes.NewBuffer(jsonBody),
	)

	resp := httptest.NewRecorder()
	ginRouter.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("expected %d, got %d", http.StatusOK, resp.Code)
	}

	var res router.DecideRes
	if err := json.Unmarshal(resp.Body.Bytes(), &res); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	if res.Target == "" || res.Reason == "" {
		t.Errorf("expected target and reason to be populated, got %+v", res)
	}

	if res.Telemetry.EvaluatedAt.IsZero() {
		t.Errorf("expected telemetry evaluated_at to be populated")
	}
}
