package controllers

import (
	TodoRequest "api/internal/todo/requests"
	TodoService "api/internal/todo/services"
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

// @Summary List todos
// @Description get todos
// @Tags todos
// @Accept  json
// @Produce  json
// @Param limit query int false "limit"
// @Success 200 {object} responses.TodosResponse
// @Router /todos [get]
func (controller *TodosController) List(c *gin.Context) {
	todos, err := controller.todoService.GetTodos(10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, todos)
}

// @Summary Create a new todo
// @Description create a new todo
// @Tags todos
// @Accept  json
// @Produce  json
// @Param todo body requests.TodoCreatePayload true "todo payload"
// @Success 200 {object} responses.TodoResponse
// @Router /todos [post]
func (controller *TodosController) Create(c *gin.Context) {
	var payload TodoRequest.TodoCreatePayload
	// Call BindJSON to bind the received JSON to
	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	todo, err := controller.todoService.CreateTodo(payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, todo)
}

// @Summary Create a new todo
// @Description create a new todo
// @Tags todos
// @Accept  json
// @Produce  json
// @Param todo body requests.TodoCreatePayload true "todo payload"
// @Success 200 {object} responses.TodoResponse
// @Router /todos [post]
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

// @Summary Update a todo
// @Description update a todo
// @Tags todos
// @Accept  json
// @Produce  json
// @Param id path int true "Todo ID"
// @Param todo body requests.TodoUpdatePayload true "todo payload"
// @Success 200 {object} responses.TodoResponse
// @Router /todos/{id} [patch]
func (controller *TodosController) Update(c *gin.Context) {
	var payload TodoRequest.TodoUpdatePayload
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	updatedTodo, err := controller.todoService.UpdateTodo(uint(id), payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, updatedTodo)
}

// @Summary Delete a todo
// @Description delete a todo
// @Tags todos
// @Accept  json
// @Produce  json
// @Param id path int true "Todo ID"
// @Success 204 {object} nil
// @Router /todos/{id} [delete]
func (controller *TodosController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	controller.todoService.DeleteTodo(uint(id))

	c.IndentedJSON(http.StatusNoContent, gin.H{"message": "Todo deleted successfully!"})
}
