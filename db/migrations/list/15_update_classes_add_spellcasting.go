package list

import (
	"dnd-api/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type UpdateClassesAddSpellcasting struct{}

func (m *UpdateClassesAddSpellcasting) GetName() string {
	return "UpdateClassesAddSpellcasting"
}

func (m *UpdateClassesAddSpellcasting) Up(con *sqlx.DB) {
	table := builder.ChangeTable("classes", con)
	table.Column("spellcasting_ability").Type("ENUM('Intelligence','Wisdom','Charisma')").Nullable()
	table.MustExec()
}

func (m *UpdateClassesAddSpellcasting) Down(con *sqlx.DB) {
	table := builder.ChangeTable("classes", con)
	table.DropColumn("spellcasting_ability")
	table.MustExec()
}
