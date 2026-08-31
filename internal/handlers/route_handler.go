package handlers

import (
	"goflagsmith/internal/service/router"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RouteHandler struct {
	eng *router.Engine
}

func NewRouteHandler(engine *router.Engine) *RouteHandler {
	return &RouteHandler{eng: engine}
}

func (rh *RouteHandler) Handle(c *gin.Context) {

	var req router.DecideReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": "Invalid request payload: " + err.Error()},
		)
		return
	}

	res := rh.eng.Route(c.Request.Context(), req)

	c.JSON(http.StatusOK, res)
}
