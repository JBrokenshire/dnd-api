package list

import (
	"dnd-api/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateCharacterInventoryItemsTable struct{}

func (m *CreateCharacterInventoryItemsTable) GetName() string {
	return "CreateCharacterInventoryItemsTable"
}

func (m *CreateCharacterInventoryItemsTable) Up(con *sqlx.DB) {
	table := builder.NewTable("character_inventory_items", con)
	table.Column("id").Type("int unsigned").NotNull().Autoincrement()
	table.PrimaryKey("id")
	table.Column("character_id").Type("int unsigned").NotNull()
	table.Column("item_id").Type("int unsigned").NotNull()
	table.String("location", 32).NotNull().Default("Backpack")
	table.Integer("quantity").NotNull().Default("1")
	table.Column("equipped").Type("boolean").Nullable()
	table.MustExec()
}

func (m *CreateCharacterInventoryItemsTable) Down(con *sqlx.DB) {
	builder.DropTable("character_inventory_items", con).MustExec()
}
