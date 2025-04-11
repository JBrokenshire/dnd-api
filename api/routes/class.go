package routes

import (
	s "dnd-api/api"
	"dnd-api/api/handlers"
	m "dnd-api/db/models"
	mw "dnd-api/pkg/middleware"
)

func classRoutes(server *s.Server) {
	classHandler := handlers.NewClassHandler(server)

	classes := server.Echo.Group("/classes")

	classes.GET("", classHandler.List, mw.HasPermission(m.SubjectClass, m.ActionRead))
	classes.GET("/:id", classHandler.Get, mw.HasPermission(m.SubjectClass, m.ActionRead))
	classes.POST("", classHandler.Create, mw.HasPermission(m.SubjectClass, m.ActionCreate))
	classes.PUT("/:id", classHandler.Update, mw.HasPermission(m.SubjectClass, m.ActionUpdate))
	classes.DELETE("/:id", classHandler.Delete, mw.HasPermission(m.SubjectClass, m.ActionDelete))
}
