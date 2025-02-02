package process

import (
	"github.com/JBrokenshire/dnd-api/db/migrations/process/list"
	gm "github.com/JBrokenshire/dnd-api/pkg/go-migrations"
	"github.com/JBrokenshire/dnd-api/pkg/go-migrations/store"
)

func Run() {
	gm.Run(getMigrationsList())
}

func getMigrationsList() []store.Migratable {
	return []store.Migratable{
		&list.CreateClassesTable{},
	}
}
