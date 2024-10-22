package list

import (
	"github.com/ShkrutDenis/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateRaceTraitsTable struct{}

func (m *CreateRaceTraitsTable) GetName() string {
	return "CreateRaceTraitsTable"
}

func (m *CreateRaceTraitsTable) Up(con *sqlx.DB) {
	table := builder.NewTable("race_traits", con)
	table.PrimaryKey("id")
	table.Integer("race_id").NotNull()
	table.ForeignKey("race_id").Reference("races").On("id").OnUpdate("cascade").OnDelete("cascade")
	table.Integer("trait_id").NotNull()
	table.ForeignKey("trait_id").Reference("traits").On("id").OnUpdate("cascade").OnDelete("cascade")
	table.MustExec()
}

func (m *CreateRaceTraitsTable) Down(con *sqlx.DB) {
	builder.DropTable("race_traits", con).MustExec()
}
