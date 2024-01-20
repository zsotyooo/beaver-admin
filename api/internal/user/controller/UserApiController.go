package controllers

import (
	"api/internal/api"
	"api/internal/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserApiController struct {
	userService user.UserService
}

func NewUserApiController() *UserApiController {
	return &UserApiController{
		userService: *user.NewUserService(),
	}
}

// @Summary List users
// @Description get users
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {object} user.UsersResponse
// @Failure 400 {object} api.ErrorResponse "Invalid request"
// @Failure 401 {object} api.ErrorResponse "Unauthorized"
// @Failure 500 {object} api.ErrorResponse "Internal server error"
// @Router /users [get]
func (controller *UserApiController) List(context *gin.Context) {
	users, total, err := controller.userService.GetUsers(10)
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, api.NewErrorResponse(err))
		return
	}
	context.IndentedJSON(http.StatusOK, user.ConvertModelsToResponse(users, total))
}

// @Summary Get a user
// @Description get a user by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} user.UserResponse
// @Failure 403 {object} api.ErrorResponse
// @Failure 500 {object} api.ErrorResponse
// @Router /users/{id} [get]
func (controller *UserApiController) Read(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, api.NewErrorResponse(err))
		return
	}
	userItem, err := controller.userService.GetUserById(uint(id))
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, api.NewErrorResponse(err))
		return
	}

	context.IndentedJSON(http.StatusOK, user.ConvertModelToResponse(userItem))
}
