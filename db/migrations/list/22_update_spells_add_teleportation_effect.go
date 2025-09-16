package list

import (
	"dnd-api/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type UpdateSpellsEffectType struct{}

func (m *UpdateSpellsEffectType) GetName() string {
	return "UpdateSpellsEffectType"
}

func (m *UpdateSpellsEffectType) Up(con *sqlx.DB) {
	table := builder.ChangeTable("spells", con)
	table.String("effect", 64).NotNull().Default("Damage").Change()
	table.MustExec()
}

func (m *UpdateSpellsEffectType) Down(con *sqlx.DB) {
	table := builder.ChangeTable("spells", con)
	table.Column("effect").Change().Type("ENUM('Damage','Utility','Charmed','Social','Incapacitated','Control','Buff')").NotNull().Default("Damage")
	table.MustExec()
}
