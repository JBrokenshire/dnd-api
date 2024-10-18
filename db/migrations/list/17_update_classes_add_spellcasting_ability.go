package list

import (
	"github.com/ShkrutDenis/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type UpdateClassesAddSpellcastingAbility struct{}

func (m *UpdateClassesAddSpellcastingAbility) GetName() string {
	return "UpdateClassesAddSpellcastingAbility"
}

func (m *UpdateClassesAddSpellcastingAbility) Up(con *sqlx.DB) {
	table := builder.ChangeTable("classes", con)
	table.Column("spellcasting_ability").Type("ENUM('Intelligence','Wisdom','Charisma')").Nullable()
	table.MustExec()
}

func (m *UpdateClassesAddSpellcastingAbility) Down(con *sqlx.DB) {
	table := builder.ChangeTable("classes", con)
	table.DropColumn("spellcasting_ability")
	table.MustExec()
}
