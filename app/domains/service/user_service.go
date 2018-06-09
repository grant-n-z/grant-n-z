package service

import (
	"authentication-server/app/domains/repository"
	"authentication-server/app/controllers"
	"authentication-server/app/domains/entity"
)

type UserService struct {}

var userRepository repository.UserRepository

func (s UserService) GetUserById(id int) controllers.BaseResponse {
	return userRepository.FindById(id)
}

func (s UserService) InsertUser(users entity.Users) bool{
	return userRepository.Save(users)
}
