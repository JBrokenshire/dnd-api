package list

import (
	"dnd-api/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateSubclassesTable struct{}

func (m *CreateSubclassesTable) GetName() string {
	return "CreateSubclassesTable"
}

func (m *CreateSubclassesTable) Up(con *sqlx.DB) {
	table := builder.NewTable("subclasses", con)
	table.Column("id").Type("int unsigned").NotNull().Autoincrement()
	table.PrimaryKey("id")
	table.Column("class_id").Type("int unsigned").NotNull()
	table.String("name", 200).NotNull()
	table.Column("short_description").Type("text").NotNull()
	table.MustExec()
}

func (m *CreateSubclassesTable) Down(con *sqlx.DB) {
	builder.DropTable("subclasses", con).MustExec()
}
