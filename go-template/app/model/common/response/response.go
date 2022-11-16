package response

import (
	"github.com/gin-gonic/gin"
)

const (
	StatusOK                  = 200 // 成功获取资源
	StatusCreated             = 201 // 成功创建或修改资源
	StatusNoContent           = 204 // 成功删除资源
	StatusBadRequest          = 400 // 请求错误，如GET参数有问题/PUT提交的数据错误等
	StatusUnauthorized        = 401 // 权限未通过认证
	StatusForbidden           = 403 // 有权限都禁止访问该资源
	StatusUnprocessableEntity = 422 // 请求格式错误
	StatusNotFound            = 404 // 请求的资源不存在
	StatusInternalServerError = 500 // 服务器端错误
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

func ResultJson(c *gin.Context, statusCode int, code int, msg string, data interface{}) {
	c.JSON(statusCode, JsonResponse{
		code,
		msg,
		data,
	})
	c.Abort()
}
func CreatedOk(c *gin.Context) {
	ResultJson(c, StatusOK, CodeOk, "", nil)
}
func CreatedOkWithData(c *gin.Context, data interface{}) {
	ResultJson(c, StatusOK, CodeOk, "", data)
}
func CreateFailUnprocessableEntityWithMsg(c *gin.Context, msg string) {
	ResultJson(c, StatusUnprocessableEntity, CodeFailure, msg, nil)
}
func CreateFailBadRequestWithMsg(c *gin.Context, msg string) {
	ResultJson(c, StatusBadRequest, CodeFailure, msg, nil)
}

func AuthorizedFail(c *gin.Context) {
	ResultJson(c, StatusUnauthorized, CodeFailure, "", nil)
}
func AuthorizedFailWithMsg(c *gin.Context, msg string) {
	ResultJson(c, StatusUnauthorized, CodeFailure, msg, nil)
}
