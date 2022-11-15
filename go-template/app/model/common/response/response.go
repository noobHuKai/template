package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	StatusOk      = 2000
	StatusFailure = 5000
)

func Result(c *gin.Context, code int, msg string, data interface{}) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		msg,
		data,
	})
	c.Abort()
}

func Ok(c *gin.Context) {
	Result(c, StatusOk, "", "")
}
func OkWithData(c *gin.Context, data interface{}) {
	Result(c, StatusOk, "", data)
}
func OkWithMessage(c *gin.Context, msg string) {
	Result(c, StatusOk, msg, "")
}
func OkWithMsgAndData(c *gin.Context, msg string, data interface{}) {
	Result(c, StatusOk, msg, data)
}

func Fail(c *gin.Context) {
	Result(c, StatusFailure, "", "")
}

func FailWithMessage(c *gin.Context, msg string) {
	Result(c, StatusFailure, msg, "")
}
func FailWithData(c *gin.Context, data interface{}) {
	Result(c, StatusFailure, "", data)
}
func FailWithMsgAndData(c *gin.Context, msg string, data interface{}) {
	Result(c, StatusFailure, msg, data)
}
