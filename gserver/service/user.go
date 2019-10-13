package service

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/satori/go.uuid"

	"github.com/tomoyane/grant-n-z/gserver/cache"
	"github.com/tomoyane/grant-n-z/gserver/common/config"
	"github.com/tomoyane/grant-n-z/gserver/common/driver"
	"github.com/tomoyane/grant-n-z/gserver/data"
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

var usInstance UserService

type UserService interface {
	EncryptPw(password string) string

	ComparePw(passwordHash string, password string) bool

	GetUserById(id int) (*entity.User, *model.ErrorResBody)

	GetUserByEmail(email string) (*entity.User, *model.ErrorResBody)

	GetUserWithOperatorPolicyByEmail(email string) (*entity.User, *model.ErrorResBody)

	InsertUser(user *entity.User) (*entity.User, *model.ErrorResBody)

	InsertUserWithUserService(user *entity.User, userService *entity.UserService) (*entity.User, *model.ErrorResBody)

	UpdateUser(user *entity.User) (*entity.User, *model.ErrorResBody)
}

type userServiceImpl struct {
	userRepository data.UserRepository
	appConfig      config.AppConfig
	redisClient    cache.RedisClient
}

func GetUserServiceInstance() UserService {
	if usInstance == nil {
		usInstance = NewUserService()
	}
	return usInstance
}

func NewUserService() UserService {
	log.Logger.Info("New `UserService` instance")
	log.Logger.Info("Inject `UserRepository`, `AppConfig`, `RedisClient` to `UserService`")
	return userServiceImpl{
		userRepository: data.GetUserRepositoryInstance(driver.Db),
		appConfig:      config.App,
		redisClient:    cache.GetRedisClientInstance(),
	}
}

func (us userServiceImpl) EncryptPw(password string) string {
	hash, err := bcrypt.GenerateFromPassword([] byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Logger.Info("Error password hash", err.Error())
		return ""
	}

	return string(hash)
}

func (us userServiceImpl) ComparePw(passwordHash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	if err != nil {
		log.Logger.Info("Error compare password", err.Error())
		return false
	}

	return true
}

func (us userServiceImpl) GetUserById(id int) (*entity.User, *model.ErrorResBody) {
	return us.userRepository.FindById(id)
}

func (us userServiceImpl) GetUserByEmail(email string) (*entity.User, *model.ErrorResBody) {
	return us.userRepository.FindByEmail(email)
}

func (us userServiceImpl) GetUserWithOperatorPolicyByEmail(email string) (*entity.User, *model.ErrorResBody) {
	return us.userRepository.FindWithOperatorPolicyByEmail(email)
}

func (us userServiceImpl) InsertUser(user *entity.User) (*entity.User, *model.ErrorResBody) {
	user.Uuid, _ = uuid.NewV4()
	user.Password = us.EncryptPw(user.Password)
	return us.userRepository.Save(*user)
}

func (us userServiceImpl) InsertUserWithUserService(user *entity.User, userService *entity.UserService) (*entity.User, *model.ErrorResBody) {
	user.Uuid, _ = uuid.NewV4()
	user.Password = us.EncryptPw(user.Password)
	return us.userRepository.SaveWithUserService(*user, userService)
}

func (us userServiceImpl) UpdateUser(user *entity.User) (*entity.User, *model.ErrorResBody) {
	user.Password = us.EncryptPw(user.Password)
	return us.userRepository.Update(*user)
}
