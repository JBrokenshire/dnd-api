package list

import (
	"dnd-api/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateCharacterUsedSpellSlotsTable struct{}

func (m *CreateCharacterUsedSpellSlotsTable) GetName() string {
	return "CreateCharacterUsedSpellSlotsTable"
}

func (m *CreateCharacterUsedSpellSlotsTable) Up(con *sqlx.DB) {
	table := builder.NewTable("character_used_spell_slots", con)
	table.Column("id").Type("int unsigned").NotNull().Autoincrement()
	table.PrimaryKey("id")
	table.Column("character_id").Type("int unsigned").NotNull()
	table.Integer("spell_level").NotNull()
	table.Integer("spell_slots_used").NotNull().Default("0")
	table.MustExec()
}

func (m *CreateCharacterUsedSpellSlotsTable) Down(con *sqlx.DB) {
	builder.DropTable("character_used_spell_slots", con).MustExec()
}
