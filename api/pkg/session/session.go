package session

import (
	"api/pkg/routing"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func InitSession() {
	router := routing.GetRouter()
	redisHost := os.Getenv("REDIS_HOST")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisPort := os.Getenv("REDIS_PORT")
	redisKey := os.Getenv("REDIS_KEY")
	store, err := redis.NewStore(10, "tcp", redisHost+":"+redisPort, redisPassword, []byte(redisKey))
	if err != nil {
		panic(err)
	}

	router.Use(sessions.Sessions("session", store))
}

func Set[T interface{}](c *gin.Context, key string, value T) error {
	session := sessions.Default(c)

	session.Set(key, value)
	return session.Save()
}

func Flash[T interface{}](c *gin.Context, key string) (val T, hasVal bool, err error) {
	session := sessions.Default(c)
	hasVal = false

	response := session.Get(key)
	err = session.Save()
	if err != nil {
		return
	}
	session.Delete(key)
	err = session.Save()
	if err != nil {
		return
	}

	if response != nil {
		val = response.(T)
		hasVal = true
	}

	return
}

func Get[T interface{}](c *gin.Context, key string) (val T, hasVal bool) {
	session := sessions.Default(c)
	hasVal = false
	response := session.Get(key)
	session.Save()

	if response != nil {
		val = response.(T)
		hasVal = true
	}

	return
}

func Remove(c *gin.Context, key string) error {
	session := sessions.Default(c)

	session.Delete(key)
	return session.Save()
}
