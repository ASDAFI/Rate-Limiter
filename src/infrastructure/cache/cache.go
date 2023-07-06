package cache

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
	"service/configs"
	"strconv"
)

type Provider struct {
	config configs.CacheConfiguration

	Client *redis.Client
}

var RedisCacheProvider Provider

func createConnection(config configs.CacheConfiguration) (*redis.Client, error) {
	DB, _ := strconv.ParseInt(config.DB, 10, 32)
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port),
		Password: config.Password, // no password set
		DB:       int(DB),         // use default DB
	})

	return client, nil
}

func CreateRedisCacheProvider(config configs.CacheConfiguration) (Provider, error) {

	log.Infof("connection to redis: host=%s port=%d user=%s dbname=%s ", config.Host, config.Port, config.Client, config.DB)
	client, err := createConnection(config)

	if err != nil {
		log.Fatal("Error in redis connection: ", err)
	}

	provider := Provider{config, client}
	log.Info("Creating redis connection has been Done.")

	return provider, nil
}
