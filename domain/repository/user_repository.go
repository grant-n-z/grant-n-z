package repository

import (
	"github.com/tomoyane/grant-n-z/domain/entity"
)

type UserRepository interface {
	FindByEmail(email string) *entity.User

	FindByUserNameAndUuid(username string, uuidStr string) *entity.User

	Save(user entity.User) *entity.User

	Update(user entity.User) *entity.User
}
