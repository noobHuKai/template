package response_model

type AdminUserLoginRes struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
