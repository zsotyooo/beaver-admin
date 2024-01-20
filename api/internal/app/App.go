package app

import (
	"api/internal/auth"
	"api/internal/user"

	"github.com/gin-gonic/gin"
)

type App struct{}

func NewApp() *App {
	return &App{}
}

func (app *App) AuthUser(context *gin.Context) (authUser user.User, hasAuthUser bool) {
	return auth.GetAuthUserFromContext(context)
}
