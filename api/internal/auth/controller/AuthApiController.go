package controllers

import (
	"api/internal/api"
	"api/internal/auth"
	"api/internal/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthApiController struct {
	authService auth.AuthService
	userService user.UserService
}

func NewAuthApiController() *AuthApiController {
	return &AuthApiController{
		authService: *auth.NewAuthService(),
		userService: *user.NewUserService(),
	}
}

// @Summary Log in via google oauth token
// @Description log in with token
// @Tags auth
// @Accept  json
// @Produce  json
// @Param payload body auth.LoginPayload true "Login payload"
// @Success 200 {object} auth.LoginResponse
// @Failure 400 {object} api.ErrorResponse
// @Router /auth/login [post]
func (controller *AuthApiController) Login(context *gin.Context) {
	var payload auth.LoginPayload

	if err := context.BindJSON(&payload); err != nil {
		err = auth.ErrorBadLoginCredentials
		context.AbortWithStatusJSON(auth.GetHttpStatusCode(err), api.NewErrorResponse(err))
		return
	}

	token, authUser, err := controller.authService.LoginWithGoogleOauthToken(payload.Token)

	if err != nil {
		err = auth.ErrorNotAuthorized
		context.AbortWithStatusJSON(auth.GetHttpStatusCode(err), api.NewErrorResponse(err))
		return
	}

	auth.SetAuthUserInContext(authUser, context)
	context.SetCookie("authToken", token, 3600, "/", "", false, true)

	context.IndentedJSON(http.StatusOK, auth.LoginResponse{Token: token, User: user.ConvertModelToResponse(authUser)})
}

// @Summary Get current user
// @Description get current user
// @Tags auth
// @Accept  json
// @Produce  json
// @Success 200 {object} user.UserResponse
// @Failure 204 {object} api.ErrorResponse
// @Router /auth/me [get]
func (controller *AuthApiController) Me(context *gin.Context) {
	authUser, _ := auth.GetAuthUserFromContext(context)

	context.IndentedJSON(http.StatusOK, user.ConvertModelToResponse(authUser))
}

// @Summary Log out
// @Description log out
// @Tags auth
// @Accept  json
// @Produce  json
// @Success 200 {object} api.SuccessResponse
// @Router /auth/logout [post]
func (controller *AuthApiController) Logout(context *gin.Context) {
	context.SetCookie("authToken", "", -1, "/", "", false, true)
	context.IndentedJSON(http.StatusOK, api.NewSuccessResponse("Logout successful!"))
}
