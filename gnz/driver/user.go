package driver

import (
	"fmt"
	"github.com/jinzhu/gorm"

	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
)

var urInstance UserRepository

type UserRepository interface {
	// FindByUuid
	// Find User by user id
	FindByUuid(uuid string) (*entity.User, error)

	// FindByEmail
	// Find User by user email
	FindByEmail(email string) (*entity.User, error)

	// FindByGroupUuid
	// Find User by group uuid
	FindByGroupUuid(groupUuid string) ([]*entity.User, error)

	// FindWithOperatorPolicyByEmail
	// Find User and operator policy by user email
	FindWithOperatorPolicyByEmail(email string) (*model.UserWithOperatorPolicy, error)

	// FindWithUserServiceWithServiceByEmail
	// Find User and UserService and service by user email
	FindWithUserServiceWithServiceByEmail(email string) (*model.UserWithUserServiceWithService, error)

	// FindUserGroupByUserUuidAndGroupUuid
	// Find UserGroup by user uuid and group uuid
	FindUserGroupByUserUuidAndGroupUuid(userUuid string, groupUuid string) (*entity.UserGroup, error)

	// FindUserServices
	// Find all UserService
	FindUserServices() ([]*entity.UserService, error)

	// FindUserServicesByUserUuid
	// Find UserService by user uuid
	FindUserServicesByUserUuid(userUuid string) ([]*entity.UserService, error)

	// FindUserServicesOffSetAndLimit
	// Find all UserService with offset and limit
	FindUserServicesOffSetAndLimit(offset int, limit int) ([]*entity.UserService, error)

	// FindUserGroupsOffSetAndLimit
	// Find all UserGroup with offset and limit
	FindUserGroupsOffSetAndLimit(offset int, limit int) ([]*entity.UserGroup, error)

	// FindUserServiceByUserUuidAndServiceUuid
	// Find UserService by user uuid and service uuid
	FindUserServiceByUserUuidAndServiceUuid(userUuid string, serviceUuid string) (*entity.UserService, error)

	// SaveUserGroup
	// Insert user_group data
	SaveUserGroup(userGroup entity.UserGroup) (*entity.UserGroup, error)

	// SaveUser
	// Save User
	SaveUser(user entity.User) (*entity.User, error)

	// SaveWithUserService
	// Save User and user service
	SaveWithUserService(user entity.User, userService entity.UserService) (*entity.User, error)

	// SaveUserService
	// Save UserService
	SaveUserService(userService entity.UserService) (*entity.UserService, error)

	// UpdateUser
	// Update User
	UpdateUser(user entity.User) (*entity.User, error)
}

// RdbmsUserRepository
// UserRepository struct
type RdbmsUserRepository struct {
	Connection *gorm.DB
}

// GetUserRepositoryInstance
// Get Policy instance.
// If use singleton pattern, call this instance method
func GetUserRepositoryInstance() UserRepository {
	if urInstance == nil {
		urInstance = NewUserRepository()
	}
	return urInstance
}

// NewUserRepository
// Constructor
func NewUserRepository() UserRepository {
	log.Logger.Info("New `UserRepository` instance")
	return RdbmsUserRepository{Connection: connection}
}

