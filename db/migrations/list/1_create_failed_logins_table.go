package list

import (
	"dnd-api/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateFailedLoginsTable struct{}

func (m *CreateFailedLoginsTable) GetName() string {
	return "CreateFailedLoginsTable"
}

func (m *CreateFailedLoginsTable) Up(con *sqlx.DB) {
	table := builder.NewTable("failed_logins", con)
	table.Column("id").Type("int unsigned").NotNull().Autoincrement()
	table.PrimaryKey("id")
	table.Column("user_id").Type("int unsigned").NotNull()
	table.WithTimestamps()
	table.MustExec()
}

func (m *CreateFailedLoginsTable) Down(con *sqlx.DB) {
	builder.DropTable("failed_logins", con).MustExec()
}
