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
		&list.CreateFilesTable{},
		&list.UpdateClassesAddDetails{},
		&list.UpdateRacesAddDetails{},
		&list.UpdateCharactersAddDetails{},
		&list.CreateSubclassesTable{},
		&list.UpdateCharactersAddAbilityScores{},
		&list.UpdateCharactersAddSettings{},
		&list.UpdateCharactersAddInspiration{},
		&list.UpdateCharactersAddHealth{},
		&list.UpdateCharactersAddSavingThrowAdjustments{},
		&list.UpdateCharactersAddSenses{},
		&list.UpdateCharactersAddProficiencies{},
	}
}
