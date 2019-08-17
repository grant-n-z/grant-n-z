package repository

import (
	"strings"

	"github.com/jinzhu/gorm"

	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepositoryImpl{
		Db: db,
	}
}

func (uri UserRepositoryImpl) FindById(id int) (*entity.User, *model.ErrorResponse) {
	user := entity.User{}
	if err := uri.Db.Where("id = ?", id).Find(&user).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError()
	}

	return &user, nil
}

func (uri UserRepositoryImpl) FindByEmail(email string) (*entity.User, *model.ErrorResponse) {
	user := entity.User{}
	if err := uri.Db.Where("email = ?", email).Find(&user).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError()
	}

	return &user, nil
}

func (uri UserRepositoryImpl) Save(user entity.User) (*entity.User, *model.ErrorResponse) {
	if err := uri.Db.Create(&user).Error; err != nil {
		log.Logger.Warn(err.Error())
		if strings.Contains(err.Error(), "1062") {
			return nil, model.Conflict("Already exit data.")
		}

		return nil, model.InternalServerError()
	}

	return &user, nil
}

func (uri UserRepositoryImpl) Update(user entity.User) *entity.User {
	if err := uri.Db.Update(&user).Error; err != nil {
		return nil
	}

	return &user
}
