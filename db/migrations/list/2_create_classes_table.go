package list

import (
	mysql "dnd-api/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateClassesTable struct{}

func (m *CreateClassesTable) GetName() string {
	return "CreateClassesTable"
}

func (m *CreateClassesTable) Up(con *sqlx.DB) {
	table := mysql.NewTable("classes", con)
	table.Column("id").Type("int unsigned").Autoincrement()
	table.PrimaryKey("id")
	table.String("name", 200).NotNull()
	table.MustExec()
}

func (m *CreateClassesTable) Down(con *sqlx.DB) {
	mysql.DropTable("classes", con).MustExec()
}
