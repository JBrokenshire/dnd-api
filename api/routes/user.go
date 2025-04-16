package routes

import (
	"dnd-api/api"
	"dnd-api/api/handlers"
)

func userRoutes(server *api.Server) {
	userHandler := handlers.NewUserHandler(server)

	user := server.Echo.Group("/users")

	user.POST("", userHandler.Create)

}
