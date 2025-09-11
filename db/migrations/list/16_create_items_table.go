package list

import (
	"dnd-api/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateItemsTable struct{}

func (m *CreateItemsTable) GetName() string {
	return "CreateItemsTable"
}

func (m *CreateItemsTable) Up(con *sqlx.DB) {
	table := builder.NewTable("items", con)
	table.Column("id").Type("int unsigned").NotNull().Autoincrement()
	table.PrimaryKey("id")
	table.String("name", 128).NotNull()
	table.Column("rarity").Type("ENUM('Common','Uncommon','Rare','Very Rare','Legendary')").NotNull().Default("Common")
	table.Column("origin").Type("text").NotNull().Default("")
	table.Column("weight").Type("float").NotNull().Default("0")
	table.Column("cost").Type("float").NotNull().Default("0")
	table.Column("equippable").Type("boolean").NotNull().Default("0")
	table.Column("notes").Type("text").NotNull()
	table.MustExec()
}

func (m *CreateItemsTable) Down(con *sqlx.DB) {
	builder.DropTable("items", con).MustExec()
}
