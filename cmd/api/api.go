package main

import (
	version "dnd-api"
	"dnd-api/api"
	"dnd-api/api/routes"
	"dnd-api/docs"
	"dnd-api/pkg/validation"
	"fmt"
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

// @title D&D Character Dashboard
// @version 0.0.1
// @description A management solution for Dungeons & Dragons character sheet information

// @contact.name Jared Brokenshire
// @contact.email jbrokenshire0306@gmail.com

// @BasePath /
func main() {

	version.Get()

	log.Printf("Starting D&D API. v%v\n", version.Get())
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	if os.Getenv("ENVIRONMENT") == "development" {
		docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("EXPOSE_PORT"))
	}

	app := api.NewServer()
	app.Echo.Validator = validation.NewCustomValidator(validator.New())

	routes.ConfigureRoutes(app)
	err = app.Start(os.Getenv("PORT"))
	if err != nil {
		log.Fatal("Port already used")
	}
}
