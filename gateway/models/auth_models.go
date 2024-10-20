package models

type RegisterRequest struct {
	Username string `	json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
