package list

import (
	"github.com/ShkrutDenis/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateCharacterSpellsTable struct{}

func (m *CreateCharacterSpellsTable) GetName() string {
	return "CreateCharacterSpellsTable"
}

func (m *CreateCharacterSpellsTable) Up(con *sqlx.DB) {
	table := builder.NewTable("character_spells", con)
	table.PrimaryKey("id")
	table.Integer("character_id").NotNull()
	table.ForeignKey("character_id").Reference("characters").On("id").OnUpdate("cascade").OnDelete("cascade")
	table.Integer("spell_id").NotNull()
	table.ForeignKey("spell_id").Reference("spells").On("id").OnUpdate("cascade").OnDelete("cascade")
	table.Column("origin").Type("MEDIUMTEXT").NotNull()
	table.MustExec()
}

func (m *CreateCharacterSpellsTable) Down(con *sqlx.DB) {
	builder.DropTable("character_spells", con).MustExec()
}
