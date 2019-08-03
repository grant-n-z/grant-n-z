package repository

import (
	"github.com/tomoyane/grant-n-z/server/model"
	"strings"

	"github.com/jinzhu/gorm"

	"github.com/tomoyane/grant-n-z/server/entity"
	"github.com/tomoyane/grant-n-z/server/log"
)

type UserServiceRepositoryImpl struct {
	Db *gorm.DB
}

func NewUserServiceRepository(db *gorm.DB) UserServiceRepository {
	return UserServiceRepositoryImpl {
		Db: db,
	}
}

func (usri UserServiceRepositoryImpl) FindAll() ([]*entity.UserService, *model.ErrorResponse) {
	var userServices []*entity.UserService
	if err := usri.Db.Find(&userServices).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError(err.Error())
	}

	return userServices, nil
}

func (usri UserServiceRepositoryImpl) FindById(id int) ([]*entity.UserService, *model.ErrorResponse) {
	var userServices []*entity.UserService
	if err := usri.Db.Where("id = ?", id).Find(&userServices).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError(err.Error())
	}

	return userServices, nil
}

func (usri UserServiceRepositoryImpl) FindByUserId(userId int) ([]*entity.UserService, *model.ErrorResponse) {
	var userServices []*entity.UserService
	if err := usri.Db.Where("user_id = ?", userId).Find(&userServices).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError(err.Error())
	}

	return userServices, nil
}

func (usri UserServiceRepositoryImpl) Save(userService entity.UserService) (*entity.UserService, *model.ErrorResponse) {
	if err := usri.Db.Create(&userService).Error; err != nil {
		errRes := model.Conflict(err.Error())
		if strings.Contains(err.Error(), "Duplicate entry") {
			log.Logger.Warn(errRes.ToJson(), errRes.Detail)
			return nil, model.Conflict(err.Error())
		}

		log.Logger.Error(errRes.ToJson(), errRes.Detail)
		return nil, model.InternalServerError(err.Error())
	}

	return &userService, nil
}
