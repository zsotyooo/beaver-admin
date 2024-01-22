package database

import (
	"os"

	"github.com/boj/redistore"
)

var RediStore *redistore.RediStore

func RedisConnect() {
	redisHost := os.Getenv("REDIS_HOST")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisPort := os.Getenv("REDIS_PORT")
	redisKey := os.Getenv("REDIS_KEY")

	store, err := redistore.NewRediStore(10, "tcp", redisHost+":"+redisPort, redisPassword, []byte(redisKey))
	if err != nil {
		panic(err)
	}
	RediStore = store
}

func RedisStore() *redistore.RediStore {
	return RediStore
}
