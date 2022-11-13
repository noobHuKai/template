package utils

import (
	"github.com/noobHuKai/internal/g"
	request "github.com/noobHuKai/internal/models/request_model"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
)

func CreateToken(username string) string {
	claim := request.Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(3 * time.Hour * time.Duration(1))), // 过期时间3小时
			IssuedAt:  jwt.NewNumericDate(time.Now()),                                       // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                                       // 生效时间
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim) // 使用HS256算法
	tokenString, err := token.SignedString(g.JWTSecret)
	if err != nil {
		g.Logger.Error(err.Error())
	}
	return tokenString
}

func ParseToken(tokenss string) (*request.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenss, &request.Claims{}, jwtSecret())
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("that's not even a token")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("token is expired")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("token not active yet")
			} else {
				return nil, errors.New("couldn't handle this token")
			}
		}
	}
	if claims, ok := token.Claims.(*request.Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("couldn't handle this token")
}

func jwtSecret() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return g.JWTSecret, nil
	}
}
