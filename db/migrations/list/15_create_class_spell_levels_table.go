package list

import (
	"dnd-api/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateClassSpellLevelTable struct{}

func (m *CreateClassSpellLevelTable) GetName() string {
	return "CreateClassSpellLevelTable"
}

func (m *CreateClassSpellLevelTable) Up(con *sqlx.DB) {
	table := builder.NewTable("class_spell_levels", con)
	table.Column("id").Type("int unsigned").NotNull().Autoincrement()
	table.PrimaryKey("id")
	table.Column("class_id").Type("int unsigned").NotNull()
	table.Integer("class_level").NotNull()
	table.Integer("spell_level").NotNull()
	table.Integer("number_of_slots").NotNull()
	table.MustExec()
}

func (m *CreateClassSpellLevelTable) Down(con *sqlx.DB) {
	builder.DropTable("class_spell_levels", con).MustExec()
}
