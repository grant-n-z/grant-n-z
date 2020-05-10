package driver

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"

	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
)

var urInstance UserRepository

type UserRepository interface {
	// Find User by user id
	FindById(id int) (*entity.User, *model.ErrorResBody)

	// Find User by user email
	FindByEmail(email string) (*entity.User, *model.ErrorResBody)

	// Find User and operator policy by user email
	FindWithOperatorPolicyByEmail(email string) (*model.UserWithOperatorPolicy, *model.ErrorResBody)

	// Find User and UserService and service by user email
	FindWithUserServiceWithServiceByEmail(email string) (*model.UserWithUserServiceWithService, *model.ErrorResBody)

	// Find UserGroup by user_id and group_id
	FindUserGroupByUserIdAndGroupId(userId int, groupId int) (*entity.UserGroup, *model.ErrorResBody)

	// Find all UserService
	FindUserServices() ([]*entity.UserService, *model.ErrorResBody)

	// Find all UserService with offset and limit
	FindUserServicesOffSetAndLimit(offset int, limit int) ([]*entity.UserService, *model.ErrorResBody)

	// Find all UserGroup with offset and limit
	FindUserGroupsOffSetAndLimit(offset int, limit int) ([]*entity.UserGroup, *model.ErrorResBody)

	// Find UserService by user_id and service_id
	FindUserServiceByUserIdAndServiceId(userId int, serviceId int) (*entity.UserService, *model.ErrorResBody)

	// Insert user_group data
	SaveUserGroup(userGroup entity.UserGroup) (*entity.UserGroup, *model.ErrorResBody)

	// Save User
	SaveUser(user entity.User) (*entity.User, *model.ErrorResBody)

	// Save User and user service
	SaveWithUserService(user entity.User, userService entity.UserService) (*entity.User, *model.ErrorResBody)

	// Save UserService
	SaveUserService(userService entity.UserService) (*entity.UserService, *model.ErrorResBody)

	// Update User
	UpdateUser(user entity.User) (*entity.User, *model.ErrorResBody)
}

// UserRepository struct
type UserRepositoryImpl struct {
	Connection *gorm.DB
}

// Get Policy instance.
// If use singleton pattern, call this instance method
func GetUserRepositoryInstance() UserRepository {
	if urInstance == nil {
		urInstance = NewUserRepository()
	}
	return urInstance
}

// Constructor
func NewUserRepository() UserRepository {
	log.Logger.Info("New `UserRepository` instance")
	return UserRepositoryImpl{Connection: connection}
}

func (uri UserRepositoryImpl) FindById(id int) (*entity.User, *model.ErrorResBody) {
	var user entity.User
	if err := uri.Connection.Where("id = ?", id).Find(&user).Error; err != nil {
		log.Logger.Warn(err.Error())
		if strings.Contains(err.Error(), "record not found") {
			return nil, model.NotFound()
		}

		return nil, model.InternalServerError()
	}

	return &user, nil
}

func (uri UserRepositoryImpl) FindByEmail(email string) (*entity.User, *model.ErrorResBody) {
	var user entity.User
	if err := uri.Connection.Where("email = ?", email).Find(&user).Error; err != nil {
		log.Logger.Warn(err.Error())
		if strings.Contains(err.Error(), "record not found") {
			return nil, model.NotFound()
		}

		return nil, model.InternalServerError()
	}

	return &user, nil
}

func (uri UserRepositoryImpl) FindWithOperatorPolicyByEmail(email string) (*model.UserWithOperatorPolicy, *model.ErrorResBody) {
	var uwo model.UserWithOperatorPolicy

	if err := uri.Connection.Table(entity.UserTable.String()).
		Select("*").
		Joins(fmt.Sprintf("LEFT JOIN %s ON %s.%s = %s.%s",
			entity.OperatorPolicyTable.String(),
			entity.UserTable.String(),
			entity.UserId,
			entity.OperatorPolicyTable.String(),
			entity.OperatorPolicyUserId)).
		Where(fmt.Sprintf("%s.%s = ?",
			entity.UserTable.String(),
			entity.UserEmail), email).
		Scan(&uwo).Error; err != nil {

		log.Logger.Warn(err.Error())
		return nil, model.InternalServerError()
	}

	return &uwo, nil
}

func (uri UserRepositoryImpl) FindWithUserServiceWithServiceByEmail(email string) (*model.UserWithUserServiceWithService, *model.ErrorResBody) {
	var uus model.UserWithUserServiceWithService

	if err := uri.Connection.Table(entity.UserTable.String()).
		Select("*").
		Joins(fmt.Sprintf("LEFT JOIN %s ON %s.%s = %s.%s",
			entity.UserServiceTable.String(),
			entity.UserTable.String(),
			entity.UserId,
			entity.UserServiceTable.String(),
			entity.UserServiceUserId)).
		Joins(fmt.Sprintf("LEFT JOIN %s ON %s.%s = %s.%s",
			entity.ServiceTable.String(),
			entity.UserServiceTable.String(),
			entity.UserServiceServiceId,
			entity.ServiceTable.String(),
			entity.ServiceId)).
		Where(fmt.Sprintf("%s.%s = ?",
			entity.UserTable.String(),
			entity.UserEmail), email).
		Scan(&uus).Error; err != nil {

		log.Logger.Warn(err.Error())
		if strings.Contains(err.Error(), "record not found") {
			return nil, model.NotFound("Not found user")
		}
		return nil, model.InternalServerError()
	}

	return &uus, nil
}

