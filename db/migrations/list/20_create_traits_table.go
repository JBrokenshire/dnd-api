package list

import (
	"dnd-api/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateTraitsTable struct{}

func (m *CreateTraitsTable) GetName() string {
	return "CreateTraitsTable"
}

func (m *CreateTraitsTable) Up(con *sqlx.DB) {
	table := builder.NewTable("traits", con)
	table.String("id", 64).NotNull()
	table.PrimaryKey("id")
	table.String("name", 64).NotNull()
	table.Column("description").Type("text").NotNull()
	table.MustExec()
}

func (m *CreateTraitsTable) Down(con *sqlx.DB) {
	builder.DropTable("traits", con).MustExec()
}
