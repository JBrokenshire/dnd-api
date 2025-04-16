package list

import (
	"dnd-api/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateUsersTable struct{}

func (m *CreateUsersTable) GetName() string {
	return "CreateUsersTable"
}

func (m *CreateUsersTable) Up(con *sqlx.DB) {
	table := builder.NewTable("users", con)
	table.Column("id").Type("int unsigned").NotNull().Autoincrement()
	table.PrimaryKey("id")
	table.String("username", 200).NotNull().Unique()
	table.String("password", 72).NotNull()
	table.Column("last_seen").Type("datetime").Nullable()
	table.Column("admin").Type("boolean").Default("0")
	table.Column("deleted_at").Type("datetime").Nullable()
	table.WithTimestamps()
	table.MustExec()
}

func (m *CreateUsersTable) Down(con *sqlx.DB) {
	builder.DropTable("users", con).MustExec()
}
