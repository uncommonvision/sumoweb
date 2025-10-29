package routes

import (
	"github.com/gin-gonic/gin"
	"sumoweb/server/internal/handlers"
	matchHandlers "sumoweb/server/internal/handlers/match"
	rikishiHandlers "sumoweb/server/internal/handlers/rikishi"
	wsHandlers "sumoweb/server/internal/handlers/websocket"
	"sumoweb/server/internal/middleware"
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

		admin := api.Group("/admin")
		{
			admin.POST("/enqueue/matches/:division/:day", matchHandlers.HandleEnqueueFetchMatches)
			admin.POST("/enqueue/rikishi/profiles", rikishiHandlers.HandleEnqueueRikishiProfiles)
		}
	}
}
