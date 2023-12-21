package models

type LoginRequest struct {
	Username string `json:"username" xml:"username" form:"username"`
	Password string `json:"password" xml:"password" form:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type User struct {
	ID             int
	Username       string
	Password       string
	FavoritePhrase string
}
