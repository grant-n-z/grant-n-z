package service

import (
	"golang.org/x/crypto/bcrypt"
	"github.com/tomo0111/grant-n-z/app/domains/repository"
	"github.com/tomo0111/grant-n-z/app/domains/entity"
)

type UserService struct{}

var userRepository = repository.UserRepositoryImpl{}.NewUserRepository()

func (s UserService) EncryptPw(password string) string {
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