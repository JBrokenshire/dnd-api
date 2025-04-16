package list

import (
	"dnd-api/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateClassesTable struct{}

func (m *CreateClassesTable) GetName() string {
	return "CreateClassesTable"
}

func (m *CreateClassesTable) Up(con *sqlx.DB) {
	table := builder.NewTable("classes", con)
	table.Column("id").Type("int unsigned").NotNull().Autoincrement()
	table.PrimaryKey("id")
	table.String("name", 200).NotNull()
	table.MustExec()
}

func (m *CreateClassesTable) Down(con *sqlx.DB) {
	builder.DropTable("classes", con).MustExec()
}
