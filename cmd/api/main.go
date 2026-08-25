package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"goflagsmith/internal/handlers"
	"goflagsmith/internal/service/flags"
	"goflagsmith/internal/service/router"
	"goflagsmith/internal/state"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	appState := state.NewState()

	flagsSvc := flags.NewClient(ctx, os.Getenv("FLAGSMITH_API_KEY"))
	appState.SetClientsReady()
	flagsSvc.MonitorFlagsReady(
		ctx, appState, 2*time.Second,
	)

	h := handlers.NewAppHandler(appState, flagsSvc)

	engine := router.NewEngine(flagsSvc)
	rh := handlers.NewRouteHandler(engine)

	router := gin.Default()

	router.GET("/healthz", func(c *gin.Context) {
		c.String(http.StatusOK, "server is running")
	})

	router.GET("/readyz", h.Readyz)

	router.POST("/decide", rh.Handle)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("FATAL: HTTP server terminated with error: %v", err)
	}
}
