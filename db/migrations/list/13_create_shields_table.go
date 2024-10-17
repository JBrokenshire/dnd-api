package list

import (
	"github.com/ShkrutDenis/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateShieldsTable struct{}

func (m *CreateShieldsTable) GetName() string { return "CreateShieldsTable" }

func (m *CreateShieldsTable) Up(con *sqlx.DB) {
	table := builder.NewTable("shields", con)
	table.Integer("item_id").Unique().NotNull().NotAutoincrement()
	table.PrimaryKey("item_id")
	table.ForeignKey("item_id").Reference("items").On("id").OnUpdate("cascade").OnDelete("cascade")
	table.Integer("bonus_ac").NotNull().Default("2")
	table.MustExec()
}

func (m *CreateShieldsTable) Down(con *sqlx.DB) {
	builder.DropTable("shields", con).MustExec()
}
