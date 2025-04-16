package process

import (
	"dnd-api/db/migrations/list"
	"dnd-api/pkg/go-migrations"
	"dnd-api/pkg/go-migrations/store"
)

func Run() {
	go_migrations.Run(getMigrationsList())
}

func getMigrationsList() []store.Migratable {
	return []store.Migratable{
		&list.CreateUsersTable{},
		&list.CreateFailedLoginsTable{},
		&list.CreateClassesTable{},
		&list.CreateRacesTable{},
		&list.CreateCharactersTable{},
	}
}
