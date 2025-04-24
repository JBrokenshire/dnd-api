package list

import (
	"dnd-api/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type UpdateClassesAddDetails struct{}

func (m *UpdateClassesAddDetails) GetName() string {
	return "UpdateClassesAddShortDescription"
}

func (m *UpdateClassesAddDetails) Up(con *sqlx.DB) {
	table := builder.ChangeTable("classes", con)
	table.Column("short_description").Type("text").NotNull()
	table.String("primary_ability", 200).NotNull()
	table.Integer("hit_point_die_value").NotNull()
	table.String("saves", 200).NotNull()
	table.MustExec()
}

func (m *UpdateClassesAddDetails) Down(con *sqlx.DB) {
	table := builder.ChangeTable("classes", con)
	table.DropColumn("short_description")
	table.DropColumn("primary_ability")
	table.DropColumn("hit_point_die_value")
	table.DropColumn("saves")
	table.MustExec()
}
