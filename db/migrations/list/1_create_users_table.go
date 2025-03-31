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
	table.Column("id").Type("int unsigned").Autoincrement()
	table.PrimaryKey("id")
	table.String("uid", 100).Unique()
	table.String("email", 200).Unique()
	table.String("name", 500).Nullable()
	table.String("password", 1000).NotNull()
	table.Column("deleted_at").Type("datetime").Nullable()
	table.Column("super_admin").Type("boolean") // is this a super admin role?
	table.WithTimestamps()
	table.MustExec()
}

func (m *CreateUsersTable) Down(con *sqlx.DB) {
	builder.DropTable("users", con).MustExec()
}
