package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Flagsmith/flagsmith-go-client/v4"
	"github.com/gin-gonic/gin"

	"goflagsmith/internal/handlers"
	"goflagsmith/internal/state"
)

func main() {

	appState := state.NewState()

	fsClient := InitFsClient()

	appState.SetClientsReady()

	h := handlers.NewAppHandler(appState, fsClient)

	router := gin.Default()

	router.GET("/healthz", func(c *gin.Context) {
		c.String(http.StatusOK, "server is running")
	})

	router.GET("/readyz", h.Readyz)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("FATAL: HTTP server terminated with error: %v", err)
	}
}

func InitFsClient() *flagsmith.Client {
	fsAPIKey := os.Getenv("FLAGSMITH_API_KEY")
	if fsAPIKey == "" {
		log.Fatalf("FATAL: FLAGSMITH_API_KEY environment variable is required")
	}

	return flagsmith.NewClient(fsAPIKey)
}
