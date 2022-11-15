package request

import "github.com/golang-jwt/jwt/v4"

type Claims struct {
	UID uint
	jwt.RegisteredClaims
}
