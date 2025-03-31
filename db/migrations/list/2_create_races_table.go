package list

import (
	mysql "dnd-api/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateRacesTable struct{}

func (m *CreateRacesTable) GetName() string {
	return "CreateRacesTable"
}

func (m *CreateRacesTable) Up(con *sqlx.DB) {
	table := mysql.NewTable("races", con)
	table.Column("id").Type("int unsigned").Autoincrement()
	table.PrimaryKey("id")
	table.String("name", 200).NotNull()
	table.MustExec()
}

func (m *CreateRacesTable) Down(con *sqlx.DB) {
	mysql.DropTable("races", con).MustExec()
}
