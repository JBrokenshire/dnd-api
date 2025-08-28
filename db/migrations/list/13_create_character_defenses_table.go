package list

import (
	"dnd-api/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateCharacterDefensesTable struct{}

func (m *CreateCharacterDefensesTable) GetName() string {
	return "CreateCharacterDefensesTable"
}

func (m *CreateCharacterDefensesTable) Up(con *sqlx.DB) {
	table := builder.NewTable("character_defenses", con)
	table.Column("id").Type("int unsigned").NotNull().Autoincrement()
	table.PrimaryKey("id")
	table.Column("character_id").Type("int unsigned").NotNull()
	table.Column("damage_type").Type("ENUM('Acid','Bludgeoning','Piercing','Slashing','Cold','Fire','Lightning','Thunder','Poison','Necrotic','Radiant','Force','Psychic')").NotNull()
	table.Column("defense_type").Type("ENUM('Resistance','Immunity','Vulnerability')").NotNull()
	table.MustExec()
}

func (m *CreateCharacterDefensesTable) Down(con *sqlx.DB) {
	builder.DropTable("character_defenses", con)
}
