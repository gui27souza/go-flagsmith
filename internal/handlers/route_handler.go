package handlers

import (
	"goflagsmith/internal/service/router"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RouteHandler struct {
	eng *router.Engine
}

func NewRouteHandler(Engine *router.Engine) *RouteHandler {
	return &RouteHandler{eng: Engine}
}

func (rh *RouteHandler) Handle(c *gin.Context) {

	var req router.DecideReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": "Invalid request payload: " + err.Error()},
		)
		return
	}

	res, err := rh.eng.Route(c.Request.Context(), req)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "Failed to evaluate route"},
		)
		return
	}

	c.JSON(http.StatusOK, res)
}
