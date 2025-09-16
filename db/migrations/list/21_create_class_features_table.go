package list

import (
	"dnd-api/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateClassFeaturesTable struct{}

func (m *CreateClassFeaturesTable) GetName() string {
	return "CreateClassFeaturesTable"
}

func (m *CreateClassFeaturesTable) Up(con *sqlx.DB) {
	table := builder.NewTable("class_features", con)
	table.String("id", 64).NotNull()
	table.PrimaryKey("id")
	table.Column("class_id").Type("int unsigned").NotNull()
	table.Integer("level").NotNull()
	table.String("name", 64).NotNull()
	table.Column("description").Type("text").NotNull()
	table.Integer("priority").NotNull().Default("0")
	table.MustExec()
}

func (m *CreateClassFeaturesTable) Down(con *sqlx.DB) {
	builder.DropTable("class_features", con).MustExec()
}
