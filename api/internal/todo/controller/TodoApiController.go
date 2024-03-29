package controllers

import (
	"api/internal/api"
	"api/internal/auth"
	"api/internal/todo"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TodoApiController struct {
	todoService todo.TodoService
}

func NewTodoApiController() *TodoApiController {
	return &TodoApiController{
		todoService: *todo.NewTodoService(),
	}
}

// @Summary List todos
// @Description get todos
// @Tags todos
// @Accept  json
// @Produce  json
// @Param   filter query TodoListFilterParams false "Filter the todo list"
// @Success 200 {object} TodosResponse
// @Failure 400 {object} api.ErrorResponse "Invalid request"
// @Failure 401 {object} api.ErrorResponse "Unauthorized"
// @Failure 500 {object} api.ErrorResponse "Internal server error"
// @Router /todos [get]
func (controller *TodoApiController) List(context *gin.Context) {
	var filterParams = TodoListFilterParams{}
	context.ShouldBindQuery(&filterParams)

	authUser := auth.NewAuthUser()
	authUser.Init(context)
	err := filterParams.Validate(authUser.User)

	if err != nil {
		context.IndentedJSON(http.StatusForbidden, api.NewErrorResponse(err))
		return
	}

	todos, total, err := controller.todoService.GetTodos(todo.TodoListFilter(filterParams), 10)
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, api.NewErrorResponse(err))
		return
	}
	context.IndentedJSON(http.StatusOK, ConvertTodoModelsToResponse(todos, total))
}

// @Summary Create a new todo
// @Description create a new todo
// @Tags todos
// @Accept  json
// @Produce  json
// @Param todo body TodoCreatePayload true "todo payload"
// @Success 200 {object} TodoResponse
// @Router /todos [post]
func (controller *TodoApiController) Create(context *gin.Context) {
	var payload TodoCreatePayload
	// Call BindJSON to bind the received JSON to
	if err := context.BindJSON(&payload); err != nil {
		context.IndentedJSON(http.StatusInternalServerError, api.NewErrorResponse(err))
		return
	}

	authUser := auth.NewAuthUser()
	authUser.Init(context)
	err := payload.Validate(authUser.User)
	if err != nil {
		context.IndentedJSON(http.StatusForbidden, api.NewErrorResponse(err))
		return
	}
	todoItem, err := controller.todoService.CreateTodo(todo.TodoFullData(payload))
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, api.NewErrorResponse(err))
		return
	}
	context.IndentedJSON(http.StatusCreated, ConvertTodoModelToResponse(todoItem))
}

// @Summary Get a todo
// @Description get a todo by ID
// @Tags todos
// @Accept  json
// @Produce  json
// @Param id path int true "Todo ID"
// @Success 200 {object} TodoResponse
// @Failure 403 {object} api.ErrorResponse
// @Failure 500 {object} api.ErrorResponse
// @Router /todos/{id} [get]
func (controller *TodoApiController) Read(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, api.NewErrorResponse(err))
		return
	}
	todoItem, err := controller.todoService.FindTodo(uint(id))
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, api.NewErrorResponse(err))
		return
	}
	authUser := auth.NewAuthUser()
	authUser.Init(context)
	err = todoItem.ValidateAccess(authUser.User)
	if err != nil {
		context.IndentedJSON(http.StatusForbidden, api.NewErrorResponse(err))
		return
	}
	context.IndentedJSON(http.StatusOK, ConvertTodoModelToResponse(todoItem))
}

// @Summary Update a todo
// @Description update a todo
// @Tags todos
// @Accept  json
// @Produce  json
// @Param id path int true "Todo ID"
// @Param todo body TodoUpdatePayload true "todo payload"
// @Success 200 {object} TodoResponse
// @Router /todos/{id} [patch]
func (controller *TodoApiController) Update(context *gin.Context) {
	var payload TodoUpdatePayload
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, api.NewErrorResponse(err))
		return
	}

	if err := context.BindJSON(&payload); err != nil {
		context.IndentedJSON(http.StatusInternalServerError, api.NewErrorResponse(err))
		return
	}

	updatedTodo, err := controller.todoService.UpdateTodo(uint(id), todo.TodoUpdatableData(payload))
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, api.NewErrorResponse(err))
		return
	}
	context.IndentedJSON(http.StatusOK, ConvertTodoModelToResponse(updatedTodo))
}

// @Summary Delete a todo
// @Description delete a todo
// @Tags todos
// @Accept  json
// @Produce  json
// @Param id path int true "Todo ID"
// @Success 204 {object} nil
// @Router /todos/{id} [delete]
func (controller *TodoApiController) Delete(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, api.NewErrorResponse(err))
		return
	}

	controller.todoService.DeleteTodo(uint(id))

	context.IndentedJSON(http.StatusNoContent, api.NewSuccessResponse("Todo deleted successfully!"))
}
