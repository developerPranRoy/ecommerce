package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	migrate "github.com/rubenv/sql-migrate"
)

func MigrateDB(db *sqlx.DB, dir string) error {
	migrationss := &migrate.FileMigrationSource{
		Dir: dir,
	}
	_, err := migrate.Exec(db.DB, "postgres", migrationss, migrate.Up)
	if err != nil {
		return err
	}
	fmt.Println("Succesfully migrated DB")
	return nil
}
