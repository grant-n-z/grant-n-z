package repository

import (
	"strings"

	"github.com/jinzhu/gorm"

	"github.com/tomoyane/grant-n-z/server/entity"
	"github.com/tomoyane/grant-n-z/server/log"
)

type UserServiceRepositoryImpl struct {
	Db *gorm.DB
}

func (usri UserServiceRepositoryImpl) FindAll() ([]*entity.UserService, *entity.ErrorResponse) {
	var userServices []*entity.UserService
	if err := usri.Db.Find(&userServices).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, entity.InternalServerError(err.Error())
	}

	return userServices, nil
}

func (usri UserServiceRepositoryImpl) FindByUserId(userId int) ([]*entity.UserService, *entity.ErrorResponse) {
	var userServices []*entity.UserService
	if err := usri.Db.Where("user_id = ?", userId).Find(&userServices).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, entity.InternalServerError(err.Error())
	}

	return userServices, nil
}

func (usri UserServiceRepositoryImpl) Save(userService entity.UserService) (*entity.UserService, *entity.ErrorResponse) {
	if err := usri.Db.Create(&userService).Error; err != nil {
		errRes := entity.Conflict(err.Error())
		if strings.Contains(err.Error(), "Duplicate entry") {
			log.Logger.Warn(errRes.ToJson(), errRes.Detail)
			return nil, entity.Conflict(err.Error())
		}

		log.Logger.Error(errRes.ToJson(), errRes.Detail)
		return nil, entity.InternalServerError(err.Error())
	}

	return &userService, nil
}
