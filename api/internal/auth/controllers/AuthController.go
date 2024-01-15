package controllers

import (
	UserModel "api/internal/user/models"
	UserRequests "api/internal/user/requests"
	UserResponses "api/internal/user/responses"
	UserService "api/internal/user/services"
	"api/pkg/googleauth"
	"net/http"

	AuthRequests "api/internal/auth/requests"
	AuthResponses "api/internal/auth/responses"

	"api/pkg/jwtauth"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthController struct {
	userService UserService.UserServiceInterface
}

func New() *AuthController {
	return &AuthController{
		userService: UserService.New(),
	}
}

// @Summary Log in via google oauth token
// @Description log in with token
// @Tags auth
// @Accept  json
// @Produce  json
// @Param payload body requests.LoginPayload true "Login payload"
// @Success 200 {object} responses.LoginResponse
// @Failure 400 {object} responses.ErrorResponse
// @Router /auth/login [post]
func (controller *AuthController) Login(context *gin.Context) {
	var payload AuthRequests.LoginPayload
	var user UserResponses.UserResponse

	if err := context.BindJSON(&payload); err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	googleUser, err := googleauth.VerifyToken(payload.Token)

	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userExists, err := controller.userService.UserExists(googleUser.Email)

	if userExists {
		user, err = controller.userService.FindUserByEmail(googleUser.Email)
	} else {
		userRole, _ := UserModel.UserUserRole.Value()
		user, err = controller.userService.CreateUser(UserRequests.UserCreatePayload{
			Email: googleUser.Email,
			Name:  googleUser.Name,
			Role:  userRole.(string),
		})
	}

	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	authToken := jwtauth.New()
	token, err := authToken.GenerateToken(user.Email)

	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.SetCookie("authToken", token, 3600, "/", "", false, true)

	context.IndentedJSON(http.StatusOK, AuthResponses.CreateLoginResponse(token, user))
}

// @Summary Get current user
// @Description get current user
// @Tags auth
// @Accept  json
// @Produce  json
// @Success 200 {object} responses.UserResponse
// @Failure 204 {object} responses.ErrorResponse
// @Router /auth/me [get]
func (controller *AuthController) Me(context *gin.Context) {
	email, authenticated := context.Get("email")

	if !authenticated {
		context.IndentedJSON(http.StatusNoContent, gin.H{"message": "Not authenticated!"})
		return
	}

	user, err := controller.userService.FindUserByEmail(email.(string))
	if err == gorm.ErrRecordNotFound {
		context.IndentedJSON(http.StatusForbidden, gin.H{"error": "User not found with the given email!"})
		return
	}

	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.IndentedJSON(http.StatusOK, user)
}

// @Summary Log out
// @Description log out
// @Tags auth
// @Accept  json
// @Produce  json
// @Success 200 {object} responses.SuccessResponse
// @Router /auth/logout [post]
func (controller *AuthController) Logout(context *gin.Context) {
	context.SetCookie("authToken", "", -1, "/", "", false, true)
	context.IndentedJSON(http.StatusOK, gin.H{"message": "Logout successful!"})
}
