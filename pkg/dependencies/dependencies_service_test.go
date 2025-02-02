package dependencies_test

import (
	"github.com/JBrokenshire/dnd-api/pkg/dependencies"
	"github.com/JBrokenshire/dnd-api/test/helpers"
	"github.com/joho/godotenv"
	"testing"
)

func TestNewDependencyService(t *testing.T) {

	err := godotenv.Load("../../.env")
	if err != nil {
		t.Fatalf("Error loading .env file")
	}

	ds := dependencies.NewDependencyService(helpers.MockDb())

	ds.PreWarmServices()

	ds.GetDB()
	ds.CreateWg()
	ds.GetWg()
	ds.WgAdd()
	ds.WgDone()
}
