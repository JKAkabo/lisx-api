package models

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type AccessToken struct {
	Token string `json:"token"`
	ExpiresAt int `json:"expiresAt"`
	UserID int `json:"userID"`
}
