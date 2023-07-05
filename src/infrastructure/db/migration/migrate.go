package migration

import (
	"service/src/infrastructure/db"
)

func MigrateDB() error {

	dbProvider := db.PostgresDBProvider

	err := dbProvider.DB.AutoMigrate()
	if err != nil {
		return err
	}

	return nil
}
