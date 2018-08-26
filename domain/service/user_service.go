package service

import (
	"golang.org/x/crypto/bcrypt"
	"github.com/tomoyane/grant-n-z/domain/entity"
	"github.com/tomoyane/grant-n-z/domain/repository"
)

type UserService struct {
	UserRepository repository.UserRepository
}

func (u UserService) EncryptPw(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([] byte(password), bcrypt.DefaultCost)
	return string(hash)
}

func (u UserService) ComparePw(passwordHash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	if err != nil {
		return false
	}

	return true
}

func (u UserService) GetUserByEmail(email string) *entity.User {
	return u.UserRepository.FindByEmail(email)
}

func (u UserService) InsertUser(user entity.User) *entity.User {
	return u.UserRepository.Save(user)
}
