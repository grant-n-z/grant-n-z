package data

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"

	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

var urInstance UserRepository

type UserRepository interface {
	// Find user by user id
	FindById(id int) (*entity.User, *model.ErrorResBody)

	// Find user by user email
	FindByEmail(email string) (*entity.User, *model.ErrorResBody)

	// Find user and operator policy by user email
	FindWithOperatorPolicyByEmail(email string) (*entity.UserWithOperatorPolicy, *model.ErrorResBody)

	// Find user and user service and service by user email
	FindWithUserServiceWithServiceByEmail(email string) (*entity.UserWithUserServiceWithService, *model.ErrorResBody)

	// Save user
	Save(user entity.User) (*entity.User, *model.ErrorResBody)

	// Save user and user service
	SaveWithUserService(user entity.User, userService entity.UserService) (*entity.User, *model.ErrorResBody)

	// Update
	Update(user entity.User) (*entity.User, *model.ErrorResBody)
}

// UserRepository struct
type UserRepositoryImpl struct {
	Db *gorm.DB
}

// Get Policy instance.
// If use singleton pattern, call this instance method
func GetUserRepositoryInstance(db *gorm.DB) UserRepository {
	if urInstance == nil {
		urInstance = NewUserRepository(db)
	}
	return urInstance
}

// Constructor
func NewUserRepository(db *gorm.DB) UserRepository {
	log.Logger.Info("New `UserRepository` instance")
	return UserRepositoryImpl{
		Db: db,
	}
}

func (uri UserRepositoryImpl) FindById(id int) (*entity.User, *model.ErrorResBody) {
	var user entity.User
	if err := uri.Db.Where("id = ?", id).Find(&user).Error; err != nil {
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
	if err := uri.Db.Where("email = ?", email).Find(&user).Error; err != nil {
		log.Logger.Warn(err.Error())
		if strings.Contains(err.Error(), "record not found") {
			return nil, model.NotFound()
		}

		return nil, model.InternalServerError()
	}

	return &user, nil
}

func (uri UserRepositoryImpl) FindWithOperatorPolicyByEmail(email string) (*entity.UserWithOperatorPolicy, *model.ErrorResBody) {
	var uwo entity.UserWithOperatorPolicy

	if err := uri.Db.Table(entity.UserTable.String()).
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

func (uri UserRepositoryImpl) FindWithUserServiceWithServiceByEmail(email string) (*entity.UserWithUserServiceWithService, *model.ErrorResBody) {
	var uus entity.UserWithUserServiceWithService

	if err := uri.Db.Table(entity.UserTable.String()).
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

func (uri UserRepositoryImpl) Save(user entity.User) (*entity.User, *model.ErrorResBody) {
	if err := uri.Db.Create(&user).Error; err != nil {
		log.Logger.Warn(err.Error())
		if strings.Contains(err.Error(), "1062") {
			return nil, model.Conflict("Already exit data.")
		}

		return nil, model.InternalServerError()
	}

	return &user, nil
}

func (uri UserRepositoryImpl) SaveWithUserService(user entity.User, userService entity.UserService) (*entity.User, *model.ErrorResBody) {
	tx := uri.Db.Begin()

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

func (uri UserRepositoryImpl) Update(user entity.User) (*entity.User, *model.ErrorResBody) {
	if err := uri.Db.Save(&user).Error; err != nil {
		log.Logger.Warn(err.Error())
		return nil, model.InternalServerError()
	}

	return &user, nil
}
