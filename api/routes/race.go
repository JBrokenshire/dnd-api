package routes

import (
	s "dnd-api/api"
	"dnd-api/api/handlers"
	m "dnd-api/db/models"
	mw "dnd-api/pkg/middleware"
)

func raceRoutes(server *s.Server) {
	raceHandler := handlers.NewRaceHandler(server)

	races := server.Echo.Group("/races")

	races.GET("", raceHandler.List, mw.HasPermission(m.SubjectRace, m.ActionRead))
	races.GET("/:id", raceHandler.Get, mw.HasPermission(m.SubjectRace, m.ActionRead))
	races.POST("", raceHandler.Create, mw.HasPermission(m.SubjectRace, m.ActionCreate))
	races.PUT("/:id", raceHandler.Update, mw.HasPermission(m.SubjectRace, m.ActionUpdate))
	races.DELETE("/:id", raceHandler.Delete, mw.HasPermission(m.SubjectRace, m.ActionDelete))
}
