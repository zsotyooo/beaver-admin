package jwtauth

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func extractAuthTokenFromHeader(context *gin.Context) (token string, err error) {
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

func extractAuthTokenFromCookie(context *gin.Context) (token string, err error) {
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

func ExtractAuthTokenFromContext(context *gin.Context) (token string, err error) {
	token, err = extractAuthTokenFromHeader(context)

	if err == nil {
		return
	}

	token, err = extractAuthTokenFromCookie(context)

	return
}
