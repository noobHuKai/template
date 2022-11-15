package middleware

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/noobHuKai/app/g"
	"github.com/noobHuKai/app/model/common/response"
)

func TokenAuthorizeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get token
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			response.FailWithMessage(c, "Not Found Token")
			return
		}
		result, err := g.RDB.Get(context.Background(), token).Result()
		if err != nil {
			g.Logger.Error(err.Error())
			response.FailWithMessage(c, "Error Token")
			return
		}
		fmt.Println(result)
		//c.Set("claims", claims)
		c.Next()
	}
}
