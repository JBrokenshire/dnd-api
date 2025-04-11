package routes

import (
	s "dnd-api/api"
	"dnd-api/api/handlers"
	m "dnd-api/db/models"
	mw "dnd-api/pkg/middleware"
)

func characterRoutes(server *s.Server) {
	characterHandler := handlers.NewCharacterHandler(server)

	characters := server.Echo.Group("/characters")

	characters.GET("", characterHandler.List, mw.HasPermission(m.SubjectCharacter, m.ActionRead))
	characters.GET("/:id", characterHandler.Get, mw.HasPermission(m.SubjectCharacter, m.ActionRead))
	characters.POST("", characterHandler.Create, mw.HasPermission(m.SubjectCharacter, m.ActionCreate))
	characters.PUT("/:id", characterHandler.Update, mw.HasPermission(m.SubjectCharacter, m.ActionUpdate))
	characters.DELETE("/:id", characterHandler.Delete, mw.HasPermission(m.SubjectCharacter, m.ActionDelete))
}
