package controllers

import (
	GoogleAuthService "api/internal/modules/auth/services"
	UserModel "api/internal/modules/user/models"
	UserRequests "api/internal/modules/user/requests"
	UserResponses "api/internal/modules/user/responses"
	UserService "api/internal/modules/user/services"
	"net/http"

	AuthRequests "api/internal/modules/auth/requests"
	AuthResponses "api/internal/modules/auth/responses"

	AuthToken "api/packages/auth-token"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthController struct {
	googleAuthService GoogleAuthService.GoogleAuthServiceInterface
	userService       UserService.UserServiceInterface
}

func New() *AuthController {
	return &AuthController{
		googleAuthService: GoogleAuthService.New(),
		userService:       UserService.New(),
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

	googleUser, err := controller.googleAuthService.VerifyToken(payload.Token)

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

	authToken := AuthToken.New()
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
	email, emailExists := context.Get("email")

	if !emailExists {
		context.IndentedJSON(http.StatusForbidden, gin.H{"error": "Email not found!"})
		return
	}

	user, err := controller.userService.FindUserByEmail(email.(string))
	if err == gorm.ErrRecordNotFound {
		context.IndentedJSON(http.StatusForbidden, gin.H{"error": "User not found!"})
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