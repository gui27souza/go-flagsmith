package handlers_test

import (
	"context"
	"encoding/json"
	"goflagsmith/internal/handlers"
	"goflagsmith/internal/state"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

type mockFlagsService struct{}

func (m *mockFlagsService) IsFeatureEnabled(ctx context.Context, featureName string) bool {
	return true
}
func (m *mockFlagsService) GetJSONConfig(ctx context.Context, configName string) (string, error) {
	return "", nil
}
func (m *mockFlagsService) MonitorFlagsReady(
	ctx context.Context,
	appState *state.State,
	interval time.Duration,
) {}

func setupRouter(h *handlers.AppHandler) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/readyz", h.Readyz)
	return r
}

func TestReadyz_NotReady(t *testing.T) {

	appState := state.NewState()
	mockFlags := &mockFlagsService{}

	h := handlers.NewAppHandler(appState, mockFlags)

	router := setupRouter(h)

	// Simulates a GET request to /readyz
	req, _ := http.NewRequest(http.MethodGet, "/readyz", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	// State was created, not ready yet
	if resp.Code != http.StatusServiceUnavailable {
		t.Errorf("expected status 503, got %d", resp.Code)
	}

	var body state.ReadinessStatus
	if err := json.Unmarshal(resp.Body.Bytes(), &body); err != nil {
		t.Fatalf("failed to decode response JSON: %v", err)
	}

	if body.Clients || body.Features {
		t.Errorf("expected clients and features to be false in payload")
	}
}

func TestReadyz_Ready(t *testing.T) {

	appState := state.NewState()

	// Sets State to be fully ready
	appState.SetClientsReady()
	appState.SetFeaturesReady()

	mockFlags := &mockFlagsService{}
	h := handlers.NewAppHandler(appState, mockFlags)
	router := setupRouter(h)

	// Simulates a GET request to /readyz
	req, _ := http.NewRequest(http.MethodGet, "/readyz", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	// State fully ready
	if resp.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.Code)
	}
}
