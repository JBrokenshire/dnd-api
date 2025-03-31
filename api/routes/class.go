package routes

import (
	s "dnd-api/api"
	"dnd-api/api/handlers"
)

func classRoutes(server *s.Server) {
	classHandler := handlers.NewClassHandler(server)

	classes := server.Echo.Group("/classes")

	classes.GET("", classHandler.List)
	classes.GET("/:id", classHandler.Get)
	classes.POST("", classHandler.Create)
	classes.PUT("/:id", classHandler.Update)
	classes.DELETE("/:id", classHandler.Delete)
}
