package list

import (
	"github.com/ShkrutDenis/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type UpdateCharacterAddNotes struct{}

func (m *UpdateCharacterAddNotes) GetName() string {
	return "UpdateCharacterAddNotes"
}

func (m *UpdateCharacterAddNotes) Up(con *sqlx.DB) {
	table := builder.ChangeTable("characters", con)
	table.Column("organisations").Type("MEDIUMTEXT").Nullable()
	table.Column("allies").Type("MEDIUMTEXT").Nullable()
	table.Column("enemies").Type("MEDIUMTEXT").Nullable()
	table.Column("backstory").Type("LONGTEXT").Nullable()
	table.MustExec()
}

func (m *UpdateCharacterAddNotes) Down(con *sqlx.DB) {
	table := builder.ChangeTable("characters", con)
	table.DropColumn("organisations")
	table.DropColumn("allies")
	table.DropColumn("enemies")
	table.DropColumn("backstory")
	table.MustExec()
}
