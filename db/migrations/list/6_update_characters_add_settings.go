package list

import (
	"dnd-api/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type UpdateCharactersAddSettings struct{}

func (m *UpdateCharactersAddSettings) GetName() string {
	return "UpdateCharactersAddSettings"
}

func (m *UpdateCharactersAddSettings) Up(con *sqlx.DB) {
	table := builder.ChangeTable("characters", con)
	table.Column("advancement_type").Type("ENUM('Milestone', 'XP')").NotNull().Default("Milestone")
	table.Column("hit_point_type").Type("ENUM('Fixed','Manual')").NotNull().Default("Fixed")
	table.MustExec()
}

func (m *UpdateCharactersAddSettings) Down(con *sqlx.DB) {
	table := builder.ChangeTable("characters", con)
	table.DropColumn("advancement_type")
	table.DropColumn("hit_point_type")
	table.MustExec()
}
