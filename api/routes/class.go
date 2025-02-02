package routes

import (
	"github.com/JBrokenshire/dnd-api/api"
	"github.com/JBrokenshire/dnd-api/api/handlers"
)

func classRoutes(server *api.Server) {

	classHandler := handlers.NewClassHandler(server)

	classes := server.Echo.Group("/classes")
	classes.GET("", classHandler.List)
	classes.POST("", classHandler.Create)
	classes.GET("/:id", classHandler.Get)
}
