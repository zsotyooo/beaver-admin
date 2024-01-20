package routes

import (
	"api/pkg/swagger"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/swagger/*any", swagger.NewSwaggerDocsController())
}
