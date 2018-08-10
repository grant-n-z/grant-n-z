package repository

import "github.com/tomoyane/grant-n-z/domain/entity"

type UserRepository interface {
	FindByEmail(email string) *entity.User

	Save(users entity.User) *entity.User

	Update(users entity.User) *entity.User
}
