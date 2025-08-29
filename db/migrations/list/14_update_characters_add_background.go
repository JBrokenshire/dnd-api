package list

import (
	"dnd-api/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type UpdateCharactersAddBackground struct{}

func (m *UpdateCharactersAddBackground) GetName() string {
	return "UpdateCharactersAddBackground"
}

func (m *UpdateCharactersAddBackground) Up(con *sqlx.DB) {
	table := builder.ChangeTable("characters", con)
	table.Column("background_id").Type("int unsigned").Default("0")
	table.String("alignment", 32).Nullable()
	table.String("gender", 32).Nullable()
	table.String("eyes", 32).Nullable()
	table.String("size", 32).Nullable()
	table.String("height", 32).Nullable()
	table.String("faith", 32).Nullable()
	table.String("hair", 32).Nullable()
	table.String("skin", 32).Nullable()
	table.String("age", 32).Nullable()
	table.String("weight", 32).Nullable()
	table.MustExec()
}

func (m *UpdateCharactersAddBackground) Down(con *sqlx.DB) {
	table := builder.ChangeTable("characters", con)
	table.DropColumn("background_id")
	table.DropColumn("alignment")
	table.DropColumn("gender")
	table.DropColumn("eyes")
	table.DropColumn("size")
	table.DropColumn("height")
	table.DropColumn("faith")
	table.DropColumn("hair")
	table.DropColumn("skin")
	table.DropColumn("age")
	table.DropColumn("weight")
	table.MustExec()
}
