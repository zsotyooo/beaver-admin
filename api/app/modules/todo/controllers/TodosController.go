package controllers

import (
	TodoRequest "api/app/modules/todo/requests"
	TodoService "api/app/modules/todo/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TodosController struct {
	todoService TodoService.TodoServiceInterface
}

func New() *TodosController {
	return &TodosController{
		todoService: TodoService.New(),
	}
}

func (controller *TodosController) List(c *gin.Context) {
	todos, err := controller.todoService.GetTodos(10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, todos)
}

func (controller *TodosController) Create(c *gin.Context) {
	var newTodo TodoRequest.TodoCreateRequest
	// Call BindJSON to bind the received JSON to
	if err := c.BindJSON(&newTodo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	todo, err := controller.todoService.CreateTodo(newTodo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, todo)
}

func (controller *TodosController) Read(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	todo, err := controller.todoService.FindTodo(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, todo)
}

func (controller *TodosController) Update(c *gin.Context) {
	var newTodo TodoRequest.TodoUpdateRequest
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&newTodo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	updatedTodo, err := controller.todoService.UpdateTodo(uint(id), newTodo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, updatedTodo)
}

func (controller *TodosController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	controller.todoService.DeleteTodo(uint(id))

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Todo deleted successfully!"})
}
