package domain

type LoginParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Role     int8   `json:"role_id"`
}