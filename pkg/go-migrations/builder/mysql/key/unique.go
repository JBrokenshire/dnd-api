package key

import (
	"dnd-api/pkg/go-migrations/builder/contract"
	"dnd-api/pkg/go-migrations/config"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

type UniqueKey struct {
	name  string
	table string
	field string

	drop bool
}

func NewUniqueKey(table, field string) contract.UniqueKey {
	uk := &UniqueKey{table: table, field: field}
	return uk
}

func (uk *UniqueKey) SetKeyName(name string) contract.UniqueKey {
	uk.name = name
	return uk
}

func (uk *UniqueKey) GenerateKeyName() contract.UniqueKey {
	uk.name = fmt.Sprintf("%v_%v_uindex", uk.table, uk.field)
	return uk
}

func (uk *UniqueKey) Drop() contract.UniqueKey {
	uk.drop = true
	return uk
}

func (uk *UniqueKey) GetSQL() string {
	if uk.drop {
		return fmt.Sprintf("Drop index %v on %v;",
			uk.name, uk.table)
	}
	return fmt.Sprintf("Create unique index %v on %v (%v);",
		uk.name, uk.table, uk.field)
}

func (uk *UniqueKey) Exec(con *sqlx.DB) error {
	if config.GetConfig().Verbose {
		log.Println(uk.GetSQL())
	}
	_, err := con.Exec(uk.GetSQL())
	return err
}

func (uk *UniqueKey) MustExec(con *sqlx.DB) {
	if config.GetConfig().Verbose {
		log.Println(uk.GetSQL())
	}
	con.MustExec(uk.GetSQL())
}

func (uk *UniqueKey) GetName() string {
	return uk.name
}
