package routes

import (
	TodosController "api/internal/todo/controllers"

	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine) {
	todosController := TodosController.New()
	router.GET("/todos", todosController.List)
	router.GET("/todos/:id", todosController.Read)
	router.POST("/todos", todosController.Create)
	router.PATCH("/todos/:id", todosController.Update)
	router.DELETE("/todos/:id", todosController.Delete)
}
