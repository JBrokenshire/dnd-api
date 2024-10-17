package list

import (
	"github.com/ShkrutDenis/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateCharacterSkillsAdvantagesTable struct{}

func (m *CreateCharacterSkillsAdvantagesTable) GetName() string {
	return "CreateCharacterSkillAdvantagesTable"
}

func (m *CreateCharacterSkillsAdvantagesTable) Up(con *sqlx.DB) {
	table := builder.NewTable("character_skills_advantages", con)
	table.PrimaryKey("id")
	table.Integer("character_id").NotNull()
	table.Column("skill_name").Type("ENUM('Acrobatics','Animal Handling','Arcana','Athletics','Deception','History','Insight','Intimidation','Investigation','Medicine','Nature','Perception','Performance','Persuasion','Religion','Sleight of Hand','Stealth','Survival')").NotNull()
	table.Column("advantage").Type("BOOLEAN").NotNull().Default("0")
	table.Column("disadvantage").Type("BOOLEAN").NotNull().Default("0")
	table.MustExec()
}

func (m *CreateCharacterSkillsAdvantagesTable) Down(con *sqlx.DB) {
	builder.DropTable("character_skills_advantages", con).MustExec()
}
