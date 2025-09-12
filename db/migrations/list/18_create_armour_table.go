package list

import (
	"dnd-api/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateArmoursTable struct{}

func (m *CreateArmoursTable) GetName() string {
	return "CreateArmoursTable"
}

func (m *CreateArmoursTable) Up(con *sqlx.DB) {
	table := builder.NewTable("armours", con)
	table.Column("item_id").Type("int unsigned").NotNull()
	table.PrimaryKey("item_id")
	table.Integer("base_ac").NotNull()
	table.Integer("str_requirement").Nullable()
	table.Column("type").Type("ENUM('Light Armour','Medium Armour','Heavy Armour')").NotNull()
	table.Integer("max_dex_modifier").Nullable()
	table.Column("stealth_disadvantage").Type("boolean").NotNull().Default("false")
	table.MustExec()
}

func (m *CreateArmoursTable) Down(con *sqlx.DB) {
	builder.DropTable("armours", con).MustExec()
}
