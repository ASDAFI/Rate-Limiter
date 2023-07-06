package migration

import (
	"service/src/infrastructure/db"
	"service/src/users"
)

func MigrateDB() error {

	dbProvider := db.PostgresDBProvider

	err := dbProvider.DB.AutoMigrate(&users.User{})
	if err != nil {
		return err
	}

	return nil
}
