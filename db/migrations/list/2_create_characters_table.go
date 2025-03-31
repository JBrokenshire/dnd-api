package list

import (
	mysql "dnd-api/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateCharactersTable struct{}

func (m *CreateCharactersTable) GetName() string {
	return "CreateCharactersTable"
}

func (m *CreateCharactersTable) Up(con *sqlx.DB) {
	table := mysql.NewTable("characters", con)
	table.Column("id").Type("int unsigned").Autoincrement()
	table.PrimaryKey("id")
	table.String("name", 200).NotNull()
	table.Column("race_id").Type("int unsigned").NotNull()
	table.ForeignKey("race_id").Reference("races").On("id").OnUpdate("cascade").OnDelete("cascade")
	table.Column("class_id").Type("int unsigned").NotNull()
	table.ForeignKey("class_id").Reference("classes").On("id").OnUpdate("cascade").OnDelete("cascade")
	table.MustExec()
}

func (m *CreateCharactersTable) Down(con *sqlx.DB) {
	mysql.DropTable("characters", con).MustExec()
}
