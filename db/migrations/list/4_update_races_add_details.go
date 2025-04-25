package list

import (
	"dnd-api/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type UpdateRacesAddDetails struct{}

func (m *UpdateRacesAddDetails) GetName() string {
	return "UpdateRacesAddDetails"
}

func (m *UpdateRacesAddDetails) Up(con *sqlx.DB) {
	table := builder.ChangeTable("races", con)
	table.Column("short_description").Type("text").NotNull()
	table.String("creature_type", 200).NotNull()
	table.String("size", 200).NotNull()
	table.Integer("base_speed").NotNull().Default("30")
	table.MustExec()
}

func (m *UpdateRacesAddDetails) Down(con *sqlx.DB) {
	table := builder.ChangeTable("races", con)
	table.DropColumn("short_description")
	table.DropColumn("creature_type")
	table.DropColumn("size")
	table.DropColumn("base_speed")
	table.MustExec()
}
