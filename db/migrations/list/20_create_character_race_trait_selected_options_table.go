package list

import (
	"dnd-api/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateCharacterSelectedRaceTraitOptionsTable struct{}

func (m *CreateCharacterSelectedRaceTraitOptionsTable) GetName() string {
	return "CreateCharacterSelectedRaceTraitOptionsTable"
}

func (m *CreateCharacterSelectedRaceTraitOptionsTable) Up(con *sqlx.DB) {
	table := builder.NewTable("character_selected_race_trait_options", con)
	table.Column("id").Type("int unsigned").NotNull().Autoincrement()
	table.PrimaryKey("id")
	table.Column("character_id").Type("int unsigned").NotNull()
	table.Column("trait_option_id").Type("int unsigned").NotNull()
	table.MustExec()
}

func (m *CreateCharacterSelectedRaceTraitOptionsTable) Down(con *sqlx.DB) {
	builder.DropTable("character_selected_race_trait_options", con).MustExec()
}
