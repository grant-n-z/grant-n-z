package repository

import (
	"authentication-server/app/domains/entity"
	"authentication-server/app"
	"authentication-server/app/controllers"
)

type UserRepository struct {}

// Find users by users.id
func (r UserRepository) FindById(id int) controllers.BaseResponse {
	var users entity.Users
	if err := app.Db.Where("id = ?", id).First(&users).Error; err != nil {
		return controllers.BaseResponse{}
	}

	response := controllers.BaseResponse {}
	response.Response = users

	return response
}

// Save to users
func (r UserRepository) Save(users entity.Users) bool {
	if err := app.Db.Create(users).Error; err != nil {
		return false
	}

	return true
}

// Update to users
func (r UserRepository) Update(users entity.Users) bool {
	if err := app.Db.Update(users).Error; err != nil {
		return false
	}

	return true
}
