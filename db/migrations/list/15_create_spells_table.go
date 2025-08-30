package list

import (
	"dnd-api/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateSpellsTable struct{}

func (m *CreateSpellsTable) GetName() string {
	return "CreateSpellsTable"
}

func (m *CreateSpellsTable) Up(con *sqlx.DB) {
	table := builder.NewTable("spells", con)
	table.Column("id").Type("int unsigned").Nullable().Autoincrement()
	table.PrimaryKey("id")
	table.String("name", 64).NotNull()
	table.Column("school").Type("ENUM('Evocation','Conjuration','Necromancy','Abjuration','Transmutation','Divination','Enchantment','Illusion')").NotNull()
	table.Integer("level").NotNull().Default("0")
	table.String("casting_time", 8).NotNull()
	table.String("distance", 16).NotNull()
	table.Column("is_attack").Type("boolean").NotNull().Default("0")
	table.Column("is_save").Type("boolean").NotNull().Default("0")
	table.Column("save_ability").Type("ENUM('STR','DEX','CON','INT','WIS','CHA')").Nullable()
	table.Column("effect").Type("ENUM('Damage','Utility','Charmed','Social','Incapacitated','Control','Buff')").Default("Damage")
	table.String("damage", 8).Nullable()
	table.Column("damage_type").Type("ENUM('Acid','Bludgeoning','Cold','Fire','Force','Lightning','Necrotic','Piercing','Poison','Psychic','Radiant','Slashing','Thunder')").Nullable()
	table.String("notes", 64).NotNull()
	table.Column("can_upcast").Type("boolean").NotNull().Default("0")
	table.MustExec()
}

func (m *CreateSpellsTable) Down(con *sqlx.DB) {
	builder.DropTable("spells", con).MustExec()
}
