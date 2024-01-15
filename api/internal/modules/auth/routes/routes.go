package routes

import (
	AuthController "api/internal/modules/auth/controllers"
	AuthMiddleware "api/internal/modules/auth/middlewares"

	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine) {
	authController := AuthController.New()
	router.POST("/auth/login", authController.Login)
	router.POST("/auth/logout", authController.Logout)
	router.GET("/auth/me", AuthMiddleware.Authorize(), authController.Me)
}