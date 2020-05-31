package service

import (
	"strings"

	"crypto/md5"
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"

	"github.com/google/uuid"
	"github.com/tomoyane/grant-n-z/gnz/cache"
	"github.com/tomoyane/grant-n-z/gnz/cache/structure"
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

	// Get User by user uuid
	GetUserByUuid(uuid string) (*entity.User, *model.ErrorResBody)

	// Get User by user email
	GetUserByEmail(email string) (*entity.User, *model.ErrorResBody)

	// Get User and operator policy by user email
	GetUserWithOperatorPolicyByEmail(email string) (*model.UserWithOperatorPolicy, *model.ErrorResBody)

	// Get User and user service and service by user email
	GetUserWithUserServiceWithServiceByEmail(email string) (*model.UserWithUserServiceWithService, *model.ErrorResBody)

	// Get UserGroup by user uuid and group uuid
	GetUserGroupByUserUuidAndGroupUuid(userUuid string, groupUuid string) (*entity.UserGroup, *model.ErrorResBody)

	// Get Users by group uuid
	GetUserByGroupUuid(groupUuid string) ([]*model.UserResponse, *model.ErrorResBody)

	// Get all UserService
	GetUserServices() ([]*entity.UserService, *model.ErrorResBody)

	// Get UserService by user uuid and service uuid
	GetUserServiceByUserUuidAndServiceUuid(userUuid string, serviceUuid string) (*entity.UserService, *model.ErrorResBody)

	// Get UserPolicy in etcd by user uuid
	GetUserPoliciesByUserUuid(userUuid string) []structure.UserPolicy

	// Get UserPolicy in etcd by user uuid
	GetUserGroupsByUserUuid(userUuid string) []structure.UserGroup

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
		return false
	}

	return true
}

func (us UserServiceImpl) GetUserByUuid(uuid string) (*entity.User, *model.ErrorResBody) {
	user, err := us.UserRepository.FindByUuid(uuid)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, model.NotFound("Not found user")
		}
		return nil, model.InternalServerError(err.Error())
	}

	return user, nil
}

func (us UserServiceImpl) GetUserByEmail(email string) (*entity.User, *model.ErrorResBody) {
	user, err := us.UserRepository.FindByEmail(email)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, model.NotFound("Not found user")
		}
		return nil, model.InternalServerError()
	}

	return user, nil
}

func (us UserServiceImpl) GetUserWithOperatorPolicyByEmail(email string) (*model.UserWithOperatorPolicy, *model.ErrorResBody) {
	userWithOperatorPolicy, err := us.UserRepository.FindWithOperatorPolicyByEmail(email)
	if err != nil {
		return nil, model.InternalServerError(err.Error())
	}

	return userWithOperatorPolicy, nil
}

func (us UserServiceImpl) GetUserWithUserServiceWithServiceByEmail(email string) (*model.UserWithUserServiceWithService, *model.ErrorResBody) {
	userWithService, err := us.UserRepository.FindWithUserServiceWithServiceByEmail(email)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, model.NotFound("Not found service of user")
		}
		return nil, model.InternalServerError()
	}

	return userWithService, nil
}

func (us UserServiceImpl) GetUserGroupByUserUuidAndGroupUuid(userUuid string, groupUuid string) (*entity.UserGroup, *model.ErrorResBody) {
	userGroup, err := us.UserRepository.FindUserGroupByUserUuidAndGroupUuid(userUuid, groupUuid)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, model.NotFound("Not found group of user")
		}
		return nil, model.InternalServerError(err.Error())
	}

	return userGroup, nil
}

func (us UserServiceImpl) GetUserServices() ([]*entity.UserService, *model.ErrorResBody) {
	userServices, err := us.UserRepository.FindUserServices()
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, model.NotFound("Not found services of user")
		}
		return nil, model.InternalServerError(err.Error())
	}

	return userServices, nil
}

func (us UserServiceImpl) GetUserServiceByUserUuidAndServiceUuid(userUuid string, serviceUuid string) (*entity.UserService, *model.ErrorResBody) {
	userServices, err := us.UserRepository.FindUserServiceByUserUuidAndServiceUuid(userUuid, serviceUuid)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, model.NotFound("Not found service of user")
		}
		return nil, model.InternalServerError(err.Error())
	}

	return userServices, nil
}

