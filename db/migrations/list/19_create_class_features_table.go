package list

import (
	"github.com/ShkrutDenis/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateClassFeaturesTable struct{}

func (m *CreateClassFeaturesTable) GetName() string {
	return "CreateClassFeaturesTable"
}

func (m *CreateClassFeaturesTable) Up(con *sqlx.DB) {
	table := builder.NewTable("class_features", con)
	table.PrimaryKey("id")
	table.Integer("class_id").NotNull()
	table.ForeignKey("class_id").Reference("classes").On("id").OnUpdate("cascade").OnDelete("cascade")
	table.Integer("feature_id").NotNull()
	table.ForeignKey("feature_id").Reference("features").On("id").OnUpdate("cascade").OnDelete("cascade")
	table.MustExec()
}

func (m *CreateClassFeaturesTable) Down(con *sqlx.DB) {
	builder.DropTable("class_features", con).MustExec()
}
