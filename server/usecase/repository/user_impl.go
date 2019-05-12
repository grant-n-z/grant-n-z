package repository

import (
	"strings"

	"github.com/jinzhu/gorm"

	"github.com/tomoyane/grant-n-z/server/entity"
	"github.com/tomoyane/grant-n-z/server/log"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func (uri UserRepositoryImpl) FindById(id int) (*entity.User, *entity.ErrorResponse) {
	user := entity.User{}
	if err := uri.Db.Where("id = ?", id).Find(&user).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, entity.InternalServerError(err.Error())
	}

	return &user, nil
}

func (uri UserRepositoryImpl) Save(user entity.User) (*entity.User, *entity.ErrorResponse) {
	if err := uri.Db.Create(&user).Error; err != nil {
		errRes := entity.Conflict(err.Error())
		if strings.Contains(err.Error(), "Duplicate entry") {
			log.Logger.Warn(errRes.ToJson(), errRes.Detail)
			return nil, entity.Conflict(err.Error())
		}

		log.Logger.Error(errRes.ToJson(), errRes.Detail)
		return nil, entity.InternalServerError(err.Error())
	}

	return &user, nil
}

func (uri UserRepositoryImpl) Update(user entity.User) *entity.User {
	if err := uri.Db.Update(&user).Error; err != nil {
		return nil
	}

	return &user
}
