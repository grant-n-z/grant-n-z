package service

import (
	"authentication-server/app/domains/repository"
	"authentication-server/app/controllers"
	"authentication-server/app/domains/entity"
)

type UserService struct {}

var userRepository repository.UserRepository

func (s UserService) GetUserByEmail(email string) controllers.BaseResponse {
	return userRepository.FindByEmail(email)
}

func (s UserService) InsertUser(users entity.Users) bool {
	return userRepository.Save(users)
}