func (uri UserRepositoryImpl) FindUserGroupByUserIdAndGroupId(userId int, groupId int) (*entity.UserGroup, *model.ErrorResBody) {
	var userGroup entity.UserGroup
	if err := uri.Connection.Where("user_id = ? AND group_id = ?", userId, groupId).First(&userGroup).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError(err.Error())
	}

	return &userGroup, nil
}

func (uri UserRepositoryImpl) FindUserServices() ([]*entity.UserService, *model.ErrorResBody) {
	var userServices []*entity.UserService
	if err := uri.Connection.Find(&userServices).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError()
	}

	return userServices, nil
}

func (uri UserRepositoryImpl) FindUserServicesOffSetAndLimit(offset int, limit int) ([]*entity.UserService, *model.ErrorResBody) {
	var userServices []*entity.UserService
	if err := uri.Connection.Limit(limit).Offset(offset).Find(&userServices).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError()
	}

	return userServices, nil
}

func (uri UserRepositoryImpl) FindUserGroupsOffSetAndLimit(offset int, limit int) ([]*entity.UserGroup, *model.ErrorResBody) {
	var userGroups []*entity.UserGroup
	if err := uri.Connection.Limit(limit).Offset(offset).Find(&userGroups).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError()
	}

	return userGroups, nil
}

func (uri UserRepositoryImpl) FindUserServiceByUserIdAndServiceId(userId int, serviceId int) (*entity.UserService, *model.ErrorResBody) {
	var userService entity.UserService
	if err := uri.Connection.Where("user_id = ? AND service_id = ?", userId, serviceId).Find(&userService).Error; err != nil {
		log.Logger.Warn(err.Error())
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError()
	}

	return &userService, nil
}

func (uri UserRepositoryImpl) SaveUserGroup(userGroup entity.UserGroup) (*entity.UserGroup, *model.ErrorResBody) {
	if err := uri.Connection.Save(&userGroup).Error; err != nil {
		log.Logger.Warn(err.Error())
		if strings.Contains(err.Error(), "1062") {
			return nil, model.Conflict("Already exit data.")
		}

		return nil, model.InternalServerError()
	}

	return &userGroup, nil
}

func (uri UserRepositoryImpl) SaveUser(user entity.User) (*entity.User, *model.ErrorResBody) {
	if err := uri.Connection.Create(&user).Error; err != nil {
		log.Logger.Warn(err.Error())
		if strings.Contains(err.Error(), "1062") {
			return nil, model.Conflict("Already exit data.")
		}

		return nil, model.InternalServerError()
	}

	return &user, nil
}

func (uri UserRepositoryImpl) SaveWithUserService(user entity.User, userService entity.UserService) (*entity.User, *model.ErrorResBody) {
	tx := uri.Connection.Begin()

	if err := tx.Create(&user).Error; err != nil {
		log.Logger.Warn(err.Error())
		tx.Rollback()
		if strings.Contains(err.Error(), "1062") {
			return nil, model.Conflict("Already exit user data.")
		}

		return nil, model.InternalServerError()
	}

	userService.UserId = user.Id
	if err := tx.Create(&userService).Error; err != nil {
		log.Logger.Warn(err.Error())
		tx.Rollback()
		if strings.Contains(err.Error(), "1062") {
			return nil, model.Conflict("Already exit service data.")
		}

		return nil, model.InternalServerError()
	}

	tx.Commit()
	return &user, nil
}

func (uri UserRepositoryImpl) UpdateUser(user entity.User) (*entity.User, *model.ErrorResBody) {
	if err := uri.Connection.Save(&user).Error; err != nil {
		log.Logger.Warn(err.Error())
		return nil, model.InternalServerError()
	}

	return &user, nil
}

func (uri UserRepositoryImpl) SaveUserService(userService entity.UserService) (*entity.UserService, *model.ErrorResBody) {
	if err := uri.Connection.Create(&userService).Error; err != nil {
		log.Logger.Warn(err.Error())
		if strings.Contains(err.Error(), "1062") {
			return nil, model.Conflict("Already exit data.")
		} else if strings.Contains(err.Error(), "1452") {
			return nil, model.BadRequest("Not register relational id.")
		}

		return nil, model.InternalServerError()
	}

	return &userService, nil
}
