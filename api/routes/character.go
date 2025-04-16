package routes

import (
	"dnd-api/api"
	"dnd-api/api/handlers"
)

func characterRoutes(server *api.Server) {
	characterHandler := handlers.NewCharacterHandler(server)

	character := restrictedRouteGroup(server, "/characters")

	character.GET("", characterHandler.List)
	character.GET("/:id", characterHandler.Get)
	character.POST("", characterHandler.Create)
	character.PUT("/:id", characterHandler.Update)
	character.DELETE("/:id", characterHandler.Delete)
}
