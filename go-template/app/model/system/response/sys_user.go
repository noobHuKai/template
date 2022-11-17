package response

type LoginResponse struct {
	Token     string `json:"token"`
	ExpiresAt string `json:"expires_at"`
}

type UserBasicInfoResponse struct {
	UserId   string `json:"userId"`
	UserName string `json:"userName"`
	UserRole string `json:"userRole"`
}
