package middlewares

import (
	"api/internal/api"
	"api/internal/auth"
	"api/internal/user"
	"api/pkg/logger"

	"github.com/gin-gonic/gin"
)

type AuthorizeMiddleware struct {
	authService auth.AuthService
}

func NewAuthorizeMiddleware() *AuthorizeMiddleware {
	return &AuthorizeMiddleware{
		authService: *auth.NewAuthService(),
	}
}

// It inserts the authUser in the context if there is a valid token in the request (Auth header -> Cookie).
// It does not abort the request if there is no valid token.
func (middleware *AuthorizeMiddleware) Authorize() gin.HandlerFunc {
	return func(context *gin.Context) {
		logger.Debug(map[string]interface{}{}, "Authorize")
		token, err := auth.GetAuthTokenFromContext(context)

		if err != nil {
			auth.DeleteAuthUserFromContext(context)
			context.Next()
			return
		}

		_, hasAuthUser := auth.GetAuthUserFromContext(context)
		if hasAuthUser {
			context.Next()
			return
		}
		authUser, err := middleware.authService.Authorize(token)

		if err == nil {
			auth.SetAuthUserInContext(authUser, context)
		} else {
			auth.DeleteAuthUserFromContext(context)
		}

		context.Next()
	}
}

// It inserts the authUser in the context if there is a valid token in the request (Auth header -> Cookie).
// It aborts the request if the Authorization fails.
func (middleware *AuthorizeMiddleware) EnsureLoggedIn() gin.HandlerFunc {
	return func(context *gin.Context) {
		_, hasAuthUser := auth.GetAuthUserFromContext(context)

		if !hasAuthUser {
			err := auth.ErrorNotAuthorized
			context.AbortWithStatusJSON(auth.GetHttpStatusCode(err), api.NewErrorResponse(err))
		}
	}
}

// It inserts the authUser in the context if there is a valid token in the request (Auth header -> Cookie).
// It aborts the request if the Authorization fails, or the user doesn't have any of the roles.
func (middleware *AuthorizeMiddleware) MustHaveRole(roles []user.UserRole) gin.HandlerFunc {
	return func(context *gin.Context) {
		token, err := auth.GetAuthTokenFromContext(context)

		if err != nil {
			auth.DeleteAuthUserFromContext(context)
			context.AbortWithStatusJSON(auth.GetHttpStatusCode(err), api.NewErrorResponse(err))
			return
		}

		_, err = middleware.authService.Authorize(token)
		if err != nil {
			context.AbortWithStatusJSON(auth.GetHttpStatusCode(err), api.NewErrorResponse(err))
			return
		}
		authUser, hasAuthUser := auth.GetAuthUserFromContext(context)

		if !hasAuthUser {
			context.AbortWithStatusJSON(auth.GetHttpStatusCode(err), api.NewErrorResponse(err))
			return
		}

		if authUser.HasRole(roles) {
			context.Next()
			return
		}

		err = auth.ErrorForbiddenForUser
		context.AbortWithStatusJSON(auth.GetHttpStatusCode(err), api.NewErrorResponse(err))
	}
}
