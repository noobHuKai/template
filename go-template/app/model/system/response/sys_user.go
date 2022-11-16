package response

type LoginResponse struct {
	Token     string `json:"token"`
	Uid       string `json:"uid"`
	ExpiresAt string `json:"expires_at"`
}
