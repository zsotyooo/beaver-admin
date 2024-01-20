package auth

import (
	"api/internal/user"
	"api/pkg/googleauth"

	"api/pkg/jwtauth"
)

type AuthService struct {
	userService *user.UserService
}

func NewAuthService() *AuthService {
	return &AuthService{
		userService: user.NewUserService(),
	}
}

func (service *AuthService) Authorize(token string) (authUser user.User, err error) {
	jwtAuth := jwtauth.NewJwtAuth()
	claims, validateErr := jwtAuth.ValidateToken(token)

	if validateErr != nil {
		return
	}

	authUser, err = service.userService.GetUserByEmail(claims.Email)

	if err != nil {
		err = Error(err)
	}

	return
}

func (service *AuthService) LoginWithGoogleOauthToken(googleToken string) (token string, authUser user.User, err error) {

	googleUser, err := googleauth.VerifyToken(googleToken)

	if err != nil {
		err = Error(err)
		return
	}

	userInDb, err := service.userService.MakeSureUserExists(googleUser.Email, googleUser.Name)

	if err != nil {
		err = Error(err)
		return
	}

	jwtAuth := jwtauth.NewJwtAuth()
	token, err = jwtAuth.GenerateToken(userInDb.Email)

	if err != nil {
		err = Error(err)
		return
	}

	authUser = userInDb
	return
}
