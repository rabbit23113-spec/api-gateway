package adapters

import (
	"fmt"
	"main/internal/ports"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Logger  ports.ZapLogger
	Service *Service
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		api.GET("/", func(ctx *gin.Context) { ctx.JSON(200, "ok") })
		api.GET("/users", func(ctx *gin.Context) {
			token := ctx.GetHeader("Authorization")
			if err := h.Service.ValidateJWT(token); err != nil {
				ctx.JSON(http.StatusUnauthorized, gin.H{"Error": err.Error()})
				return
			}
			ctx.Redirect(http.StatusPermanentRedirect, fmt.Sprintf("http://localhost:11090/%s", ctx.Request.URL))
		})
	}

	return router
}
