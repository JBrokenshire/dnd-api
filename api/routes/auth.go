package routes

import (
	s "dnd-api/api"
	"dnd-api/api/handlers"
	mw "dnd-api/pkg/middleware"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
	"time"
)

func authRoutes(server *s.Server) {
	authHandler := handlers.NewAuthHandler(server)

	auth := server.Echo.Group("")
	auth.Use(mw.Security(false))

	limiterStore := middleware.NewRateLimiterMemoryStore(rate.Limit(time.Minute))

	auth.POST("/auth/login", authHandler.Login, middleware.RateLimiter(limiterStore))
	auth.GET("/auth/refresh", authHandler.RefreshToken)
	auth.GET("/auth/logout", authHandler.Logout)
}
