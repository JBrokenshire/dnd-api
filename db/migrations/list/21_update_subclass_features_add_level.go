package list

import (
	"github.com/ShkrutDenis/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type UpdateSubclassFeaturesAddLevel struct{}

func (m *UpdateSubclassFeaturesAddLevel) GetName() string {
	return "UpdateSubclassFeaturesAddLevel"
}

func (m *UpdateSubclassFeaturesAddLevel) Up(con *sqlx.DB) {
	table := builder.ChangeTable("subclass_features", con)
	table.Integer("level").NotNull().Default("1")
	table.MustExec()
}

func (m *UpdateSubclassFeaturesAddLevel) Down(con *sqlx.DB) {
	table := builder.ChangeTable("subclass_features", con)
	table.DropColumn("level")
	table.MustExec()
}
