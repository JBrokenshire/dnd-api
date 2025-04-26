package list

import (
	"dnd-api/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type UpdateCharactersAddDetails struct{}

func (m *UpdateCharactersAddDetails) GetName() string {
	return "UpdateCharactersAddDetails"
}

func (m *UpdateCharactersAddDetails) Up(con *sqlx.DB) {
	table := builder.ChangeTable("characters", con)
	table.String("pronouns", 32).Nullable()
	table.Integer("level").NotNull().Default("1")
	table.MustExec()
}

func (m *UpdateCharactersAddDetails) Down(con *sqlx.DB) {
	table := builder.ChangeTable("characters", con)
	table.DropColumn("pronouns")
	table.DropColumn("level")
	table.MustExec()
}
