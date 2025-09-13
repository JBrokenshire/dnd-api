package list

import (
	"dnd-api/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateRaceTraitsTable struct{}

func (m *CreateRaceTraitsTable) GetName() string {
	return "CreateRaceTraitsTable"
}

func (m *CreateRaceTraitsTable) Up(con *sqlx.DB) {
	table := builder.NewTable("race_traits", con)
	table.Column("id").Type("int unsigned").NotNull().Autoincrement()
	table.PrimaryKey("id")
	table.Column("race_id").Type("int unsigned").NotNull()
	table.String("trait_id", 64).NotNull()
	table.MustExec()
}

func (m *CreateRaceTraitsTable) Down(con *sqlx.DB) {
	builder.DropTable("race_traits", con).MustExec()
}
