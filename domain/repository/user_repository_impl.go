package repository

import (
	"github.com/tomoyane/grant-n-z/domain/entity"
	"github.com/tomoyane/grant-n-z/common"
)

type UserRepositoryImpl struct{
}

func (r UserRepositoryImpl) NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

// Find users by users.email
func (r UserRepositoryImpl) FindByEmail(email string) *entity.User {
	users := entity.User{}

	if err := common.Db.Where("email = ?", email).First(&users).Error; err != nil {
		if err.Error() == "record not found" {
			return &entity.User{}
		}
		return nil
	}

	return &users
}

// Save to user
func (r UserRepositoryImpl) Save(user entity.User) *entity.User {
	if err := common.Db.Create(&user).Error; err != nil {
		return nil
	}

	return &user
}

// Update to user
func (r UserRepositoryImpl) Update(user entity.User) *entity.User {
	if err := common.Db.Update(&user).Error; err != nil {
		return nil
	}

	return &user
}
