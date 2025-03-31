package process

import (
	"dnd-api/db/migrations/list"
	gm "dnd-api/pkg/go-migrations"
	gmStore "dnd-api/pkg/go-migrations/store"
)

func Run() {
	gm.Run(getMigrationList())
}

func getMigrationList() []gmStore.Migratable {
	return []gmStore.Migratable{
		&list.CreatePermissionsTable{},
		&list.CreateRolePermissionsTable{},
		&list.CreateRolesTable{},
		&list.CreateUserRolesTable{},
		&list.CreateUsersTable{},
		&list.CreateClassesTable{},
		&list.CreateRacesTable{},
		&list.CreateCharactersTable{},
	}
}
