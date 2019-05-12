package service

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/satori/go.uuid"

	"github.com/tomoyane/grant-n-z/server/config"
	"github.com/tomoyane/grant-n-z/server/entity"
	"github.com/tomoyane/grant-n-z/server/log"
	"github.com/tomoyane/grant-n-z/server/usecase/repository"
)

type userServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUserService() UserService {
	log.Logger.Info("Inject `userRepository` to `UserService`")
	return userServiceImpl{userRepository: repository.UserRepositoryImpl{Db: config.Db}}
}

func (us userServiceImpl) EncryptPw(password string) string {
	hash, err := bcrypt.GenerateFromPassword([] byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Logger.Error("Error password hash", err)
		panic(err)
	}

	return string(hash)
}

func (us userServiceImpl) ComparePw(passwordHash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	if err != nil {
		log.Logger.Error("Error compare password", err)
		return false
	}

	return true
}

func (us userServiceImpl) GetUserById(id int) (*entity.User, *entity.ErrorResponse) {
	return us.userRepository.FindById(id)
}

func (us userServiceImpl) InsertUser(user *entity.User) (*entity.User, *entity.ErrorResponse) {
	user.Uuid, _ = uuid.NewV4()
	user.Password = us.EncryptPw(user.Password)
	return us.userRepository.Save(*user)
}
