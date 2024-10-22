package list

import (
	"github.com/ShkrutDenis/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateFeaturesTable struct{}

func (m *CreateFeaturesTable) GetName() string {
	return "CreateFeaturesTable"
}

func (m *CreateFeaturesTable) Up(con *sqlx.DB) {
	table := builder.NewTable("features", con)
	table.PrimaryKey("id")
	table.String("name", 255).NotNull()
	table.Column("description").Type("LONGTEXT").NotNull()
	table.Column("action").Type("MEDIUMTEXT").Nullable()
	table.Column("action_type").Type("ENUM('Action','Bonus Action','Reaction')").Nullable()
	table.Integer("action_uses").Nullable()
	table.Column("action_reset").Type("ENUM('Short Rest','Long Rest')").Nullable()
	table.MustExec()
}

func (m *CreateFeaturesTable) Down(con *sqlx.DB) {
	builder.DropTable("features", con).MustExec()
}
