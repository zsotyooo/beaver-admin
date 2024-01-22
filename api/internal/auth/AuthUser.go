package auth

import (
	"api/internal/user"
	"api/pkg/session"

	"github.com/gin-gonic/gin"
)

type AuthUser struct {
	User user.User
}

func NewAuthUser() *AuthUser {
	return &AuthUser{}
}

func (authUser *AuthUser) SetUser(u user.User) {
	authUser.User = u
}

func (authUser *AuthUser) Init(context *gin.Context) {
	authUserInContext, hasAuthUser := context.Get("authUser")
	if hasAuthUser {
		authUser.User = authUserInContext.(user.User)
		return
	}
	authUserInSession, hasAuthUser := session.Get[user.User](context, "authUser")
	if hasAuthUser {
		authUser.User = authUserInSession
		return
	}

	authUser.SetUser(user.User{})
}

func (authUser *AuthUser) Store(context *gin.Context) {
	if authUser.User.ID == 0 {
		session.Remove(context, "authUser")
		context.Set("authUser", nil)
		return
	}
	context.Set("authUser", authUser.User)
	session.Set[user.User](context, "authUser", authUser.User)
}

func (_ *AuthUser) Delete(context *gin.Context) {
	context.Set("authUser", nil)
}

func (authUser *AuthUser) IsAuthenticated() bool {
	return authUser.User.ID > 0
}

func (authUser *AuthUser) IsAuthUser(user user.User) bool {
	return authUser.IsAuthenticated() && authUser.User.ID == user.ID
}
