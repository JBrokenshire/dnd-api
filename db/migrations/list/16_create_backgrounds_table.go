package list

import (
	"github.com/ShkrutDenis/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateBackgroundsTable struct{}

func (m *CreateBackgroundsTable) GetName() string {
	return "CreateBackgroundsTable"
}

func (m *CreateBackgroundsTable) Up(con *sqlx.DB) {
	table := builder.NewTable("backgrounds", con)
	table.String("name", 64).NotNull()
	table.PrimaryKey("name")
	table.String("feature", 64).NotNull()
	table.Column("description").Type("LONGTEXT").NotNull()
	table.MustExec()
}

func (m *CreateBackgroundsTable) Down(con *sqlx.DB) {
	builder.DropTable("backgrounds", con).MustExec()
}
