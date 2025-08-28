package list

import (
	"dnd-api/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type UpdateCharactersAddProficiencies struct{}

func (m *UpdateCharactersAddProficiencies) GetName() string {
	return "UpdateCharactersAddProficiencies"
}

func (m *UpdateCharactersAddProficiencies) Up(con *sqlx.DB) {
	table := builder.ChangeTable("characters", con)
	table.Column("proficiencies").Type("LONGTEXT").NotNull()
	table.MustExec()
}

func (m *UpdateCharactersAddProficiencies) Down(con *sqlx.DB) {
	table := builder.ChangeTable("characters", con)
	table.DropColumn("proficiencies")
	table.MustExec()
}
