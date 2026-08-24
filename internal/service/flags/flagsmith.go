package flags

import (
	"context"
	"goflagsmith/internal/state"
	"log"
	"time"

	"github.com/Flagsmith/flagsmith-go-client/v4"
)

type flagsmithSDK interface {
	GetEnvironmentFlags(ctx context.Context) (flagsmith.Flags, error)
}

type Service interface {
	IsFeatureEnabled(ctx context.Context, featureName string) bool
}

type Client struct {
	sdk flagsmithSDK
}

func NewClient(ctx context.Context, apiKey string) *Client {
	if apiKey == "" {
		log.Fatalf("FATAL: FLAGSMITH_API_KEY environment variable is required")
	}

	sdk := flagsmith.NewClient(
		apiKey,
		flagsmith.WithLocalEvaluation(ctx),
		flagsmith.WithEnvironmentRefreshInterval(60*time.Second),
		flagsmith.WithRequestTimeout(5*time.Second),
	)

	return &Client{sdk}
}

// Testable Constructor
func NewClientWithSDK(sdk flagsmithSDK) *Client {
	return &Client{sdk: sdk}
}

func (c *Client) MonitorFlagsReady(
	ctx context.Context, appState *state.State, interval time.Duration,
) {

	go func() {

		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				if appState.Snapshot().Features {
					return
				}

				ctxCheck, cancel := context.WithTimeout(ctx, 1*time.Second)
				_, err := c.sdk.GetEnvironmentFlags(ctxCheck)
				cancel()

				if err == nil {
					appState.SetFeaturesReady()
					log.Println("INFO: Flags cache ready, service is now fully ready for traffic")
					return
				}
			}
		}
	}()
}

func (c *Client) IsFeatureEnabled(ctx context.Context, featureName string) bool {

	flags, err := c.sdk.GetEnvironmentFlags(ctx)
	if err != nil {
		return false
	}

	enabled, err := flags.IsFeatureEnabled(featureName)
	if err != nil {
		return false
	}
	return enabled
}
