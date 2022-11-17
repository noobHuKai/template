package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/noobHuKai/app/g"
	"github.com/noobHuKai/app/model/common/response"
)

var ctx = context.Background()

func TokenAuthorizeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get token
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			response.FailUnauthorized(c, "Not Found Token")
			return
		}
		uid, err := g.RDB.Get(ctx, token).Result()
		if err != nil {
			g.Logger.Error(err.Error())
			response.FailUnauthorized(c, "token is expired")
			return
		}

		RefreshToken(token)

		c.Set("uid", uid)
		c.Next()
	}
}

func RefreshToken(token string) {
	duration := g.RDB.TTL(ctx, token).Val()
	if duration.Minutes() < g.TimeRefreshToken {
		g.RDB.Expire(ctx, token, g.TimeExpireToken)
	}
}
