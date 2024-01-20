package jwtauth

import (
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

type JwtAuth struct {
	SecretKey         string
	Issuer            string
	ExpirationMinutes int64
}

type Claims struct {
	Email string
	jwt.RegisteredClaims
}

func NewJwtAuth() *JwtAuth {
	return &JwtAuth{
		SecretKey:         os.Getenv("JWT_SECRET"),
		Issuer:            "BeaverAdmin",
		ExpirationMinutes: 24 * 7 * 60,
	}
}

func (jwtAuth *JwtAuth) GenerateToken(email string) (signedToken string, err error) {
	claims := &Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Minute * time.Duration(jwtAuth.ExpirationMinutes))),
			Issuer:    jwtAuth.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err = token.SignedString([]byte(jwtAuth.SecretKey))
	if err != nil {
		err = Error(ErrorCreatingToken)
		signedToken = ""
	}
	return
}

func (jwtAuth *JwtAuth) ValidateToken(signedToken string) (claims *Claims, err error) {
	claims, err = jwtAuth.parseTokenClaims(signedToken)

	if claims.ExpiresAt.Time.Unix() < time.Now().Local().Unix() {
		err = ErrorTokenExpired
	}
	if err != nil {
		err = Error(err)
		claims = nil
		return
	}
	return
}

func (jwtAuth *JwtAuth) parseTokenClaims(signedToken string) (claims *Claims, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtAuth.SecretKey), nil
		},
	)
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			err = Error(ErrorInvalidTokenSignature)
		}
		err = Error(ErrorInvalidToken)
		return
	}
	if !token.Valid {
		err = Error(ErrorInvalidToken)
		return
	}
	claims, ok := token.Claims.(*Claims)
	if !ok {
		err = Error(ErrorInvalidClaims)
	}
	return
}