func (uri RdbmsUserRepository) FindByUuid(uuid string) (*entity.User, error) {
	var user entity.User
	if err := uri.Connection.Where("uuid = ?", uuid).Find(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (uri RdbmsUserRepository) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := uri.Connection.Where("email = ?", email).Find(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (uri RdbmsUserRepository) FindByGroupUuid(groupUuid string) ([]*entity.User, error) {
	var users []*entity.User

	target := entity.UserTable.String() + "." +
		entity.UserId.String() + "," +
		entity.UserTable.String() + "." +
		entity.UserUuid.String() + "," +
		entity.UserTable.String() + "." +
		entity.UserUsername.String() + "," +
		entity.UserTable.String() + "." +
		entity.UserEmail.String()

	if err := uri.Connection.Table(entity.UserGroupTable.String()).
		Select(target).
		Joins(fmt.Sprintf("LEFT JOIN %s ON %s.%s = %s.%s",
			entity.UserTable.String(),
			entity.UserGroupTable.String(),
			entity.UserGroupUserUuid.String(),
			entity.UserTable.String(),
			entity.UserUuid.String())).
		Where(fmt.Sprintf("%s.%s = ?",
			entity.UserGroupTable.String(),
			entity.UserGroupGroupUuid.String()), groupUuid).
		Scan(&users).Error; err != nil {

		return nil, err
	}

	return users, nil
}

func (uri RdbmsUserRepository) FindWithOperatorPolicyByEmail(email string) (*model.UserWithOperatorPolicy, error) {
	var uwo model.UserWithOperatorPolicy

	if err := uri.Connection.Table(entity.UserTable.String()).
		Select("*").
		Joins(fmt.Sprintf("LEFT JOIN %s ON %s.%s = %s.%s",
			entity.OperatorPolicyTable.String(),
			entity.OperatorPolicyTable.String(),
			entity.OperatorPolicyUserUuid.String(),
			entity.UserTable.String(),
			entity.UserUuid.String())).
		Joins(fmt.Sprintf("LEFT JOIN %s ON %s.%s = %s.%s",
			entity.RoleTable.String(),
			entity.RoleTable.String(),
			entity.RoleUuid.String(),
			entity.OperatorPolicyTable.String(),
			entity.OperatorPolicyRoleUuid.String())).
		Where(fmt.Sprintf("%s.%s = ?",
			entity.UserTable.String(),
			entity.UserEmail.String()), email).
		Scan(&uwo).Error; err != nil {

		return nil, err
	}

	return &uwo, nil
}

func (uri RdbmsUserRepository) FindWithUserServiceWithServiceByEmail(email string) (*model.UserWithUserServiceWithService, error) {
	var uus model.UserWithUserServiceWithService

	if err := uri.Connection.Table(entity.UserTable.String()).
		Select("*").
		Joins(fmt.Sprintf("LEFT JOIN %s ON %s.%s = %s.%s",
			entity.UserServiceTable.String(),
			entity.UserTable.String(),
			entity.UserUuid.String(),
			entity.UserServiceTable.String(),
			entity.UserServiceUserUuid.String())).
		Joins(fmt.Sprintf("LEFT JOIN %s ON %s.%s = %s.%s",
			entity.ServiceTable.String(),
			entity.UserServiceTable.String(),
			entity.UserServiceServiceUuid.String(),
			entity.ServiceTable.String(),
			entity.ServiceUuid.String())).
		Where(fmt.Sprintf("%s.%s = ?",
			entity.UserTable.String(),
			entity.UserEmail.String()), email).
		Scan(&uus).Error; err != nil {

		return nil, err
	}

	return &uus, nil
}

func (uri RdbmsUserRepository) FindUserGroupByUserUuidAndGroupUuid(userUuid string, groupUuid string) (*entity.UserGroup, error) {
	var userGroup entity.UserGroup
	if err := uri.Connection.Where("user_uuid = ? AND group_uuid = ?", userUuid, groupUuid).First(&userGroup).Error; err != nil {
		return nil, err
	}

	return &userGroup, nil
}

func (uri RdbmsUserRepository) FindUserServices() ([]*entity.UserService, error) {
	var userServices []*entity.UserService
	if err := uri.Connection.Find(&userServices).Error; err != nil {
		return nil, err
	}

	return userServices, nil
}

func (uri RdbmsUserRepository) FindUserServicesByUserUuid(userUuid string) ([]*entity.UserService, error) {
	var userServices []*entity.UserService
	if err := uri.Connection.Where("user_uuid = ?", userUuid).Find(&userServices).Error; err != nil {
		return nil, err
	}

	return userServices, nil
}

func (uri RdbmsUserRepository) FindUserServicesOffSetAndLimit(offset int, limit int) ([]*entity.UserService, error) {
	var userServices []*entity.UserService
	if err := uri.Connection.Limit(limit).Offset(offset).Find(&userServices).Error; err != nil {
		return nil, err
	}

	return userServices, nil
}

func (uri RdbmsUserRepository) FindUserGroupsOffSetAndLimit(offset int, limit int) ([]*entity.UserGroup, error) {
	var userGroups []*entity.UserGroup
	if err := uri.Connection.Limit(limit).Offset(offset).Find(&userGroups).Error; err != nil {
		return nil, err
	}

	return userGroups, nil
}

func (uri RdbmsUserRepository) FindUserServiceByUserUuidAndServiceUuid(userUuid string, serviceUuid string) (*entity.UserService, error) {
	var userService entity.UserService
	if err := uri.Connection.Where("user_uuid = ? AND service_uuid = ?", userUuid, serviceUuid).Find(&userService).Error; err != nil {
		return nil, err
	}

	return &userService, nil
}

func (uri RdbmsUserRepository) SaveUserGroup(userGroup entity.UserGroup) (*entity.UserGroup, error) {
	if err := uri.Connection.Save(&userGroup).Error; err != nil {
		return nil, err
	}

	return &userGroup, nil
}

func (uri RdbmsUserRepository) SaveUser(user entity.User) (*entity.User, error) {
	if err := uri.Connection.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (uri RdbmsUserRepository) SaveWithUserService(user entity.User, userService entity.UserService) (*entity.User, error) {
	tx := uri.Connection.Begin()

	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	userService.UserUuid = user.Uuid
	if err := tx.Create(&userService).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return &user, nil
}

func (uri RdbmsUserRepository) SaveUserService(userService entity.UserService) (*entity.UserService, error) {
	if err := uri.Connection.Create(&userService).Error; err != nil {
		return nil, err
	}

	return &userService, nil
}

func (uri RdbmsUserRepository) UpdateUser(user entity.User) (*entity.User, error) {
	if err := uri.Connection.Save(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
