package routes

import (
	"dnd-api/server"
)

func ConfigureRoutes(server *server.Server) {
	charactersRoutes(server)
	classRoutes(server)
	raceRoutes(server)
	itemRoutes(server)
	weaponRoutes(server)
}
