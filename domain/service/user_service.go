package service

import (
	"golang.org/x/crypto/bcrypt"
	"github.com/tomoyane/grant-n-z/domain/entity"
	"github.com/tomoyane/grant-n-z/domain/repository"
)

var userRepository = repository.UserRepositoryImpl{}.NewUserRepository()

func EncryptPw(password string) string {
	hash, _ := bcrypt.GenerateFromPassword(
		[] byte(password),
		bcrypt.DefaultCost,
	)
	return string(hash)
}

func GetUserByEmail(email string) *entity.User {
	return userRepository.FindByEmail(email)
}

func InsertUser(user entity.User) *entity.User {
	return userRepository.Save(user)
}
