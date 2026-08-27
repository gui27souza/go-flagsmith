package flags

import (
	"context"
	"goflagsmith/internal/state"
	"time"
)

type MockReader struct {
	FeatureEnabled bool
	JSONConfig     string
	ErrJSONConfig  error
}

func (m *MockReader) IsFeatureEnabled(ctx context.Context, featureName string) bool {
	return m.FeatureEnabled
}
func (m *MockReader) GetJSONConfig(ctx context.Context, configName string) (string, error) {
	return m.JSONConfig, m.ErrJSONConfig
}

func NewMockReader(
	featEnabled bool, jsonConfig string, errJson error,
) *MockReader {
	return &MockReader{
		FeatureEnabled: featEnabled,
		JSONConfig:     jsonConfig,
		ErrJSONConfig:  errJson,
	}
}

type MockService struct {
	*MockReader
}

func NewMockService(reader *MockReader) *MockService {
	return &MockService{MockReader: reader}
}

func (m *MockService) MonitorFlagsReady(
	ctx context.Context, appState *state.State, interval time.Duration,
) {
}
