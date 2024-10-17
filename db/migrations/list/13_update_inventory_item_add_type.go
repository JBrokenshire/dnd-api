package list

import (
	"github.com/ShkrutDenis/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type UpdateInventoryItemAddType struct{}

func (m *UpdateInventoryItemAddType) GetName() string {
	return "UpdateInventoryItemAddType"
}

func (m *UpdateInventoryItemAddType) Up(con *sqlx.DB) {
	table := builder.ChangeTable("character_inventory_items", con)
	table.Column("type").Type("ENUM('item','armour','weapon','shield')").NotNull()
	table.MustExec()
}

func (m *UpdateInventoryItemAddType) Down(con *sqlx.DB) {
	table := builder.ChangeTable("character_inventory_items", con)
	table.DropColumn("type")
	table.MustExec()
}
