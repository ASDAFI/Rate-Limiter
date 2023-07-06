package cache

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
	"service/configs"
)

type Provider struct {
	config configs.CacheConfiguration

	Client *redis.Client
}

var RedisCacheProvider Provider

func getConnectionString(config configs.CacheConfiguration) string {
	connectionString := fmt.Sprintf("redis://%s:%s@%s:%d/%s", config.Client, config.Password,
		config.Host, config.Port, config.DB)
	return connectionString
}

func createConnection(connectionString string) (*redis.Client, error) {
	opt, err := redis.ParseURL(connectionString)
	if err != nil {
		return nil, err
	}
	client := redis.NewClient(opt)
	return client, nil
}

func CreateRedisCacheProvider(config configs.CacheConfiguration) (Provider, error) {

	log.Infof("connection to redis: host=%s port=%d user=%s dbname=%s ", config.Host, config.Port, config.Client, config.DB)
	connectionString := getConnectionString(config)
	client, err := createConnection(connectionString)

	if err != nil {
		log.Fatal("Error in redis connection: ", err)
	}

	provider := Provider{config, client}
	log.Info("Creating redis connection has been Done.")

	return provider, nil
}
