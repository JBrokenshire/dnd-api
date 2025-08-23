package list

import (
	"dnd-api/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type UpdateCharactersAddSavingThrowAdjustments struct{}

func (m *UpdateCharactersAddSavingThrowAdjustments) GetName() string {
	return "UpdateCharactersAddSavingThrowAdjustments"
}

func (m *UpdateCharactersAddSavingThrowAdjustments) Up(con *sqlx.DB) {
	table := builder.ChangeTable("characters", con)
	table.Integer("strength_save_adjustment").Default("0").NotNull()
	table.Integer("dexterity_save_adjustment").Default("0").NotNull()
	table.Integer("constitution_save_adjustment").Default("0").NotNull()
	table.Integer("intelligence_save_adjustment").Default("0").NotNull()
	table.Integer("wisdom_save_adjustment").Default("0").NotNull()
	table.Integer("charisma_save_adjustment").Default("0").NotNull()
	table.MustExec()
}

func (m *UpdateCharactersAddSavingThrowAdjustments) Down(con *sqlx.DB) {
	table := builder.ChangeTable("characters", con)
	table.DropColumn("strength_save_adjustment")
	table.DropColumn("dexterity_save_adjustment")
	table.DropColumn("constitution_save_adjustment")
	table.DropColumn("intelligence_save_adjustment")
	table.DropColumn("wisdom_save_adjustment")
	table.DropColumn("charisma_save_adjustment")
	table.MustExec()
}
