package handlers

import (
	"net/http"

	"github.com/Flagsmith/flagsmith-go-client/v4"
	"github.com/gin-gonic/gin"

	"goflagsmith/internal/state"
)

type AppHandler struct {
	state       *state.State
	flagsClient *flagsmith.Client
}

func NewAppHandler(s *state.State, fs *flagsmith.Client) *AppHandler {
	return &AppHandler{
		state:       s,
		flagsClient: fs,
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
