package list

import (
	"github.com/ShkrutDenis/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type UpdateCharactersAddCharacteristics struct{}

func (m *UpdateCharactersAddCharacteristics) GetName() string {
	return "UpdateCharactersAddCharacteristics"
}

func (m *UpdateCharactersAddCharacteristics) Up(con *sqlx.DB) {
	table := builder.ChangeTable("characters", con)
	table.Column("alignment").Type("ENUM('Lawful Good','Neutral Good','Chaotic Good','Lawful Neutral','True Neutral','Chaotic Neutral','Lawful Evil','Neutral Evil','Chaotic Evil')").Nullable()
	table.Column("gender").Type("TINYTEXT").Nullable()
	table.Column("eyes").Type("TINYTEXT").Nullable()
	table.Column("size").Type("ENUM('Tiny','Small','Medium','Large','Huge','Gargantuan')").NotNull()
	table.Column("height").Type("TINYTEXT").Nullable()
	table.Column("faith").Type("TINYTEXT").Nullable()
	table.Column("hair").Type("TINYTEXT").Nullable()
	table.Column("skin").Type("TINYTEXT").Nullable()
	table.Integer("age").Nullable()
	table.Integer("weight").Nullable()
	table.MustExec()
}

func (m *UpdateCharactersAddCharacteristics) Down(con *sqlx.DB) {
	table := builder.ChangeTable("characters", con)
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
