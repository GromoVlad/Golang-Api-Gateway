package migration

import (
	"gin_tonic/internal/pgdb"
	"log"
)

type Migration struct {
	Id        int    `db:"id"`
	Timestamp int    `db:"timestamp"`
	Name      string `db:"name"`
}

func FindAllMigration() []Migration {
	var migration []Migration
	err := pgdb.DB().Select(&migration, "SELECT * FROM migrations")
	if err != nil {
		log.Fatalln(err)
	}
	return migration
}
