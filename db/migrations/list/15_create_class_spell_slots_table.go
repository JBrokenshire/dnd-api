package list

import (
	"dnd-api/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateClassSpellSlotsTable struct{}

func (m *CreateClassSpellSlotsTable) GetName() string {
	return "CreateClassSpellSlotsTable"
}

func (m *CreateClassSpellSlotsTable) Up(con *sqlx.DB) {
	table := builder.NewTable("class_spell_slots", con)
	table.Column("id").Type("int unsigned").NotNull().Autoincrement()
	table.PrimaryKey("id")
	table.Column("class_id").Type("int unsigned").NotNull()
	table.Integer("class_level").NotNull()
	table.Integer("spell_level").NotNull()
	table.Integer("spell_slots").NotNull().Default("0")
	table.MustExec()
}

func (m *CreateClassSpellSlotsTable) Down(con *sqlx.DB) {
	builder.DropTable("class_spell_slots", con).MustExec()
}
