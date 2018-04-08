package entity

type UserEntity struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
