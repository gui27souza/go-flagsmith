package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"goflagsmith/internal/service/flags"
	"goflagsmith/internal/state"
)

type AppHandler struct {
	state    *state.State
	flagsSvc flags.Service
}

func NewAppHandler(s *state.State, fSvc flags.Service) *AppHandler {
	return &AppHandler{
		state:    s,
		flagsSvc: fSvc,
	}
}

func (h *AppHandler) Readyz(c *gin.Context) {

	s := h.state.Snapshot()

	if !s.Clients || !s.Features {
		c.JSON(http.StatusServiceUnavailable, s)
		return
	}

	c.JSON(http.StatusOK, s)
}
