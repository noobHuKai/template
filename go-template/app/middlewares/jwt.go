package middlewares

import (
	"github.com/gin-gonic/gin"
	response "github.com/noobHuKai/app/models/response_model"
	"github.com/noobHuKai/app/utils"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get token
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			response.JWTFail(c, "Not Found Token")
			return
		}
		claims, err := utils.ParseToken(token)
		if err != nil {
			response.JWTFail(c, err.Error())
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}
