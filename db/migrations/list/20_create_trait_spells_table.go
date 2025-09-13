package list

import (
	"dnd-api/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateTraitSpellsTable struct{}

func (m *CreateTraitSpellsTable) GetName() string {
	return "CreateTraitSpellsTable"
}

func (m *CreateTraitSpellsTable) Up(con *sqlx.DB) {
	table := builder.NewTable("trait_spells", con)
	table.Column("id").Type("int unsigned").NotNull().Autoincrement()
	table.PrimaryKey("id")
	table.String("trait_id", 64).NotNull()
	table.Column("spell_id").Type("int unsigned").NotNull()
	table.Integer("uses").Nullable()
	table.Column("reset").Type("ENUM('Long Rest','Short Rest')").Nullable()
	table.MustExec()
}

func (m *CreateTraitSpellsTable) Down(con *sqlx.DB) {
	builder.DropTable("trait_spells", con).MustExec()
}
