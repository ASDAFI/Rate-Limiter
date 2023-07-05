package cli

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"service/configs"
)

var configFile string

func initConfig() {
	viper.SetConfigName(configFile)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Fatal error loading config file: %s \n", err)
	}
	if err := viper.Unmarshal(&configs.Config); err != nil {
		log.Fatalf("Fatal error marshalng config file: %s \n", err)

	}

	log.Info("Configuration initialized!")

	if configs.Config.Credential.TokenSecret == "" {
		log.Fatal("There is no token secret in config file\n")
	}

	//dbProvider, err := infrastructure.CreateDBProvider(configs.Config.Database)
	//if err != nil {
	//	log.Fatalf("Fatal error on create db: %s \n", err)
	//}
	//infrastructure.PostgresDBProvider = dbProvider

	log.Info("Configuration initialized!")

	if configs.Config.Credential.TokenSecret == "" {
		log.Fatal("There is no token secret in config file\n")
	}

	//dbProvider, err := infrastructure.CreateDBProvider(configs.Config.Database)
	//if err != nil {
	//	log.Fatalf("Fatal error on create db: %s \n", err)
	//}
	//infrastructure.PostgresDBProvider = dbProvider
}
