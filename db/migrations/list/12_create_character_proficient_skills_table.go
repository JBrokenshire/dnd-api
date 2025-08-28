package list

import (
	"dnd-api/pkg/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateCharacterProficientSkillsTable struct{}

func (m *CreateCharacterProficientSkillsTable) GetName() string {
	return "CreateCharacterProficientSkillsTable"
}

func (m *CreateCharacterProficientSkillsTable) Up(con *sqlx.DB) {
	table := builder.NewTable("character_proficient_skills", con)
	table.Column("id").Type("int unsigned").NotNull().Autoincrement()
	table.PrimaryKey("id")
	table.Column("character_id").Type("int unsigned").NotNull()
	table.Column("skill").Type("ENUM('Acrobatics','Animal Handling','Arcana','Athletics','Deception','History','Insight','Intimidation','Investigation','Medicine','Nature','Perception','Performance','Persuasion','Religion','Sleight of Hand','Stealth','Survival')").NotNull()
	table.Column("proficiency_type").Type("ENUM('Half Proficiency', 'Proficiency', 'Expertise')").NotNull().Default("Proficiency")
	table.MustExec()
}

func (m *CreateCharacterProficientSkillsTable) Down(con *sqlx.DB) {
	builder.DropTable("character_proficient_skills", con).MustExec()
}
