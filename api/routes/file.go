package routes

import (
	"dnd-api/api"
	"dnd-api/api/handlers"
)

func fileRoutes(server *api.Server) {
	fileHandler := handlers.NewFileHandler(server)

	file := server.Echo.Group("/files")

	file.GET("/:filepath", fileHandler.Get)
}
