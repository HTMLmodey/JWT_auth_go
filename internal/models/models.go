package models

type TokenRequest struct {
	UserID string `json:"user_id"`
	IP     string `json:"ip"`
}

type RefreshRequest struct {
	UserID       string `json:"user_id"`
	RefreshToken string `json:"refresh_token"`
	IP           string `json:"ip"`
	CurrentIP    string `json:"current_ip"`
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
