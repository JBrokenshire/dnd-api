package list

import (
	"dnd-api/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateRolePermissionsTable struct{}

func (m *CreateRolePermissionsTable) GetName() string {
	return "CreateRolePermissionsTable"
}

func (m *CreateRolePermissionsTable) Up(con *sqlx.DB) {
	_ = con.MustExec("CREATE TABLE role_permissions(role_id int NOT NULL, permission_id int NOT NULL, PRIMARY KEY(role_id,permission_id));")
}

func (m *CreateRolePermissionsTable) Down(con *sqlx.DB) {
	builder.DropTable("role_permissions", con).MustExec()
}
