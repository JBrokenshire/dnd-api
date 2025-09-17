package list

import (
	"dnd-api/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type UpdateCharactersAddAttacksPerAction struct{}

func (m *UpdateCharactersAddAttacksPerAction) GetName() string {
	return "UpdateCharactersAddAttacksPerAction"
}

func (m *UpdateCharactersAddAttacksPerAction) Up(con *sqlx.DB) {
	table := builder.ChangeTable("characters", con)
	table.Integer("attacks_per_action").NotNull().Default("1")
	table.MustExec()
}

func (m *UpdateCharactersAddAttacksPerAction) Down(con *sqlx.DB) {
	table := builder.ChangeTable("characters", con)
	table.DropColumn("attacks_per_action")
	table.MustExec()
}
