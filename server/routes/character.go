package routes

import (
	"dnd-api/server"
	"dnd-api/server/controllers"
)

func charactersRoutes(server *server.Server) {
	characterController := controllers.CharacterController{Server: *server}
	skillsController := controllers.CharacterSkillsController{Server: *server}
	sensesController := controllers.CharacterSensesController{Server: *server}
	proficienciesController := controllers.CharacterProficienciesController{Server: *server}
	defensesController := controllers.CharacterDefensesController{Server: *server}
	conditionsController := controllers.CharacterConditionsController{Server: *server}
	inventoryController := controllers.CharacterInventoryController{Server: *server}
	moneyController := controllers.CharacterMoneyController{Server: *server}

	characters := server.Echo.Group("/characters")
	characters.GET("", characterController.GetAll)
	characters.POST("", characterController.Create)

	characters.GET("/:id", characterController.Get)
	characters.PUT("/:id", characterController.Update)
	characters.DELETE("/:id", characterController.Delete)
	characters.GET("/:id/inspiration", characterController.ToggleInspiration)
	characters.GET("/:id/level-up", characterController.LevelUp)
	characters.PUT("/:id/heal/:value", characterController.Heal)
	characters.PUT("/:id/damage/:value", characterController.Damage)
	characters.GET("/:id/armour-class", characterController.GetArmourClass)

	characters.GET("/:id/proficient-skills", skillsController.GetProficientByCharacterID)
	characters.GET("/:id/skills-advantages", skillsController.GetAdvantagesByCharacterID)
	characters.PUT("/:id/skills/:name/disadvantage", skillsController.ToggleCharacterSkillDisadvantage)
	characters.GET("/:id/senses", sensesController.GetSensesByCharacterID)

	characters.GET("/:id/proficient/armour", proficienciesController.GetCharacterProficientArmourTypes)
	characters.GET("/:id/proficient/weapons", proficienciesController.GetCharacterProficientWeapons)
	characters.GET("/:id/proficient/tools", proficienciesController.GetCharacterProficientTools)
	characters.GET("/:id/proficient/languages", proficienciesController.GetCharacterLanguages)

	characters.GET("/:id/defenses", defensesController.GetCharacterDefenses)
	characters.GET("/:id/conditions", conditionsController.GetCharacterConditions)

	characters.GET("/:id/inventory", inventoryController.GetCharacterInventory)
	characters.GET("/:id/inventory/equipped-weapons", inventoryController.GetCharacterEquippedWeapons)
	characters.GET("/:id/inventory/equipped-armour", inventoryController.GetCharacterEquippedArmour)
	characters.GET("/:id/inventory/money", moneyController.GetCharacterMoney)
	characters.PUT("/:characterID/inventory/:itemID", inventoryController.ToggleItemEquipped)
}
