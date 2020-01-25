package data

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"

	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

var usrInstance UserServiceRepository

type UserServiceRepository interface {
	FindAll() ([]*entity.UserService, *model.ErrorResBody)

	FindServicesByUserId(userId int) ([]*entity.Service, *model.ErrorResBody)

	FindByUserIdAndServiceId(userId int, serviceId int) (*entity.UserService, *model.ErrorResBody)

	Save(userService entity.UserService) (*entity.UserService, *model.ErrorResBody)
}

type UserServiceRepositoryImpl struct {
	Db *gorm.DB
}

func GetUserServiceRepositoryInstance(db *gorm.DB) UserServiceRepository {
	if usrInstance == nil {
		usrInstance = NewUserServiceRepository(db)
	}
	return usrInstance
}

func NewUserServiceRepository(db *gorm.DB) UserServiceRepository {
	log.Logger.Info("New `UserServiceRepository` instance")
	return UserServiceRepositoryImpl{
		Db: db,
	}
}

func (usri UserServiceRepositoryImpl) FindAll() ([]*entity.UserService, *model.ErrorResBody) {
	var userServices []*entity.UserService
	if err := usri.Db.Find(&userServices).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError()
	}

	return userServices, nil
}

func (usri UserServiceRepositoryImpl) FindServicesByUserId(userId int) ([]*entity.Service, *model.ErrorResBody) {
	var services []*entity.Service

	if err := usri.Db.Table(entity.ServiceTable.String()).
		Select("*").
		Joins(fmt.Sprintf("LEFT JOIN %s ON %s.%s = %s.%s",
			entity.UserServiceTable.String(),
			entity.ServiceTable.String(),
			entity.ServiceId,
			entity.UserServiceTable.String(),
			entity.UserServiceServiceId)).
		Where(fmt.Sprintf("%s.%s = ?",
			entity.UserServiceTable.String(),
			entity.UserServiceUserId), userId).
		Scan(&services).Error; err != nil {

		log.Logger.Warn(err.Error())
		return nil, model.InternalServerError()
	}

	return services, nil
}

func (usri UserServiceRepositoryImpl) FindByUserIdAndServiceId(userId int, serviceId int) (*entity.UserService, *model.ErrorResBody) {
	var userService entity.UserService
	if err := usri.Db.Where("user_id = ? AND service_id = ?", userId, serviceId).Find(&userService).Error; err != nil {
		log.Logger.Warn(err.Error())
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError()
	}

	return &userService, nil
}

func (usri UserServiceRepositoryImpl) Save(userService entity.UserService) (*entity.UserService, *model.ErrorResBody) {
	if err := usri.Db.Create(&userService).Error; err != nil {
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
