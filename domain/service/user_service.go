package service

import (
	"golang.org/x/crypto/bcrypt"
	"github.com/tomoyane/grant-n-z/domain/entity"
	"github.com/tomoyane/grant-n-z/domain/repository"
	"github.com/satori/go.uuid"
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

func (u UserService) GetUserByUuid(username string, uuid string) *entity.User {
	return u.UserRepository.FindByUserNameAndUuid(username, uuid)
}

func (u UserService) InsertUser(user entity.User) *entity.User {
	user.Uuid, _ = uuid.NewV4()
	user.Password = u.EncryptPw(user.Password)
	return u.UserRepository.Save(user)
}
