package database

import (
	"github.com/go-redis/redis"
	"github.com/auth-core/config"
)

var client *redis.Client

func RedisOpen() *redis.Client {

	serverMode := config.MustGetString("server.mode")
	redisHost := config.MustGetString(serverMode + ".redis_host")
	redisPort := config.MustGetString(serverMode + ".redis_port")
	// redisUsername := config.MustGetString(serverMode + ".redis_username")
	redisPassword := config.MustGetString(serverMode + ".redis_password")

	client := redis.NewClient(&redis.Options{
		Addr:     "" + redisHost + ":" + redisPort + "",
		Password: redisPassword, // no password set
		DB:       0,             // use default DB
	})

	return client
}
