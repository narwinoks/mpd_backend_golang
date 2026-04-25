package user

type LoginRequest struct {
	Identity string `json:"identity" binding:"required" label:"Username or Email"`
	Password string `json:"password" binding:"required"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}
