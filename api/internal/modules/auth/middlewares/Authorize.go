package middlewares

import (
	AuthToken "api/packages/auth-token"
	"api/packages/logger"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func getTokenFromHeader(context *gin.Context) (token string, err error) {
	header := context.Request.Header.Get("Authorization")

	if header == "" {
		err = errors.New("Token not found!")
		return
	}

	splitHeader := strings.Split(header, "Bearer ")
	if len(splitHeader) == 2 {
		token = strings.TrimSpace(splitHeader[1])
		return
	}

	err = errors.New("Token not found!")
	return
}

func getTokenFromCookie(context *gin.Context) (token string, err error) {
	token, err = context.Cookie("authToken")

	if err != nil {
		return
	}

	if token == "" {
		err = errors.New("Token not found!")
		return
	}

	return
}

func Authorize() gin.HandlerFunc {
	return func(context *gin.Context) {
		logger.Debug(map[string]interface{}{"message": "Authorizing user"}, "Authorizing user")
		token, err := getTokenFromHeader(context)

		if err != nil {
			token, err = getTokenFromCookie(context)
		}

		if err != nil {
			context.IndentedJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			context.Abort()
			return
		}

		authToken := AuthToken.New()
		claims, err := authToken.ValidateToken(token)

		if err != nil {
			context.IndentedJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			context.Abort()
			return
		}

		logger.Debug(map[string]interface{}{
			"email": claims.Email,
		}, "Authorized user")

		context.Set("email", claims.Email)
		context.Next()
	}
}
