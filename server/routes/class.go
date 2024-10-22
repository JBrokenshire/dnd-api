package routes

import (
	"dnd-api/server"
	"dnd-api/server/controllers"
)

func classRoutes(server *server.Server) {
	classController := controllers.ClassController{Server: *server}

	classes := server.Echo.Group("/classes")
	classes.GET("", classController.GetAll)
	classes.GET("/:id", classController.Get)
	classes.PUT("/:id", classController.Update)
	classes.GET("/:id/features", classController.GetFeatures)
}
