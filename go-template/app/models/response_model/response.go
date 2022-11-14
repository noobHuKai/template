package response_model

import (
	"github.com/gin-gonic/gin"
	"github.com/noobHuKai/app/constants"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Ok(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code:    constants.StatusOk,
		Message: "",
		Data:    nil,
	})
	c.Abort()
}
func OkWithData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    constants.StatusOk,
		Message: "",
		Data:    data,
	})
	c.Abort()
}

func Fail(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code:    constants.StatusFailure,
		Message: "",
		Data:    nil,
	})
	c.Abort()
}

func FailWithMsg(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, Response{
		Code:    constants.StatusFailure,
		Message: msg,
		Data:    nil,
	})
	c.Abort()

}

func FailWithMsgAndData(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    constants.StatusFailure,
		Message: msg,
		Data:    data,
	})
	c.Abort()
}

func JWTFail(c *gin.Context, msg string) {
	JWTFailWithMsg(c, msg)
}
func JWTFailWithMsg(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, Response{
		Code:    constants.StatusFailure,
		Message: msg,
		Data:    nil,
	})
	c.Abort()
}
