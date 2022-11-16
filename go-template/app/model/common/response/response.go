package response

import (
	"github.com/gin-gonic/gin"
)

const (
	StatusCodeOk                  = 200 // 成功获取资源
	StatusCodeCreated             = 201 // 成功创建或修改资源
	StatusCodeNoContent           = 204 // 成功删除资源
	StatusCodeBadRequest          = 400 // 请求错误，如GET参数有问题/PUT提交的数据错误等
	StatusCodeUnauthorized        = 401 // 权限未通过认证
	StatusCodeForbidden           = 403 // 有权限都禁止访问该资源
	StatusCodeNotFound            = 404 // 请求的资源不存在
	StatusCodeInternalServerError = 500 // 服务器端错误
)
const (
	CodeOk      = 2000
	CodeFailure = 5000
)

type JsonResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResultJson(c *gin.Context,statusCode int, code int, msg string, data interface{}) {
	// 开始时间
	c.JSON(statusCode, JsonResponse{
		code,
		msg,
		data,
	})
	c.Abort()
}
func OkCreated(c *gin.Context){
	ResultJson(c,StatusCodeCreated,CodeOk,"","")
}

//	func Ok(c *gin.Context) {
//		Result(c, StatusOk, "", "")
//	}
//func OkWithData(c *gin.Context, data interface{}) {
//	Result(c, StatusOk, "", data)
//}
//func OkWithMessage(c *gin.Context, msg string) {
//	Result(c, StatusOk, msg, "")
//}
//func OkWithMsgAndData(c *gin.Context, msg string, data interface{}) {
//	Result(c, StatusOk, msg, data)
//}
//
//func Fail(c *gin.Context) {
//	Result(c, StatusFailure, "", "")
//}
//
//func FailWithMessage(c *gin.Context, msg string) {
//	Result(c, StatusFailure, msg, "")
//}
//func FailWithData(c *gin.Context, data interface{}) {
//	Result(c, StatusFailure, "", data)
//}
//func FailWithMsgAndData(c *gin.Context, msg string, data interface{}) {
//	Result(c, StatusFailure, msg, data)
//}
