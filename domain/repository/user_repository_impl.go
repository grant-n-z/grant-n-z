package repository

import (
	"github.com/tomoyane/grant-n-z/domain/entity"
	"github.com/tomoyane/grant-n-z/infra"
)

type UserRepositoryImpl struct {
}

// Find user by users.email
func (r UserRepositoryImpl) FindByEmail(email string) *entity.User {
	user := entity.User{}

	if err := infra.Db.Where("email = ?", email).First(&user).Error; err != nil {
		if err.Error() == "record not found" {
			return &entity.User{}
		}
		return nil
	}

	return &user
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
