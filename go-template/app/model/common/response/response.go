package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	CodeOk                  = 200 // 通用成功
	CodeBadRequest          = 400 // 通用失败
	CodeUnauthorized        = 401 // 未授权
	CodeUnProcessableEntity = 422 // 请求格式错误

)

type JsonResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResultJson(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(http.StatusOK, JsonResponse{
		code,
		msg,
		data,
	})
	c.Abort()
}

func Ok(c *gin.Context) {
	ResultJson(c, CodeOk, "", nil)
}
func OkWithMsg(c *gin.Context, msg string) {
	ResultJson(c, CodeOk, msg, nil)
}
func OkWithData(c *gin.Context, data interface{}) {
	ResultJson(c, CodeOk, "", data)
}

func Fail(c *gin.Context) {
	ResultJson(c, CodeBadRequest, "", nil)
}
func FailWithMsg(c *gin.Context, msg string) {
	ResultJson(c, CodeBadRequest, msg, nil)
}
func FailWithData(c *gin.Context, data interface{}) {
	ResultJson(c, CodeBadRequest, "", data)
}

func FailFormatError(c *gin.Context, msg string) {
	ResultJson(c, CodeUnProcessableEntity, msg, nil)
}

func FailUnauthorized(c *gin.Context, msg string) {
	ResultJson(c, CodeUnauthorized, msg, nil)
}
