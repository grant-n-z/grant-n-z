package service

import (
	"authentication-server/app/domains/repository"
	"authentication-server/app/domains/entity"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct{}

var userRepository repository.UserRepository

func (s UserService) BcryptPw(password string) string {
	hash, _ := bcrypt.GenerateFromPassword(
		[] byte(password),
		bcrypt.DefaultCost,
	)
	return string(hash)
}

func (s UserService) GetUserByEmail(email string) *entity.Users {
	return userRepository.FindByEmail(email)
}

func (s UserService) InsertUser(users entity.Users) *entity.Users {
	return userRepository.Save(users)
}