package list

import (
	"dnd-api/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type UpdateCharactersAddNotes struct{}

func (m *UpdateCharactersAddNotes) GetName() string {
	return "UpdateCharactersAddNotes"
}

func (m *UpdateCharactersAddNotes) Up(con *sqlx.DB) {
	table := builder.ChangeTable("characters", con)
	table.Column("personality_traits").Type("LONGTEXT").NotNull()
	table.Column("ideals").Type("LONGTEXT").NotNull()
	table.Column("bonds").Type("LONGTEXT").NotNull()
	table.Column("flaws").Type("LONGTEXT").NotNull()
	table.Column("organisations").Type("LONGTEXT").NotNull()
	table.Column("allies").Type("LONGTEXT").NotNull()
	table.Column("enemies").Type("LONGTEXT").NotNull()
	table.Column("backstory").Type("LONGTEXT").NotNull()
	table.MustExec()
}

func (m *UpdateCharactersAddNotes) Down(con *sqlx.DB) {
	table := builder.ChangeTable("characters", con)
	table.DropColumn("personality_traits")
	table.DropColumn("ideals")
	table.DropColumn("bonds")
	table.DropColumn("flaws")
	table.DropColumn("organisations")
	table.DropColumn("allies")
	table.DropColumn("enemies")
	table.DropColumn("backstory")
	table.MustExec()
}
