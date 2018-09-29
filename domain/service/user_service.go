package service

import (
	"github.com/satori/go.uuid"
	"github.com/tomoyane/grant-n-z/domain/entity"
	"github.com/tomoyane/grant-n-z/domain/repository"
	"github.com/tomoyane/grant-n-z/handler"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepository repository.UserRepository
}

func (u UserService) EncryptPw(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([] byte(password), bcrypt.DefaultCost)
	return string(hash)
}

func (u UserService) GetUserByEmail(email string) *entity.User {
	return u.UserRepository.FindByEmail(email)
}

func (u UserService) GetUserByNameAndUuid(username string, uuid string) *entity.User {
	return u.UserRepository.FindByUserNameAndUuid(username, uuid)
}

func (u UserService) InsertUser(user entity.User) *entity.User {
	user.Uuid, _ = uuid.NewV4()
	user.Password = u.EncryptPw(user.Password)
	return u.UserRepository.Save(user)
}

func (u UserService) UpdateUser(user entity.User) *entity.User {
	user.Password = u.EncryptPw(user.Password)
	return u.UserRepository.Update(user)
}

func (u UserService) UpdateUserColumn(user entity.User, column string) *entity.User {
	user.Password = u.EncryptPw(user.Password)
	return u.UserRepository.UpdateUserColumn(user, column)
}

func (u UserService) PostUserData(user *entity.User) *handler.ErrorResponse {
	userData := u.GetUserByEmail(user.Email)
	if userData == nil {
		return handler.InternalServerError("")
	}

	if len(userData.Email) > 0 {
		return handler.Conflict("")
	}

	userData = u.InsertUser(*user)
	if userData == nil {
		return handler.InternalServerError("")
	}

	return nil
}

func (u UserService) PutUserColumnData(user *entity.User, column string) *handler.ErrorResponse {
	userData := u.GetUserByEmail(user.Email)
	if userData == nil {
		return handler.InternalServerError("")
	}

	if len(userData.Email) == 0 {
		return handler.NotFound("")
	}

	userData = u.UpdateUserColumn(*user, column)
	if userData == nil {
		return handler.InternalServerError("")
	}

	return nil
}