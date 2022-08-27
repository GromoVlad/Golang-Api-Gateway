package migration

import (
	_ "database/sql"
	"gin_tonic/internal/database/DB"
	_ "github.com/lib/pq"
	"log"
)

type Migration struct {
	Id        int    `db:"id"`
	Timestamp int    `db:"timestamp"`
	Name      string `db:"name"`
}

func FindAllMigration() []Migration {
	var migration []Migration
	err := DB.Connect().Select(&migration, "SELECT * FROM migrations")
	if err != nil {
		log.Fatalln(err)
	}
	return migration
}
