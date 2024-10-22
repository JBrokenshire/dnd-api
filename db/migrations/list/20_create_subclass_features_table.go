package list

import (
	"github.com/ShkrutDenis/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateSubclassFeaturesTable struct{}

func (m *CreateSubclassFeaturesTable) GetName() string {
	return "CreateSubclassFeaturesTable"
}

func (m *CreateSubclassFeaturesTable) Up(con *sqlx.DB) {
	table := builder.NewTable("subclass_features", con)
	table.PrimaryKey("id")
	table.Integer("subclass_id").NotNull()
	table.ForeignKey("subclass_id").Reference("subclasses").On("id").OnUpdate("cascade").OnDelete("cascade")
	table.Integer("feature_id").NotNull()
	table.ForeignKey("feature_id").Reference("features").On("id").OnUpdate("cascade").OnDelete("cascade")
	table.MustExec()
}

func (m *CreateSubclassFeaturesTable) Down(con *sqlx.DB) {
	builder.DropTable("subclass_features", con).MustExec()
}
