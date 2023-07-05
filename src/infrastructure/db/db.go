package db

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"service/configs"
	"time"
)

type Provider struct {
	config configs.DatabaseConfiguration

	DB *gorm.DB
}

var PostgresDBProvider Provider

func getConnectionString(config configs.DatabaseConfiguration) string {
	connectionString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, config.Password, config.DB)
	return connectionString
}

func createConnection(connectionString string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  connectionString,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	return db, err
}

func CreateDBProvider(config configs.DatabaseConfiguration) (Provider, error) {

	connectionString := getConnectionString(config)
	log.Info("host=%s port=%d user=%s dbname=%s ", config.Host, config.Port, config.User, config.DB)

	db, err := createConnection(connectionString)
	provider := &Provider{
		config: config,
		DB:     db,
	}

	if err != nil {
		log.Fatal("Error in Create db: ", err)
		return *provider, err
	}

	db1, err := db.DB()
	if err != nil {
		log.Info("Error in Create db: ", err)
		return *provider, err
	}
	db1.SetConnMaxLifetime(time.Minute * time.Duration(config.ConnectionMaxLifetime))
	db1.SetMaxIdleConns(config.MaxIdleConnections)
	db1.SetMaxOpenConns(config.MaxOpenConnections)

	log.Info("Create Db Done.")
	return *provider, nil
}
