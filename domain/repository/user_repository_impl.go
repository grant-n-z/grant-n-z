package repository

import (
	"github.com/tomoyane/grant-n-z/domain/entity"
	"github.com/tomoyane/grant-n-z/infra"
)

type UserRepositoryImpl struct {
}

// Find user by users.email
func (u UserRepositoryImpl) FindByEmail(email string) *entity.User {
	user := entity.User{}

	if err := infra.Db.Where("email = ?", email).First(&user).Error; err != nil {
		if err.Error() == "record not found" {
			return &entity.User{}
		}
		return nil
	}

	return &user
}

// Find user by user.username and users.uuid
func (u UserRepositoryImpl) FindByUserNameAndUuid(username string, uuidStr string) *entity.User  {
	user := entity.User{}

	if err := infra.Db.Where("username = ? AND uuid = ?", username, uuidStr).First(&user).Error; err != nil {
		if err.Error() == "record not found" {
			return &entity.User{}
		}
		return nil
	}

	return &user
}

// Save to user
func (u UserRepositoryImpl) Save(user entity.User) *entity.User {
	if err := infra.Db.Create(&user).Error; err != nil {
		return nil
	}

	return &user
}

// Update to user
func (u UserRepositoryImpl) Update(user entity.User) *entity.User {
	if err := infra.Db.Save(&user).Error; err != nil {
		return nil
	}

	return &user
}

// Update to user column
func (u UserRepositoryImpl) UpdateUserColumn(user entity.User, column string) *entity.User {
	var data string

	switch column {
	case "username":
		data = user.Username
	case "email":
		data = user.Email
	case "password":
		data = user.Password
	}

	if err := infra.Db.Model(&entity.User{}).Where("email = ?", user.Email).Update(column, data).Error; err != nil {
		return nil
	}

	return &user
}
