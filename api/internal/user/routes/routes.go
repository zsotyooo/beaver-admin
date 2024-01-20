package routes

import (
	authMiddlewares "api/internal/auth/middleware"
	userControllers "api/internal/user/controller"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	userController := userControllers.NewUserApiController()
	authMiddleware := authMiddlewares.NewAuthorizeMiddleware()
	group := router.Group("/users")
	{
		group.Use(authMiddleware.Authorize())
		group.GET("", authMiddleware.EnsureLoggedIn(), userController.List)
		group.GET("/:id", authMiddleware.EnsureLoggedIn(), userController.Read)
	}

}
