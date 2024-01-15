package routing

import (
	"github.com/gin-gonic/gin"
)

func Init() {
	router = gin.Default()
	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{"127.0.0.1", "localhost"})
}
