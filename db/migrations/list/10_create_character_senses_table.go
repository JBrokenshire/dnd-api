package list

import (
	"dnd-api/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateCharacterSensesTable struct{}

func (m *CreateCharacterSensesTable) GetName() string {
	return "CreateCharacterSensesTable"
}

func (m *CreateCharacterSensesTable) Up(con *sqlx.DB) {
	table := builder.NewTable("character_senses", con)
	table.Column("id").Type("int unsigned").NotNull().Autoincrement()
	table.PrimaryKey("id")
	table.Column("character_id").Type("int unsigned").NotNull()
	table.String("sense", 64).NotNull()
	table.MustExec()
}

func (m *CreateCharacterSensesTable) Down(con *sqlx.DB) {
	builder.DropTable("character_senses", con).MustExec()
}
