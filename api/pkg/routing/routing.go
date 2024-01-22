package routing

import (
	"api/pkg/config"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func Init() {
	router = gin.Default()
	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{"127.0.0.1", "localhost"})
}

func GetRouter() *gin.Engine {
	return router
}

func Serve() {
	router := GetRouter()

	configs := config.Get()

	err := router.Run(fmt.Sprintf("%s:%s", configs.Server.Host, configs.Server.Port))

	if err != nil {
		log.Fatal("Error in routing")
		return
	}
}
