package list

import (
	"dnd-api/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type UpdateCharactersAddSenses struct{}

func (m *UpdateCharactersAddSenses) GetName() string {
	return "UpdateCharactersAddSenses"
}

func (m *UpdateCharactersAddSenses) Up(con *sqlx.DB) {
	table := builder.ChangeTable("characters", con)
	table.Column("senses").Type("LONGTEXT").NotNull()
	table.MustExec()
}

func (m *UpdateCharactersAddSenses) Down(con *sqlx.DB) {
	table := builder.ChangeTable("characters", con)
	table.DropColumn("senses")
	table.MustExec()
}
