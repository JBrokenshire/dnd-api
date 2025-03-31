package list

import (
	"dnd-api/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreatePermissionsTable struct{}

func (m *CreatePermissionsTable) GetName() string {
	return "CreatePermissionsTable"
}

func (m *CreatePermissionsTable) Up(con *sqlx.DB) {
	table := builder.NewTable("permissions", con)
	table.Column("id").Type("int unsigned").Autoincrement()
	table.PrimaryKey("id")
	table.String("subject", 200).NotNull()
	table.String("action", 200).NotNull()
	table.String("description", 1000).Nullable()
	table.WithTimestamps()
	table.MustExec()

	// Let's make sure there is a composite index in subject and action
	con.MustExec("ALTER TABLE `permissions` ADD UNIQUE `subject_action_index`(`subject`, `action`);")
}

func (m *CreatePermissionsTable) Down(con *sqlx.DB) {
	builder.DropTable("permissions", con).MustExec()
}
