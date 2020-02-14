package service

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/google/uuid"

	"github.com/tomoyane/grant-n-z/gnz/cache"
	"github.com/tomoyane/grant-n-z/gnz/config"
	"github.com/tomoyane/grant-n-z/gnz/driver"
	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnzserver/ctx"
	"github.com/tomoyane/grant-n-z/gnz/data"
	"github.com/tomoyane/grant-n-z/gnzserver/entity"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
)

var usInstance UserService

type UserService interface {
	// Encrypt password
	EncryptPw(password string) string

	// Compare encrypt password and decrypt password
	ComparePw(passwordHash string, password string) bool

	// Get User by user id
	GetUserById(id int) (*entity.User, *model.ErrorResBody)

	// Get User by user email
	GetUserByEmail(email string) (*entity.User, *model.ErrorResBody)

	// Get User and operator policy by user email
	GetUserWithOperatorPolicyByEmail(email string) (*entity.UserWithOperatorPolicy, *model.ErrorResBody)

	// Get User and user service and service by user email
	GetUserWithUserServiceWithServiceByEmail(email string) (*entity.UserWithUserServiceWithService, *model.ErrorResBody)

	// Get UserGroup by user_id and group_id
	GetUserGroupByUserIdAndGroupId(userId int, groupId int) (*entity.UserGroup, *model.ErrorResBody)

	// Get all UserService
	GetUserServices() ([]*entity.UserService, *model.ErrorResBody)

	// Get UserService by user_id and service_id
	GetUserServiceByUserIdAndServiceId(userId int, serviceId int) (*entity.UserService, *model.ErrorResBody)

	// Insert UserGroup
	InsertUserGroup(userGroup entity.UserGroup) (*entity.UserGroup, *model.ErrorResBody)

	// Insert User
	InsertUser(user entity.User) (*entity.User, *model.ErrorResBody)

	// Insert User and UserService
	InsertUserWithUserService(user entity.User, userService entity.UserService) (*entity.User, *model.ErrorResBody)

	// Insert UserService
	InsertUserService(userServiceEntity entity.UserService) (*entity.UserService, *model.ErrorResBody)

	// Update User
	UpdateUser(user entity.User) (*entity.User, *model.ErrorResBody)
}

// UserService struct
type userServiceImpl struct {
	userRepository data.UserRepository
	appConfig      config.AppConfig
	redisClient    cache.RedisClient
}

// Get Policy instance.
// If use singleton pattern, call this instance method
func GetUserServiceInstance() UserService {
	if usInstance == nil {
		usInstance = NewUserService()
	}
	return usInstance
}

// Constructor
func NewUserService() UserService {
	log.Logger.Info("New `UserService` instance")
	return userServiceImpl{
		userRepository: data.GetUserRepositoryInstance(driver.Rdbms),
		appConfig:      config.App,
		redisClient:    cache.GetRedisClientInstance(),
	}
}

func (us userServiceImpl) EncryptPw(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Logger.Info("Failed to password hash", err.Error())
		return ""
	}

	return string(hash)
}

func (us userServiceImpl) ComparePw(passwordHash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	if err != nil {
		log.Logger.Info("Failed to compare password", err.Error())
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

func (us userServiceImpl) GetUserWithOperatorPolicyByEmail(email string) (*entity.UserWithOperatorPolicy, *model.ErrorResBody) {
	return us.userRepository.FindWithOperatorPolicyByEmail(email)
}

func (us userServiceImpl) GetUserWithUserServiceWithServiceByEmail(email string) (*entity.UserWithUserServiceWithService, *model.ErrorResBody) {
	return us.userRepository.FindWithUserServiceWithServiceByEmail(email)
}

func (us userServiceImpl) GetUserGroupByUserIdAndGroupId(userId int, groupId int) (*entity.UserGroup, *model.ErrorResBody) {
	return us.userRepository.FindUserGroupByUserIdAndGroupId(userId, groupId)
}

func (us userServiceImpl) GetUserServices() ([]*entity.UserService, *model.ErrorResBody) {
	return us.userRepository.FindUserServices()
}

func (us userServiceImpl) GetUserServiceByUserIdAndServiceId(userId int, serviceId int) (*entity.UserService, *model.ErrorResBody) {
	return us.userRepository.FindUserServiceByUserIdAndServiceId(userId, serviceId)
}

func (us userServiceImpl) InsertUserGroup(userGroupEntity entity.UserGroup) (*entity.UserGroup, *model.ErrorResBody) {
	userGroup, err := us.GetUserGroupByUserIdAndGroupId(userGroupEntity.UserId, userGroupEntity.GroupId)
	if err != nil || userGroup != nil {
		conflictErr := model.Conflict("This user already joins group")
		return nil, conflictErr
	}
	return us.userRepository.SaveUserGroup(userGroupEntity)
}

func (us userServiceImpl) InsertUser(user entity.User) (*entity.User, *model.ErrorResBody) {
	user.Uuid = uuid.New()
	user.Password = us.EncryptPw(user.Password)
	return us.userRepository.SaveUser(user)
}

func (us userServiceImpl) InsertUserWithUserService(user entity.User, userService entity.UserService) (*entity.User, *model.ErrorResBody) {
	user.Uuid = uuid.New()
	user.Password = us.EncryptPw(user.Password)
	return us.userRepository.SaveWithUserService(user, userService)
}

func (us userServiceImpl) InsertUserService(userServiceEntity entity.UserService) (*entity.UserService, *model.ErrorResBody) {
	userService, err := us.userRepository.FindUserServiceByUserIdAndServiceId(userServiceEntity.UserId, userServiceEntity.ServiceId)
	if err != nil || userService != nil {
		return nil, model.Conflict("Already the user has this service account")
	}
	return us.userRepository.SaveUserService(userServiceEntity)
}

func (us userServiceImpl) UpdateUser(user entity.User) (*entity.User, *model.ErrorResBody) {
	user.Id = ctx.GetUserId().(int)
	user.Uuid = ctx.GetUserUuid().(uuid.UUID)
	user.Password = us.EncryptPw(user.Password)
	return us.userRepository.UpdateUser(user)
}
