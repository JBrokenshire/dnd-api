package list

import (
	"dnd-api/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type UpdateCharactersAddHealth struct{}

func (m *UpdateCharactersAddHealth) GetName() string {
	return "UpdateCharactersAddHealth"
}

func (m *UpdateCharactersAddHealth) Up(con *sqlx.DB) {
	table := builder.ChangeTable("characters", con)
	table.Column("current_hit_points").Type("int unsigned").NotNull().Default("0")
	table.Column("max_hit_points").Type("int unsigned").NotNull().Default("0")
	table.Column("temp_hit_points").Type("int unsigned").NotNull().Default("0")
	table.MustExec()
}

func (m *UpdateCharactersAddHealth) Down(con *sqlx.DB) {
	table := builder.ChangeTable("characters", con)
	table.DropColumn("current_hit_points")
	table.DropColumn("max_hit_points")
	table.DropColumn("temp_hit_points")
	table.MustExec()
}
