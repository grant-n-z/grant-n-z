package entity

type Users struct {
	Id        int       `json:"id"`
	Uuid      string    `json:"uuid"`
	Username  string    `validate:"required"json:"username"`
	Email     string    `validate:"required,email"json:"email"`
	Password  string    `validate:"required"json:"password"`
}

func (Users) TableName() string {
	return "users"
}
