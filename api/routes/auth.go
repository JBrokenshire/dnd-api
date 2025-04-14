package routes

import (
	"dnd-api/api"
	"dnd-api/api/handlers"
	"dnd-api/config"
	mw "dnd-api/pkg/middleware"
	"github.com/labstack/echo/v4/middleware"
)

func authRoutes(server *api.Server) {
	authHandler := handlers.NewAuthHandler(server)

	auth := server.Echo.Group("/auth")
	auth.Use(mw.Security(false))

	limiterStore := middleware.NewRateLimiterMemoryStore(config.Get().LoginRateLimit)

	auth.POST("/login", authHandler.Login, middleware.RateLimiter(limiterStore))

}
