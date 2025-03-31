package list

import (
	"dnd-api/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateRolesTable struct{}

func (m *CreateRolesTable) GetName() string {
	return "CreateRolesTable"
}

func (m *CreateRolesTable) Up(con *sqlx.DB) {
	table := builder.NewTable("roles", con)
	table.Column("id").Type("int unsigned").Autoincrement()
	table.PrimaryKey("id")
	table.String("name", 200).NotNull().Unique()
	table.Column("created_by").Type("int").Nullable()
	table.Column("deleted_by").Type("int").Nullable()
	table.Column("deleted_at").Type("datetime").Nullable()
	table.WithTimestamps()
	table.MustExec()

	// Add in a main Admin role which every enterprise will have.
	con.MustExec("INSERT INTO `roles` (`name`, `created_by`) VALUES ('Admin', '1');")
}

func (m *CreateRolesTable) Down(con *sqlx.DB) {
	builder.DropTable("roles", con).MustExec()
}
