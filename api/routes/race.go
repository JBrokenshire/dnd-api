package routes

import (
	"dnd-api/api"
	"dnd-api/api/handlers"
)

func raceRoutes(server *api.Server) {
	raceHandler := handlers.NewRaceHandler(server)

	race := restrictedRouteGroup(server, "/races")

	race.GET("", raceHandler.List)
	race.GET("/:id", raceHandler.Get)
	race.POST("", raceHandler.Create)
	race.PUT("/:id", raceHandler.Update)
	race.DELETE("/:id", raceHandler.Delete)
}
