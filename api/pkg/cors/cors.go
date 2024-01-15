package cors

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init() {
	router := gin.Default()
	// config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"http://localhost:3000"}
	router.Use(cors.Default())
}
