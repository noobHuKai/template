package request_model

// AdminUserLoginReq 管理员用户登录请求结构体
type AdminUserLoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
