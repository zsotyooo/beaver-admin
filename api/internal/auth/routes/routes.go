package routes

import (
	controllers "api/internal/auth/controller"
	authMiddlewares "api/internal/auth/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	authController := controllers.NewAuthApiController()
	authMiddleware := authMiddlewares.NewAuthorizeMiddleware()
	group := router.Group("/auth")
	{
		group.Use(authMiddleware.Authorize())
		group.POST("/login", authController.Login)
		group.POST("/logout", authController.Logout)
		group.GET("/me", authMiddleware.EnsureLoggedIn(), authController.Me)
	}

}
