package list

import (
	"dnd-api/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type UpdateCharactersAddAbilityScores struct{}

func (m *UpdateCharactersAddAbilityScores) GetName() string {
	return "UpdateCharactersAddAbilityScores"
}

func (m *UpdateCharactersAddAbilityScores) Up(con *sqlx.DB) {
	table := builder.ChangeTable("characters", con)
	table.Column("strength").Type("int unsigned").NotNull().Default("10")
	table.Column("dexterity").Type("int unsigned").NotNull().Default("10")
	table.Column("constitution").Type("int unsigned").NotNull().Default("10")
	table.Column("intelligence").Type("int unsigned").NotNull().Default("10")
	table.Column("wisdom").Type("int unsigned").NotNull().Default("10")
	table.Column("charisma").Type("int unsigned").NotNull().Default("10")
	table.MustExec()
}

func (m *UpdateCharactersAddAbilityScores) Down(con *sqlx.DB) {
	table := builder.ChangeTable("characters", con)
	table.DropColumn("strength")
	table.DropColumn("dexterity")
	table.DropColumn("constitution")
	table.DropColumn("intelligence")
	table.DropColumn("wisdom")
	table.DropColumn("charisma")
	table.MustExec()
}
