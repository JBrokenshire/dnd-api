package list

import (
	"dnd-api/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateWeaponsTable struct{}

func (m *CreateWeaponsTable) GetName() string {
	return "CreateWeaponsTable"
}

func (m *CreateWeaponsTable) Up(con *sqlx.DB) {
	table := builder.NewTable("weapons", con)
	table.Column("item_id").Type("int unsigned").NotNull()
	table.PrimaryKey("item_id")
	table.Column("weapon_type").Type("ENUM('Melee Weapon','Ranged Weapon')").NotNull()
	table.Integer("distance").NotNull().Default("5")
	table.Integer("alt_distance").Nullable()
	table.Column("ability").Type("ENUM('STR','DEX')").NotNull()
	table.String("damage", 16).NotNull()
	table.String("alt_damage", 16).Nullable()
	table.Column("damage_type").Type("ENUM('Acid','Bludgeoning','Cold','Fire','Force','Lightning','Necrotic','Piercing','Poison','Psychic','Radiant','Slashing','Thunder')").NotNull()
	table.Integer("bonus").NotNull().Default("0")
	table.Column("proficiencies").Type("text").NotNull()
	table.MustExec()
}

func (m *CreateWeaponsTable) Down(con *sqlx.DB) {
	builder.DropTable("weapons", con).MustExec()
}
