package list

import (
	"dnd-api/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateCharactersTable struct{}

func (m *CreateCharactersTable) GetName() string {
	return "CreateCharactersTable"
}

func (m *CreateCharactersTable) Up(con *sqlx.DB) {
	table := builder.NewTable("characters", con)
	table.Column("id").Type("int unsigned").NotNull().Autoincrement()
	table.PrimaryKey("id")
	table.Column("user_id").Type("int unsigned").NotNull()
	table.String("name", 200).NotNull()
	table.Column("class_id").Type("int unsigned").NotNull()
	table.Column("race_id").Type("int unsigned").NotNull()
	table.MustExec()
}

func (m *CreateCharactersTable) Down(con *sqlx.DB) {
	builder.DropTable("characters", con).MustExec()
}