func (us UserServiceImpl) GetUserByGroupUuid(groupUuid string) ([]*model.UserResponse, *model.ErrorResBody) {
	users, err := us.UserRepository.FindByGroupUuid(groupUuid)
	if err != nil {
		return nil, model.InternalServerError(err.Error())
	}

	var userResponse []*model.UserResponse
	for _, user := range users {
		userResponse = append(userResponse, &model.UserResponse{Uuid: user.Uuid.String(), Username: user.Username, Email: user.Email})
	}

	return userResponse, nil
}

func (us UserServiceImpl) GetUserPoliciesByUserUuid(userUuid string) []structure.UserPolicy {
	return us.EtcdClient.GetUserPolicy(userUuid)
}

func (us UserServiceImpl) GetUserGroupsByUserUuid(userUuid string) []structure.UserGroup {
	return us.EtcdClient.GetUserGroup(userUuid)
}

func (us UserServiceImpl) InsertUserGroup(userGroupEntity entity.UserGroup) (*entity.UserGroup, *model.ErrorResBody) {
	userGroupMd5 := md5.Sum(uuid.New().NodeID())
	userGroupEntity.InternalId = hex.EncodeToString(userGroupMd5[:])

	_, err := us.UserRepository.FindUserGroupByUserUuidAndGroupUuid(userGroupEntity.UserUuid.String(), userGroupEntity.GroupUuid.String())
	if err != nil {
		conflictErr := model.Conflict("This user already joins group")
		return nil, conflictErr
	}

	savedUserGroup, err := us.UserRepository.SaveUserGroup(userGroupEntity)
	if err != nil {
		if strings.Contains(err.Error(), "1062") {
			return nil, model.Conflict("Already exit data.")
		}
		return nil, model.InternalServerError(err.Error())
	}

	return savedUserGroup, nil
}

func (us UserServiceImpl) InsertUser(user entity.User) (*entity.User, *model.ErrorResBody) {
	uid := uuid.New()
	uidMd5 := md5.Sum(uid.NodeID())
	user.InternalId = hex.EncodeToString(uidMd5[:])
	user.Uuid = uid
	user.Password = us.EncryptPw(user.Password)

	savedUser, err := us.UserRepository.SaveUser(user)
	if err != nil {
		if strings.Contains(err.Error(), "1062") {
			return nil, model.Conflict("Already exit data.")
		}
		return nil, model.InternalServerError(err.Error())
	}

	return savedUser, nil
}

func (us UserServiceImpl) InsertUserWithUserService(user entity.User, userService entity.UserService) (*entity.User, *model.ErrorResBody) {
	uid := uuid.New()
	uidMd5 := md5.Sum(uid.NodeID())
	user.InternalId = hex.EncodeToString(uidMd5[:])
	user.Uuid = uid
	user.Password = us.EncryptPw(user.Password)

	userServiceIdMd5 := md5.Sum(uuid.New().NodeID())
	userService.InternalId = hex.EncodeToString(userServiceIdMd5[:])
	savedUser, err := us.UserRepository.SaveWithUserService(user, userService)
	if err != nil {
		if strings.Contains(err.Error(), "1062") {
			return nil, model.Conflict("Already exit service data.")
		}
		return nil, model.InternalServerError(err.Error())
	}

	return savedUser, nil
}

func (us UserServiceImpl) InsertUserService(userServiceEntity entity.UserService) (*entity.UserService, *model.ErrorResBody) {
	userServiceMd5 := md5.Sum(uuid.New().NodeID())
	userServiceEntity.InternalId = hex.EncodeToString(userServiceMd5[:])

	_, err := us.UserRepository.FindUserServiceByUserUuidAndServiceUuid(userServiceEntity.UserUuid.String(), userServiceEntity.ServiceUuid.String())
	if err != nil {
		return nil, model.Conflict("Already the user has this service account")
	}

	savedUserService, err := us.UserRepository.SaveUserService(userServiceEntity)
	if err != nil {
		if strings.Contains(err.Error(), "1062") {
			return nil, model.Conflict("Already exit data.")
		} else if strings.Contains(err.Error(), "1452") {
			return nil, model.BadRequest("Not register relational id.")
		} else {
			return nil, model.InternalServerError(err.Error())
		}
	}

	return savedUserService, nil
}

func (us UserServiceImpl) UpdateUser(user entity.User) (*entity.User, *model.ErrorResBody) {
	uidMd5 := md5.Sum(user.Uuid.NodeID())
	user.InternalId = hex.EncodeToString(uidMd5[:])
	user.Password = us.EncryptPw(user.Password)

	updatedUser, err := us.UserRepository.UpdateUser(user)
	if err != nil {
		return nil, model.InternalServerError(err.Error())
	}

	return updatedUser, nil
}
