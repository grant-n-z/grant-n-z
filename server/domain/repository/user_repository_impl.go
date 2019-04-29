package repository

import (
	"strings"

	"github.com/tomoyane/grant-n-z/server/config"
	"github.com/tomoyane/grant-n-z/server/domain/entity"
)

type UserRepositoryImpl struct {
}

func (u UserRepositoryImpl) Save(user entity.User) (*entity.User, *entity.ErrorResponse) {
	if err := config.Db.Create(&user).Error; err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return nil, entity.Conflict()
		}
		return nil, entity.InternalServerError()
	}

	return &user, nil
}

func (u UserRepositoryImpl) Update(user entity.User) *entity.User {
	if err := config.Db.Update(&user).Error; err != nil {
		return nil
	}

	return &user
}
