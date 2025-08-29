package list

import (
	"dnd-api/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateBackgroundsTable struct{}

func (m *CreateBackgroundsTable) GetName() string {
	return "CreateBackgroundsTable"
}

func (m *CreateBackgroundsTable) Up(con *sqlx.DB) {
	table := builder.NewTable("backgrounds", con)
	table.Column("id").Type("int unsigned").NotNull().Autoincrement()
	table.PrimaryKey("id")
	table.String("name", 64).NotNull()
	table.String("feature", 128).NotNull()
	table.Column("description").Type("LONGTEXT").NotNull()
	table.MustExec()
}

func (m *CreateBackgroundsTable) Down(con *sqlx.DB) {
	builder.DropTable("backgrounds", con).MustExec()
}
