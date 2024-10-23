package list

import (
	"github.com/ShkrutDenis/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type UpdateClassFeaturesAddLevel struct{}

func (m *UpdateClassFeaturesAddLevel) GetName() string {
	return "UpdateClassFeaturesAddLevel"
}

func (m *UpdateClassFeaturesAddLevel) Up(con *sqlx.DB) {
	table := builder.ChangeTable("class_features", con)
	table.Integer("level").NotNull().Default("1")
	table.MustExec()
}

func (m *UpdateClassFeaturesAddLevel) Down(con *sqlx.DB) {
	table := builder.ChangeTable("class_features", con)
	table.DropColumn("level")
	table.MustExec()
}
