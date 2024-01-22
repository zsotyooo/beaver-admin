package middlewares

import (
	"api/internal/api"
	"api/internal/auth"
	"api/internal/user"
	"api/pkg/jwtauth"
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

// It inserts the authUser in the context and session if there is a valid token in the request (Auth header -> Cookie).
// It does not abort the request if there is no valid token.
func (middleware *AuthorizeMiddleware) Authorize() gin.HandlerFunc {
	return func(context *gin.Context) {
		logger.Debug(map[string]interface{}{}, "Authorize")
		token, err := jwtauth.ExtractAuthTokenFromContext(context)

		authUser := auth.NewAuthUser()

		if err != nil {
			authUser.Delete(context)
			context.Next()
			return
		}

		authUser.Init(context)

		if authUser.IsAuthenticated() {
			context.Next()
			return
		}
		identifiedUser, err := middleware.authService.Authorize(token)

		if err == nil {
			authUser.SetUser(identifiedUser)
			authUser.Store(context)
		} else {
			authUser.Delete(context)
		}

		context.Next()
	}
}

// It inserts the authUser in the context if there is a valid token in the request (Auth header -> Cookie).
// It aborts the request if the Authorization fails.
func (middleware *AuthorizeMiddleware) EnsureLoggedIn() gin.HandlerFunc {
	return func(context *gin.Context) {
		authUser := auth.NewAuthUser()
		authUser.Init(context)

		if !authUser.IsAuthenticated() {
			err := auth.ErrorNotAuthorized
			context.AbortWithStatusJSON(auth.GetHttpStatusCode(err), api.NewErrorResponse(err))
		}
	}
}

// It inserts the authUser in the context if there is a valid token in the request (Auth header -> Cookie).
// It aborts the request if the Authorization fails, or the user doesn't have any of the roles.
func (middleware *AuthorizeMiddleware) MustHaveRole(roles []user.UserRole) gin.HandlerFunc {
	return func(context *gin.Context) {
		authUser := auth.NewAuthUser()
		authUser.Init(context)

		if !authUser.IsAuthenticated() {
			err := auth.ErrorNotAuthorized
			context.AbortWithStatusJSON(auth.GetHttpStatusCode(err), api.NewErrorResponse(err))
			return
		}

		if authUser.User.HasRole(roles) {
			context.Next()
			return
		}

		err := auth.ErrorForbiddenForUser
		context.AbortWithStatusJSON(auth.GetHttpStatusCode(err), api.NewErrorResponse(err))
	}
}
