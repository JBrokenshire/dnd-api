package list

import (
	"github.com/ShkrutDenis/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateTraitsTable struct{}

func (m *CreateTraitsTable) GetName() string {
	return "CreateTraitsTable"
}

func (m *CreateTraitsTable) Up(con *sqlx.DB) {
	table := builder.NewTable("traits", con)
	table.PrimaryKey("id")
	table.String("name", 255).NotNull()
	table.Column("description").Type("LONGTEXT").NotNull()
	table.Column("action").Type("MEDIUMTEXT").Nullable()
	table.Column("action_type").Type("ENUM('Action','Bonus Action','Reaction')").Nullable()
	table.Integer("action_uses").Nullable()
	table.Column("action_reset").Type("ENUM('Short Rest','Long Rest')").Nullable()
	table.MustExec()
}

func (m *CreateTraitsTable) Down(con *sqlx.DB) {
	builder.DropTable("traits", con).MustExec()
}
