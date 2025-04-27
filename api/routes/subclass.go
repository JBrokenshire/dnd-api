package routes

import (
	"dnd-api/api"
	"dnd-api/api/handlers"
)

func subclassRoutes(server *api.Server) {
	subclassHandler := handlers.NewSubclassHandler(server)

	subclass := restrictedRouteGroup(server, "/subclasses")

	subclass.GET("/:classId", subclassHandler.List)
	subclass.GET("/:classId/:subclassId", subclassHandler.Get)
	subclass.POST("", subclassHandler.Create)
	subclass.PUT("/:classId/:subclassId", subclassHandler.Update)
	subclass.DELETE("/:classId/:subclassId", subclassHandler.Delete)
	subclass.POST("/:classId/:subclassId/upload/logo", subclassHandler.UploadLogo)
}
