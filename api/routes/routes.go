package routes

import (
	s "github.com/JBrokenshire/dnd-api/api"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ConfigureRoutes(server *s.Server) {
	// Configure CORS
	server.Echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"http://localdnd.com:3000",
		},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowCredentials: true,
	}))

	// Class routes
	// PATH: /classes
	classRoutes(server)
}
