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
	"goflagsmith/internal/util/hash"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	appState := state.NewState()

	flagsSvc, err := flags.NewClient(ctx, os.Getenv("FLAGSMITH_API_KEY"))
	if err != nil {
		// TODO - gracefull shutdown
	}
	appState.SetClientsReady()

	interval := 2 * time.Second
	flagsSvc.MonitorFlagsReady(
		ctx, appState, interval,
	)
	flagsSvc.StartRulesSync(ctx, interval)

	h := handlers.NewAppHandler(appState, flagsSvc)

	engine := router.NewEngine(
		flagsSvc, appState, hash.NormalizedHash, time.Now,
	)
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
