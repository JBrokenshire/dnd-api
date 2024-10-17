package list

import (
	"github.com/ShkrutDenis/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateArmourTable struct{}

func (m *CreateArmourTable) GetName() string { return "CreateArmourTable" }

func (m *CreateArmourTable) Up(con *sqlx.DB) {
	table := builder.NewTable("armour", con)
	table.Integer("item_id").Unique().NotNull().NotAutoincrement()
	table.PrimaryKey("item_id")
	table.ForeignKey("item_id").Reference("items").On("id").OnUpdate("cascade").OnDelete("cascade")
	table.Column("type").Type("ENUM('Light Armour', 'Medium Armour', 'Heavy Armour', 'Shields')").NotNull()
	table.Integer("base_ac").NotNull()
	table.Integer("max_dexterity_modifier").NotNull().Default("0")
	table.Integer("strength_requirement").Nullable()
	table.Column("stealth_disadvantage").Type("BOOLEAN").NotNull()
	table.MustExec()
}

func (m *CreateArmourTable) Down(con *sqlx.DB) {
	builder.DropTable("armour", con).MustExec()
}
