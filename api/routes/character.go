package routes

import (
	"dnd-api/api"
	"dnd-api/api/handlers"
)

func characterRoutes(server *api.Server) {
	characterHandler := handlers.NewCharacterHandler(server)
	inspirationHandler := handlers.NewCharacterInspirationHandler(server)
	healthHandler := handlers.NewCharacterHealthHandler(server)
	inventoryItemHandler := handlers.NewCharacterInventoryItemHandler(server)
	spellSlotHandler := handlers.NewCharacterUsedSpellSlotHandler(server)

	character := restrictedRouteGroup(server, "/characters")

	character.GET("", characterHandler.List)
	character.GET("/:id", characterHandler.Get)
	character.POST("", characterHandler.Create)
	character.PUT("/:id", characterHandler.Update)
	character.DELETE("/:id", characterHandler.Delete)
	character.POST("/:id/upload/profile-picture", characterHandler.UploadProfilePicture)

	// Inspiration
	character.PUT("/:id/inspiration", inspirationHandler.Update)

	// Health
	character.PUT("/:id/health", healthHandler.Update)

	// Inventory Items
	character.PUT("/:id/inventory-item/:inventoryItemId", inventoryItemHandler.Update)

	// Used Spell Slots
	character.PUT("/:id/spell-slots/:spellLevel", spellSlotHandler.Update)
}
