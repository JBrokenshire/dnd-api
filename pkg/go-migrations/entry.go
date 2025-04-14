package go_migrations

import (
	"dnd-api/pkg/go-migrations/config"
	"dnd-api/pkg/go-migrations/db"
	"dnd-api/pkg/go-migrations/model"
	"dnd-api/pkg/go-migrations/store"
	"flag"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"log"
)

var connection *sqlx.DB

func Run(migs []store.Migratable) {
	parseFlags()
	prepare()

	store.RegisterMigrations(migs)

	if config.GetConfig().IsRollback {
		rollBack()
		return
	}

	if !config.GetConfig().FirstRun {
		//ok, m := model.GetLastMigration(connection)
		ok, m := model.GetAllMigrations(connection)
		if ok {
			store.FilterToMissingMigrations(m)
		}
	}
	up()
}

func rollBack() {
	forRollback := model.GetLastMigrations(connection, config.GetConfig().LastBatch)
	for _, m := range forRollback {
		forRun := store.FindMigration(m.Name)
		if forRun == nil {
			log.Fatal("Migration", m.Name, "not found")
		}
		log.Println("Rolling back", forRun.GetName())
		forRun.Down(connection)
		log.Println("Rolled back", forRun.GetName())
	}
	model.RemoveLastBatch(connection, config.GetConfig().LastBatch)
}

func up() {
	for _, m := range store.GetMigrationsList() {
		log.Println("Migrating", m.GetName())
		m.Up(connection)
		model.AddMigrationRaw(connection, m.GetName(), config.GetConfig().LastBatch+1)
		log.Println("Migrated", m.GetName())
	}
}

func parseFlags() {
	isRollback := flag.Bool("rollback", false, "Flag for init rollback.")
	envPath := flag.String("env-path", "", "Path to .env file.")
	envFile := flag.String("env-file", ".env", "Env file name.")
	verbose := flag.Bool("verbose", false, "Flag for show more info.")
	flag.Parse()
	config.GetConfig().IsRollback = *isRollback
	config.GetConfig().EnvPath = *envPath
	config.GetConfig().EnvFile = *envFile
	config.GetConfig().Verbose = *verbose
}

func prepare() {
	if config.GetConfig().Verbose {
		log.Println("load env file from:", config.GetConfig().GetEnvFullPath())
	}
	err := godotenv.Load(config.GetConfig().GetEnvFullPath())
	if err != nil {
		log.Println("Error loading .env file")
	}

	connector := db.NewConnector()
	connection, err = connector.Connect()
	if err != nil {
		log.Fatal(err)
	}
	config.GetConfig().LastBatch = model.GetLastBatch(connection)
	config.GetConfig().FirstRun = model.CreateMigrationsTable(connection)
}
