package routes

import (
	"sumoweb/server/internal/handlers"
	matchHandlers "sumoweb/server/internal/handlers/match"
	rikishiHandlers "sumoweb/server/internal/handlers/rikishi"
	wsHandlers "sumoweb/server/internal/handlers/websocket"
	"sumoweb/server/internal/middleware"
	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {
	middleware.SetupLogger(router)
	middleware.SetupRecovery(router)
	middleware.SetupCORS(router)

	router.GET("/health", handlers.HealthCheck)

	api := router.Group("api/v1")
	{
		api.GET("/ws/:id", wsHandlers.HandleWebSocket)
		api.GET("/matches/:division/:day", matchHandlers.HandleGetMatches)
		api.GET("/rikishi/:id", rikishiHandlers.HandleGetRikishi)
		api.GET("/rikishi/:id/:lang/image.png", rikishiHandlers.HandleGetRikishiImage)
	}
}
