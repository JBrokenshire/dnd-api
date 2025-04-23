package dependencies_test

import (
	"dnd-api/pkg/dependencies"
	"dnd-api/tests/helpers"
	"github.com/joho/godotenv"
	"testing"
)

func TestNewDependencyService(t *testing.T) {

	err := godotenv.Load("../../.env")
	if err != nil {
		t.Fatalf("Error loading .env file")
	}

	ds := dependencies.NewDependencyService(helpers.MockDb())

	ds.GetDB()

}
