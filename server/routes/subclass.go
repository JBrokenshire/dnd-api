package routes

import (
	"dnd-api/server"
	"dnd-api/server/controllers"
)

func subclassRoutes(server *server.Server) {
	subclassController := controllers.SubclassController{Server: *server}

	subclasses := server.Echo.Group("/subclasses")
	subclasses.GET("/:id/features", subclassController.GetFeatures)
}
