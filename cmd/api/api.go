package main

import (
	version "dnd-api"
	"dnd-api/api"
	"dnd-api/api/routes"
	"dnd-api/pkg/validation"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"log"
	"os"
)

// @title D&D Character Hub API
// @version 0.0.1
// @description Solution for managing character sheet information for Dungeons & Dragon 5E

// @contact.name Jared Brokenshire
// @contact.email jbrokenshire0306@gmail.com

// @BasePath /
func main() {
	log.Printf("Starting D&D API. v%v\n", version.Get())

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	app := api.NewServer()
	app.Echo.Validator = validation.NewCustomValidator(validator.New())
	// Check Permission and Sync
	app.SetupPermissions()

	routes.ConfigureRoutes(app)
	err = app.Start(os.Getenv("PORT"))
	if err != nil {
		log.Fatal("Port already in use")
	}
}
