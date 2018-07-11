package repository

import (
	"github.com/tomo0111/grant-n-z/app/infrastructures"
	"github.com/tomo0111/grant-n-z/app/domains/entity"
	"github.com/tomo0111/grant-n-z/app"
)

type UserRepositoryImpl struct{}

// Infrastructure UserRepository implementation
func (r UserRepositoryImpl) NewUserRepository() infrastructures.UserRepository {
	return &UserRepositoryImpl{}
}

// Find users by users.email
func (r UserRepositoryImpl) FindByEmail(email string) *entity.Users {
	users := entity.Users{}

	if err := app.Db.Where("email = ?", email).First(&users).Error; err != nil {
		if err.Error() == "record not found" {
			return &entity.Users{}
		}
		return nil
	}

	return &users
}

// Save to users
func (r UserRepositoryImpl) Save(users entity.Users) *entity.Users {
	if err := app.Db.Create(&users).Error; err != nil {
		return nil
	}

	return &users
}

// Update to users
func (r UserRepositoryImpl) Update(users entity.Users) *entity.Users {
	if err := app.Db.Update(&users).Error; err != nil {
		return nil
	}

	return &users
}
