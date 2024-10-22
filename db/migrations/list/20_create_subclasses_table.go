package list

import (
	"github.com/ShkrutDenis/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateSubclassesTable struct{}

func (m *CreateSubclassesTable) GetName() string {
	return "CreateSubclassesTable"
}

func (m *CreateSubclassesTable) Up(con *sqlx.DB) {
	table := builder.NewTable("subclasses", con)
	table.PrimaryKey("id")
	table.Integer("class_id").NotNull()
	table.ForeignKey("class_id").Reference("classes").On("id").OnUpdate("cascade").OnDelete("cascade")
	table.String("name", 255).NotNull()
	table.MustExec()
}

func (m *CreateSubclassesTable) Down(con *sqlx.DB) {
	builder.DropTable("subclasses", con).MustExec()
}
