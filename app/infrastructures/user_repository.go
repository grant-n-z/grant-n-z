package infrastructures

import "grant-n-z/app/domains/entity"

type UserRepository interface {
	FindByEmail(email string) *entity.Users

	Save(users entity.Users) *entity.Users

	Update(users entity.Users) *entity.Users
}