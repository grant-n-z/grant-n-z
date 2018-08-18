package repository

import (
	"github.com/tomoyane/grant-n-z/domain/entity"
	"github.com/tomoyane/grant-n-z/infra"
)

type UserRepositoryImpl struct {
}

// Find users by users.email
func (r UserRepositoryImpl) FindByEmail(email string) *entity.User {
	users := entity.User{}

	if err := infra.Db.Where("email = ?", email).First(&users).Error; err != nil {
		if err.Error() == "record not found" {
			return &entity.User{}
		}
		return nil
	}

	return &users
}

// Save to user
func (r UserRepositoryImpl) Save(user entity.User) *entity.User {
	if err := infra.Db.Create(&user).Error; err != nil {
		return nil
	}

	return &user
}

// Update to user
func (r UserRepositoryImpl) Update(user entity.User) *entity.User {
	if err := infra.Db.Update(&user).Error; err != nil {
		return nil
	}

	return &user
}
