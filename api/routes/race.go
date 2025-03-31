package routes

import (
	s "dnd-api/api"
	"dnd-api/api/handlers"
)

func raceRoutes(server *s.Server) {
	raceHandler := handlers.NewRaceHandler(server)

	races := server.Echo.Group("/races")

	races.GET("", raceHandler.List)
	races.GET("/:id", raceHandler.Get)
	races.POST("", raceHandler.Create)
	races.PUT("/:id", raceHandler.Update)
	races.DELETE("/:id", raceHandler.Delete)
}
