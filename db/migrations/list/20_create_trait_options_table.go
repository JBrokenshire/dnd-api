package list

import (
	"dnd-api/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateTraitOptionsTable struct{}

func (m *CreateTraitOptionsTable) GetName() string {
	return "CreateTraitOptionsTable"
}

func (m *CreateTraitOptionsTable) Up(con *sqlx.DB) {
	table := builder.NewTable("trait_options", con)
	table.Column("id").Type("int unsigned").NotNull().Autoincrement()
	table.PrimaryKey("id")
	table.String("trait_id", 64).NotNull()
	table.String("name", 64).NotNull()
	table.Column("description").Type("text").NotNull()
	table.MustExec()
}

func (m *CreateTraitOptionsTable) Down(con *sqlx.DB) {
	builder.DropTable("trait_options", con).MustExec()
}
