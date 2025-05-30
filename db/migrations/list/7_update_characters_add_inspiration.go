package list

import (
	"dnd-api/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type UpdateCharactersAddInspiration struct{}

func (m *UpdateCharactersAddInspiration) GetName() string {
	return "UpdateCharactersAddInspiration"
}

func (m *UpdateCharactersAddInspiration) Up(con *sqlx.DB) {
	table := builder.ChangeTable("characters", con)
	table.Column("inspiration").Type("BOOLEAN").NotNull().Default("0")
	table.MustExec()
}

func (m *UpdateCharactersAddInspiration) Down(con *sqlx.DB) {
	table := builder.ChangeTable("characters", con)
	table.DropColumn("inspiration")
	table.MustExec()
}
