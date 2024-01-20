package auth

import (
	"api/internal/api"
	"api/internal/user"
	"strings"

	"github.com/gin-gonic/gin"
)

func SetAuthUserInContext(user user.User, context *gin.Context) {
	context.Set("authUser", user)
}

func DeleteAuthUserFromContext(context *gin.Context) {
	context.Set("authUser", nil)
}

func GetAuthUserFromContext(context *gin.Context) (authUser user.User, hasAuthUser bool) {
	authUserInContext, hasAuthUser := context.Get("authUser")
	if !hasAuthUser || authUserInContext == nil {
		hasAuthUser = false
		return
	}
	authUser = authUserInContext.(user.User)
	return
}

func getAuthTokenFromHeader(context *gin.Context) (token string, err error) {
	header := context.Request.Header.Get("Authorization")

	if header == "" {
		err = Error(ErrorAuthTokenNotFoundInRequest)
		return
	}

	splitHeader := strings.Split(header, "Bearer ")
	if len(splitHeader) == 2 {
		token = strings.TrimSpace(splitHeader[1])
		return
	}

	err = Error(ErrorAuthTokenNotFoundInRequest)
	return
}

func getAuthTokenFromCookie(context *gin.Context) (token string, err error) {
	token, err = context.Cookie("authToken")

	if err != nil {
		err = Error(ErrorAuthTokenNotFoundInRequest)
		return
	}

	if token == "" {
		err = Error(ErrorAuthTokenNotFoundInRequest)
		return
	}

	return
}

func GetAuthTokenFromContext(context *gin.Context) (token string, err error) {
	token, err = getAuthTokenFromHeader(context)

	if err == nil {
		return
	}

	token, err = getAuthTokenFromCookie(context)

	return
}

func IsAuthUser(u user.User, context *gin.Context) bool {
	authUser, hasAuthUser := GetAuthUserFromContext(context)
	if !hasAuthUser {
		return false
	}
	return u.ID == authUser.ID
}

func abortWithError(context *gin.Context, err error) {
	context.AbortWithStatusJSON(GetHttpStatusCode(err), api.NewErrorResponse(err))
}
