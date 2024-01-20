package routes

import (
	authMiddlewares "api/internal/auth/middleware"
	todoControllers "api/internal/todo/controller"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	todoController := todoControllers.NewTodoApiController()
	authMiddleware := authMiddlewares.NewAuthorizeMiddleware()
	group := router.Group("/todos")
	{
		group.Use(authMiddleware.Authorize())
		group.GET("", authMiddleware.EnsureLoggedIn(), todoController.List)
		group.GET("/:id", authMiddleware.EnsureLoggedIn(), todoController.Read)
		group.POST("", authMiddleware.EnsureLoggedIn(), todoController.Create)
		group.PATCH("/:id", authMiddleware.EnsureLoggedIn(), todoController.Update)
		group.DELETE("/:id", authMiddleware.EnsureLoggedIn(), todoController.Delete)
	}

}
