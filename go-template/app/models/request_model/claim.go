package request_model

import "github.com/golang-jwt/jwt/v4"

type Claims struct {
	UID uint `json:"uid"`
	jwt.RegisteredClaims
}
