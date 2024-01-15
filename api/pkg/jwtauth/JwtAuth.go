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

func New() *JwtAuth {
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
		err = ErrorCreatingToken
		signedToken = ""
	}
	return
}

func (jwtAuth *JwtAuth) ValidateToken(signedToken string) (claims *Claims, err error) {
	claims, err = jwtAuth.parseTokenClaims(signedToken)

	if claims.ExpiresAt.Time.Unix() < time.Now().Local().Unix() {
		err = ErrorTokenExpired
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
			err = ErrorInvalidTokenSignature
		}
		err = ErrorInvalidToken
		return
	}
	if !token.Valid {
		err = ErrorInvalidToken
		return
	}
	claims, ok := token.Claims.(*Claims)
	if !ok {
		err = ErrorInvalidClaims
	}
	return
}
