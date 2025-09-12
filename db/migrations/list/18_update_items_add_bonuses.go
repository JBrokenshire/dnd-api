package list

import (
	"dnd-api/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type UpdateItemsAddBonuses struct{}

func (m *UpdateItemsAddBonuses) GetName() string {
	return "UpdateItemsAddBonuses"
}

func (m *UpdateItemsAddBonuses) Up(con *sqlx.DB) {
	table := builder.ChangeTable("items", con)
	table.Integer("strength_bonus").NotNull().Default("0")
	table.Integer("dexterity_bonus").NotNull().Default("0")
	table.Integer("constitution_bonus").NotNull().Default("0")
	table.Integer("intelligence_bonus").NotNull().Default("0")
	table.Integer("wisdom_bonus").NotNull().Default("0")
	table.Integer("charisma_bonus").NotNull().Default("0")
	table.Integer("strength_save_bonus").NotNull().Default("0")
	table.Integer("dexterity_save_bonus").NotNull().Default("0")
	table.Integer("constitution_save_bonus").NotNull().Default("0")
	table.Integer("intelligence_save_bonus").NotNull().Default("0")
	table.Integer("wisdom_save_bonus").NotNull().Default("0")
	table.Integer("charisma_save_bonus").NotNull().Default("0")
	table.Column("saving_throw_bonus_type").Type("ENUM('advantage','add')").Nullable()
	table.String("saving_throw_bonus_text", 64).NotNull()
	table.Integer("initiative_bonus").NotNull().Default("0")
	table.Column("initiative_advantage").Type("boolean").NotNull().Default("false")
	table.Integer("armour_class_bonus").NotNull().Default("0")
	table.MustExec()
}

func (m *UpdateItemsAddBonuses) Down(con *sqlx.DB) {
	table := builder.ChangeTable("items", con)
	table.DropColumn("strength_bonus")
	table.DropColumn("dexterity_bonus")
	table.DropColumn("constitution_bonus")
	table.DropColumn("intelligence_bonus")
	table.DropColumn("wisdom_bonus")
	table.DropColumn("charisma_bonus")
	table.DropColumn("strength_save_bonus")
	table.DropColumn("dexterity_save_bonus")
	table.DropColumn("constitution_save_bonus")
	table.DropColumn("intelligence_save_bonus")
	table.DropColumn("wisdom_save_bonus")
	table.DropColumn("charisma_save_bonus")
	table.DropColumn("saving_throw_bonus_type")
	table.Column("saving_throw_bonus_text").NotNull()
	table.DropColumn("initiative_bonus")
	table.DropColumn("initiative_advantage")
	table.DropColumn("armour_class_bonus")
	table.MustExec()
}
