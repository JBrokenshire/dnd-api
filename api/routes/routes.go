package routes

import (
	s "dnd-api/api"
	mw "dnd-api/pkg/middleware"
	"dnd-api/services/jwt_service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
	"strings"
)

var (
	userJwtMiddleware echo.MiddlewareFunc
)

func ConfigureRoutes(server *s.Server) {

	// Set up JWT config
	config := jwt_service.JWTConfig{
		Claims:     &jwt_service.JwtCustomClaims{},
		SigningKey: []byte(os.Getenv("ACCESS_SECRET")),
	}
	userJwtMiddleware = jwt_service.JWTWithConfig(config)

	// Configure CORS
	server.Echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"http://localdnd.com:8080",
		},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowCredentials: true,
	}))

	// Make sure none of the queries are cached.
	server.Echo.Use(mw.NoCacheWithConfig(mw.NoCacheConfig{
		Skipper: func(c echo.Context) bool {
			// Skip no cache for FILES
			if strings.HasPrefix(c.Request().RequestURI, "/files") {
				return true
			}
			return false
		},
	}))

	// Add additional headers
	server.Echo.Use(mw.ServerHeader)

	// Auth Routes
	authRoutes(server)
	// User Routes
	userRoutes(server)
	// Class Routes
	classRoutes(server)
	// Race Routes
	raceRoutes(server)
	// Character Routes
	characterRoutes(server)

}

func restrictedRouteGroup(server *s.Server, prefix string) *echo.Group {
	// Secure Endpoints under restricted group with JWT protection
	restricted := server.Echo.Group(prefix)
	restricted.Use(userJwtMiddleware)
	restricted.Use(mw.Security(false))
	restricted.Use(mw.LoadUser(server.Db)) // Check user with Database. Add user to context
	restricted.Use(mw.Authorised)          //Make sure the JWT token Authorised Flag is True
	return restricted
}
