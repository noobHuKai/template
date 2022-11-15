package utils

import (
	"github.com/noobHuKai/app/g"
	"github.com/noobHuKai/app/model/common/request"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
)

func CreateToken(uid uint, duration time.Duration) string {
	claim := request.Claims{
		UID: uid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)), // 过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),               // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),               // 生效时间
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim) // 使用HS256算法
	tokenString, err := token.SignedString(g.JWTSecret)
	if err != nil {
		g.Logger.Error(err.Error())
	}
	return tokenString
}

func ParseToken(tokens string) (*request.Claims, error) {
	token, err := jwt.ParseWithClaims(tokens, &request.Claims{}, jwtSecret())
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
