package list

import (
	"github.com/ShkrutDenis/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateSpellsTable struct{}

func (m *CreateSpellsTable) GetName() string {
	return "CreateSpellsTable"
}

func (m *CreateSpellsTable) Up(con *sqlx.DB) {
	table := builder.NewTable("spells", con)
	table.PrimaryKey("id")
	table.String("name", 128).NotNull()
	table.Integer("level").NotNull()
	table.Column("school").Type("ENUM('Abjuration', 'Conjuration', 'Divination', 'Enchantment', 'Evocation', 'Illusion', 'Necromancy', 'Transmutation')")
	table.String("casting_time", 8).NotNull()
	table.String("distance", 32).NotNull()
	table.String("effect", 32).Nullable()
	table.String("damage", 16).Nullable()
	table.Column("damage_type").Type("ENUM('Acid','Bludgeoning','Cold','Fire','Force','Lightning','Necrotic','Piercing','Poison','Psychic','Radiant','Slashing','Thunder')").Nullable()
	table.Column("save").Type("ENUM('STR','DEX','CON','INT','WIS','CHA')").Nullable()
	table.Column("notes").Type("MEDIUMTEXT").NotNull()
	table.MustExec()
}

func (m *CreateSpellsTable) Down(con *sqlx.DB) {
	builder.DropTable("spells", con).MustExec()
}
