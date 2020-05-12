package service

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/tomoyane/grant-n-z/gnz/cache"
	"github.com/tomoyane/grant-n-z/gnz/ctx"
	"github.com/tomoyane/grant-n-z/gnz/driver"
	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
)

var usInstance UserService

type UserService interface {
	// Generate initial user name
	GenInitialName() string

	// Encrypt password
	EncryptPw(password string) string

	// Compare encrypt password and decrypt password
	ComparePw(passwordHash string, password string) bool

	// Get User by user id
	GetUserById(id int) (*entity.User, *model.ErrorResBody)

	// Get User by user email
	GetUserByEmail(email string) (*entity.User, *model.ErrorResBody)

	// Get User and operator policy by user email
	GetUserWithOperatorPolicyByEmail(email string) (*model.UserWithOperatorPolicy, *model.ErrorResBody)

	// Get User and user service and service by user email
	GetUserWithUserServiceWithServiceByEmail(email string) (*model.UserWithUserServiceWithService, *model.ErrorResBody)

	// Get UserGroup by user_id and group_id
	GetUserGroupByUserIdAndGroupId(userId int, groupId int) (*entity.UserGroup, *model.ErrorResBody)

	// Get Users by group_id
	GetUserByGroupId(groupId int) ([]*model.UserResponse, *model.ErrorResBody)

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
type UserServiceImpl struct {
	UserRepository driver.UserRepository
	EtcdClient     cache.EtcdClient
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
	return UserServiceImpl{
		UserRepository: driver.GetUserRepositoryInstance(),
		EtcdClient:     cache.GetEtcdClientInstance(),
	}
}

func (us UserServiceImpl) GenInitialName() string {
	uid := uuid.New().String()
	return string([]rune(uid)[:6])
}

func (us UserServiceImpl) EncryptPw(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash)
}

func (us UserServiceImpl) ComparePw(passwordHash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	if err != nil {
		log.Logger.Info("Failed to compare password", err.Error())
		return false
	}

	return true
}

func (us UserServiceImpl) GetUserById(id int) (*entity.User, *model.ErrorResBody) {
	return us.UserRepository.FindById(id)
}

func (us UserServiceImpl) GetUserByEmail(email string) (*entity.User, *model.ErrorResBody) {
	return us.UserRepository.FindByEmail(email)
}

func (us UserServiceImpl) GetUserWithOperatorPolicyByEmail(email string) (*model.UserWithOperatorPolicy, *model.ErrorResBody) {
	return us.UserRepository.FindWithOperatorPolicyByEmail(email)
}

func (us UserServiceImpl) GetUserWithUserServiceWithServiceByEmail(email string) (*model.UserWithUserServiceWithService, *model.ErrorResBody) {
	return us.UserRepository.FindWithUserServiceWithServiceByEmail(email)
}

func (us UserServiceImpl) GetUserGroupByUserIdAndGroupId(userId int, groupId int) (*entity.UserGroup, *model.ErrorResBody) {
	return us.UserRepository.FindUserGroupByUserIdAndGroupId(userId, groupId)
}

func (us UserServiceImpl) GetUserServices() ([]*entity.UserService, *model.ErrorResBody) {
	return us.UserRepository.FindUserServices()
}

func (us UserServiceImpl) GetUserServiceByUserIdAndServiceId(userId int, serviceId int) (*entity.UserService, *model.ErrorResBody) {
	userService := us.EtcdClient.GetUserService(userId, serviceId)
	if userService != nil {
		return userService, nil
	}
	return us.UserRepository.FindUserServiceByUserIdAndServiceId(userId, serviceId)
}

func (us UserServiceImpl) GetUserByGroupId(groupId int) ([]*model.UserResponse, *model.ErrorResBody) {
	users, err := us.UserRepository.FindByGroupId(groupId)
	if err != nil{
		return nil, err
	}

	var userResponse []*model.UserResponse
	for _, user := range users {
		userResponse = append(userResponse, &model.UserResponse{Uuid: user.Uuid.String(), Username: user.Username, Email: user.Email})
	}
	return userResponse, nil
}

func (us UserServiceImpl) InsertUserGroup(userGroupEntity entity.UserGroup) (*entity.UserGroup, *model.ErrorResBody) {
	userGroup, err := us.GetUserGroupByUserIdAndGroupId(userGroupEntity.UserId, userGroupEntity.GroupId)
	if err != nil || userGroup != nil {
		conflictErr := model.Conflict("This user already joins group")
		return nil, conflictErr
	}
	return us.UserRepository.SaveUserGroup(userGroupEntity)
}

func (us UserServiceImpl) InsertUser(user entity.User) (*entity.User, *model.ErrorResBody) {
	user.Uuid = uuid.New()
	user.Password = us.EncryptPw(user.Password)
	return us.UserRepository.SaveUser(user)
}

func (us UserServiceImpl) InsertUserWithUserService(user entity.User, userService entity.UserService) (*entity.User, *model.ErrorResBody) {
	user.Uuid = uuid.New()
	user.Password = us.EncryptPw(user.Password)
	return us.UserRepository.SaveWithUserService(user, userService)
}

func (us UserServiceImpl) InsertUserService(userServiceEntity entity.UserService) (*entity.UserService, *model.ErrorResBody) {
	userService, err := us.UserRepository.FindUserServiceByUserIdAndServiceId(userServiceEntity.UserId, userServiceEntity.ServiceId)
	if err != nil || userService != nil {
		return nil, model.Conflict("Already the user has this service account")
	}
	return us.UserRepository.SaveUserService(userServiceEntity)
}

func (us UserServiceImpl) UpdateUser(user entity.User) (*entity.User, *model.ErrorResBody) {
	user.Id = ctx.GetUserId().(int)
	user.Uuid = ctx.GetUserUuid().(uuid.UUID)
	user.Password = us.EncryptPw(user.Password)
	return us.UserRepository.UpdateUser(user)
}
