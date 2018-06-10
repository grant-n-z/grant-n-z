package service

import (
	"authentication-server/app/domains/repository"
	"authentication-server/app/domains/entity"
	"golang.org/x/crypto/bcrypt"
	"authentication-server/app/controllers/base"
)

type UserService struct {}

var userRepository repository.UserRepository

func (s UserService) BcryptPw(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([] byte(password), bcrypt.DefaultCost)
	return string(hash)
}

func (s UserService) GetUserByEmail(email string) base.BaseResponse {
	return userRepository.FindByEmail(email)
}

func (s UserService) InsertUser(users entity.Users) base.BaseResponse {
	return userRepository.Save(users)
}