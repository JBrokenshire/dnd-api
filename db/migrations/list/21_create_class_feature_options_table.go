package list

import (
	"dnd-api/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateClassFeatureOptionsTable struct{}

func (m *CreateClassFeatureOptionsTable) GetName() string {
	return "CreateClassFeatureOptionsTable"
}

func (m *CreateClassFeatureOptionsTable) Up(con *sqlx.DB) {
	table := builder.NewTable("class_feature_options", con)
	table.Column("id").Type("int unsigned").NotNull().Autoincrement()
	table.PrimaryKey("id")
	table.String("class_feature_id", 64).NotNull()
	table.String("name", 64).NotNull()
	table.Column("description").Type("text").NotNull()
	table.MustExec()
}

func (m *CreateClassFeatureOptionsTable) Down(con *sqlx.DB) {
	builder.DropTable("class_feature_options", con).MustExec()
}
