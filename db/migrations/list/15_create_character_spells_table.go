package list

import (
	"dnd-api/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateCharacterSpellsTable struct{}

func (m *CreateCharacterSpellsTable) GetName() string {
	return "CreateCharacterSpellsTable"
}

func (m *CreateCharacterSpellsTable) Up(con *sqlx.DB) {
	table := builder.NewTable("character_spells", con)
	table.Column("id").Type("int unsigned").NotNull().Autoincrement()
	table.PrimaryKey("id")
	table.Column("character_id").Type("int unsigned").NotNull()
	table.Column("spell_id").Type("int unsigned").NotNull()
	table.MustExec()
}

func (m *CreateCharacterSpellsTable) Down(con *sqlx.DB) {
	builder.DropTable("character_spells", con).MustExec()
}
