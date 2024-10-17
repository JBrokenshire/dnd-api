package list

import (
	"github.com/ShkrutDenis/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type UpdateCharactersAddBackground struct{}

func (m *UpdateCharactersAddBackground) GetName() string {
	return "UpdateCharactersAddBackground"
}

func (m *UpdateCharactersAddBackground) Up(con *sqlx.DB) {
	table := builder.ChangeTable("characters", con)
	table.String("background_name", 64).NotNull()
	table.MustExec()
}

func (m *UpdateCharactersAddBackground) Down(con *sqlx.DB) {
	table := builder.ChangeTable("characters", con)
	table.DropColumn("background_name")
	table.MustExec()
}
