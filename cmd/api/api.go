package main

import (
	"fmt"
	"github.com/JBrokenshire/dnd-api/api"
	"github.com/JBrokenshire/dnd-api/api/routes"
	"github.com/JBrokenshire/dnd-api/docs"
	"github.com/JBrokenshire/dnd-api/pkg/validation"
	"github.com/go-playground/validator/v10"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// @title D&D Hub
// @version 1.0
// @description Hub of information for Dungeons & Dragons games

// @contact.name Jared Brokenshire
// @contact.email jbrokenshire0306@gmail.com

// @host api.localdnd.com:7788
// @BasePath /
func main() {

	log.Println("Starting D&D")
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
