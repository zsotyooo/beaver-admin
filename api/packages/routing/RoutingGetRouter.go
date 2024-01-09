package routing

import (
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	return router
}
