package repository

import (
	"strings"

	"github.com/tomoyane/grant-n-z/server/config"
	"github.com/tomoyane/grant-n-z/server/domain/entity"
	"github.com/tomoyane/grant-n-z/server/log"
)

type UserRepositoryImpl struct {
}

func (uri UserRepositoryImpl) Save(user entity.User) (*entity.User, *entity.ErrorResponse) {
	if err := config.Db.Create(&user).Error; err != nil {
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
	if err := config.Db.Update(&user).Error; err != nil {
		return nil
	}

	return &user
}
