package list

import (
	"github.com/ShkrutDenis/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type UpdateCharactersAddSubclassID struct{}

func (m *UpdateCharactersAddSubclassID) GetName() string {
	return "UpdateCharactersAddSubclassID"
}

func (m *UpdateCharactersAddSubclassID) Up(con *sqlx.DB) {
	table := builder.ChangeTable("characters", con)
	table.Integer("subclass_id").Nullable()
	table.ForeignKey("subclass_id").Reference("subclasses").On("id").OnUpdate("cascade").OnDelete("cascade")
	table.MustExec()
}

func (m *UpdateCharactersAddSubclassID) Down(con *sqlx.DB) {
	table := builder.ChangeTable("characters", con)
	table.DropColumn("subclass_id")
	table.MustExec()
}
