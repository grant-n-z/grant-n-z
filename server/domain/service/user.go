package service

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/satori/go.uuid"

	"github.com/tomoyane/grant-n-z/server/domain/entity"
	"github.com/tomoyane/grant-n-z/server/domain/repository"
	"github.com/tomoyane/grant-n-z/server/log"
)

type UserService struct {
	UserRepository repository.UserRepository
}

func NewUserService() UserService {
	return UserService{UserRepository: repository.UserRepositoryImpl{}}
}

func (us UserService) EncryptPw(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([] byte(password), bcrypt.DefaultCost)
	return string(hash)
}

func (us UserService) ComparePw(passwordHash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	if err != nil {
		log.Logger.Error("error compare password", err)
		return false
	}

	return true
}

func (us UserService) GetUserById(id int) (*entity.User, *entity.ErrorResponse) {
	return us.UserRepository.FindById(id)
}

func (us UserService) InsertUser(user *entity.User) (*entity.User, *entity.ErrorResponse) {
	user.Uuid, _ = uuid.NewV4()
	user.Password = us.EncryptPw(user.Password)
	return us.UserRepository.Save(*user)
}
