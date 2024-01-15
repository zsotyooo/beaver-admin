package routes

import (
	"api/pkg/swagger"

	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine) {
	router.GET("/swagger/*any", swagger.NewSwaggerDocsController())
}
