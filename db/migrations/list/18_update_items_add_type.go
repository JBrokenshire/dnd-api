package list

import (
	"dnd-api/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type UpdateItemsAddType struct{}

func (m *UpdateItemsAddType) GetName() string {
	return "UpdateItemsAddType"
}

func (m *UpdateItemsAddType) Up(con *sqlx.DB) {
	table := builder.ChangeTable("items", con)
	table.Column("type").Type("ENUM('armour','weapon','item')").NotNull().Default("item")
	table.MustExec()
}

func (m *UpdateItemsAddType) Down(con *sqlx.DB) {
	table := builder.ChangeTable("items", con)
	table.DropColumn("type")
	table.MustExec()
}
