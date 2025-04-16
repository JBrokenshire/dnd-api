package list

import (
	"dnd-api/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateRacesTable struct{}

func (m *CreateRacesTable) GetName() string {
	return "CreateRacesTable"
}

func (m *CreateRacesTable) Up(con *sqlx.DB) {
	table := builder.NewTable("races", con)
	table.Column("id").Type("int unsigned").NotNull().Autoincrement()
	table.PrimaryKey("id")
	table.String("name", 200).NotNull()
	table.MustExec()
}

func (m *CreateRacesTable) Down(con *sqlx.DB) {
	builder.DropTable("races", con).MustExec()
}
