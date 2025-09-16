package list

import (
	"dnd-api/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateCharacterSelectedClassFeatureOptionsTable struct{}

func (m *CreateCharacterSelectedClassFeatureOptionsTable) GetName() string {
	return "CreateCharacterSelectedClassFeatureOptionsTable"
}

func (m *CreateCharacterSelectedClassFeatureOptionsTable) Up(con *sqlx.DB) {
	table := builder.NewTable("character_selected_class_feature_options", con)
	table.Column("id").Type("int unsigned").NotNull().Autoincrement()
	table.PrimaryKey("id")
	table.Column("character_id").Type("int unsigned").NotNull()
	table.Column("class_feature_option_id").Type("int unsigned").NotNull()
	table.MustExec()
}

func (m *CreateCharacterSelectedClassFeatureOptionsTable) Down(con *sqlx.DB) {
	builder.DropTable("character_selected_clas_feature_options", con).MustExec()
}
