package routes

import (
	AuthController "api/internal/modules/auth/controllers"
	AuthMiddleware "api/internal/modules/auth/middlewares"

	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine) {
	authController := AuthController.New()
	router.POST("/user/login", authController.Login)
	router.POST("/user/logout", authController.Logout)
	router.GET("/user/me", AuthMiddleware.Authorize(), authController.Me)
}
