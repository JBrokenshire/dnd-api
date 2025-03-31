package list

import (
	"dnd-api/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateUserRolesTable struct{}

func (m *CreateUserRolesTable) GetName() string {
	return "CreateUserRolesTable"
}

func (m *CreateUserRolesTable) Up(con *sqlx.DB) {
	_ = con.MustExec("CREATE TABLE user_roles(user_id int NOT NULL, role_id int NOT NULL, PRIMARY KEY(user_id,role_id));")
}

func (m *CreateUserRolesTable) Down(con *sqlx.DB) {
	builder.DropTable("user_roles", con).MustExec()
}
