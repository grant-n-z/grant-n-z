package repository

import (
	"strings"

	"github.com/jinzhu/gorm"

	"github.com/tomoyane/grant-n-z/server/entity"
	"github.com/tomoyane/grant-n-z/server/log"
	"github.com/tomoyane/grant-n-z/server/model"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func (uri UserRepositoryImpl) FindById(id int) (*entity.User, *model.ErrorResponse) {
	user := entity.User{}
	if err := uri.Db.Where("id = ?", id).Find(&user).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError(err.Error())
	}

	return &user, nil
}

func (uri UserRepositoryImpl) FindByEmail(email string) (*entity.User, *model.ErrorResponse) {
	user := entity.User{}
	if err := uri.Db.Where("email = ?", email).Find(&user).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError(err.Error())
	}

	return &user, nil
}

func (uri UserRepositoryImpl) Save(user entity.User) (*entity.User, *model.ErrorResponse) {
	if err := uri.Db.Create(&user).Error; err != nil {
		errRes := model.Conflict(err.Error())
		if strings.Contains(err.Error(), "Duplicate entry") {
			log.Logger.Warn(errRes.ToJson(), errRes.Detail)
			return nil, model.Conflict(err.Error())
		}

		log.Logger.Error(errRes.ToJson(), errRes.Detail)
		return nil, model.InternalServerError(err.Error())
	}

	return &user, nil
}

func (uri UserRepositoryImpl) Update(user entity.User) *entity.User {
	if err := uri.Db.Update(&user).Error; err != nil {
		return nil
	}

	return &user
}
