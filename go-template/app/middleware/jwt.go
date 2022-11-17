package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/noobHuKai/app/model/common/response"
	"github.com/noobHuKai/app/utils"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get token
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			response.FailUnauthorized(c, "Not Found Token")
			return
		}
		claims, err := utils.ParseToken(token)
		if err != nil {
			response.FailUnauthorized(c, err.Error())
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}
