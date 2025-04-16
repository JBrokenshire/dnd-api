package routes

import (
	"dnd-api/api"
	"dnd-api/api/handlers"
)

func classRoutes(server *api.Server) {
	classHandler := handlers.NewClassHandler(server)

	class := restrictedRouteGroup(server, "/classes")

	class.GET("", classHandler.List)
	class.GET("/:id", classHandler.Get)
	class.POST("", classHandler.Create)
	class.PUT("/:id", classHandler.Update)
	class.DELETE("/:id", classHandler.Delete)
}
