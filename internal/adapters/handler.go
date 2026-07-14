package adapters

import (
	"main/internal/ports"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Logger ports.ZapLogger
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		api.GET("/", func(ctx *gin.Context) { ctx.JSON(200, "ok") })
	}

	return router
}
