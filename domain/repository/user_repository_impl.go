package repository

import (
	"github.com/tomoyane/revel-performance/app"
	"github.com/tomoyane/grant-n-z/domain/entity"
)

type UserRepositoryImpl struct{
}

func (r UserRepositoryImpl) NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

// Find users by users.email
func (r UserRepositoryImpl) FindByEmail(email string) *entity.User {
	users := entity.User{}

	if err := app.Db.Where("email = ?", email).First(&users).Error; err != nil {
		if err.Error() == "record not found" {
			return &entity.User{}
		}
		return nil
	}

	return &users
}

// Save to user
func (r UserRepositoryImpl) Save(user entity.User) *entity.User {
	if err := app.Db.Create(&user).Error; err != nil {
		return nil
	}

	return &user
}

// Update to user
func (r UserRepositoryImpl) Update(user entity.User) *entity.User {
	if err := app.Db.Update(&user).Error; err != nil {
		return nil
	}

	return &user
}
