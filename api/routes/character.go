package routes

import (
	s "dnd-api/api"
	"dnd-api/api/handlers"
)

func characterRoutes(server *s.Server) {
	characterHandler := handlers.NewCharacterHandler(server)

	characters := server.Echo.Group("/characters")

	characters.GET("", characterHandler.List)
	characters.GET("/:id", characterHandler.Get)
	characters.POST("", characterHandler.Create)
	characters.PUT("/:id", characterHandler.Update)
	characters.DELETE("/:id", characterHandler.Delete)
}
