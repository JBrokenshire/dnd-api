package list

import (
	"dnd-api/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateFilesTable struct{}

func (m *CreateFilesTable) GetName() string {
	return "CreateFilesTable"
}

func (m *CreateFilesTable) Up(con *sqlx.DB) {
	table := builder.NewTable("files", con)
	table.Column("id").Type("int unsigned").NotNull().Autoincrement()
	table.PrimaryKey("id")
	table.String("model", 100).NotNull()
	table.Column("model_id").Type("int unsigned")
	table.String("filename", 1000)
	table.String("file_location", 1000)

	table.Column("deleted_at").Type("datetime").Nullable()
	table.WithTimestamps()
	table.MustExec()
}

func (m *CreateFilesTable) Down(con *sqlx.DB) {
	builder.DropTable("files", con).MustExec()
}
