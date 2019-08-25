package repository

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"

	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

var urInstance UserRepository

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func GetUserRepositoryInstance(db *gorm.DB) UserRepository {
	if urInstance == nil {
		urInstance = NewUserRepository(db)
	}
	return urInstance
}

func NewUserRepository(db *gorm.DB) UserRepository {
	log.Logger.Info("New `UserRepository` instance")
	log.Logger.Info("Inject `gorm.DB` to `UserRepository`")
	return UserRepositoryImpl{
		Db: db,
	}
}

func (uri UserRepositoryImpl) FindById(id int) (*entity.User, *model.ErrorResponse) {
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

func (uri UserRepositoryImpl) FindByEmail(email string) (*entity.User, *model.ErrorResponse) {
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

func (uri UserRepositoryImpl) FindUserWithRoleByEmail(email string) (*model.UserOperatorMemberRole, *model.ErrorResponse) {
	var userOperatorMemberRole model.UserOperatorMemberRole
	if err := uri.Db.Table(entity.User{}.TableName()).
		Select("*").
		Joins(fmt.Sprintf("LEFT JOIN %s ON %s.%s = %s.%s",
			entity.OperatorMemberRole{}.TableName(),
			entity.User{}.TableName(),
			entity.UserId,
			entity.OperatorMemberRole{}.TableName(),
			entity.OperatorMemberRoleUserId)).
		Where(fmt.Sprintf("%s.%s = ?",
			entity.User{}.TableName(),
			entity.UserEmail), email).
		Find(&userOperatorMemberRole).Error; err != nil {

		log.Logger.Warn(err.Error())
		return nil, model.InternalServerError()
	}
	return &userOperatorMemberRole, nil
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

func (uri UserRepositoryImpl) SaveUserWithUserService(user entity.User, userService *entity.UserService) (*entity.User, *model.ErrorResponse) {
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

func (uri UserRepositoryImpl) Update(user entity.User) (*entity.User, *model.ErrorResponse) {
	if err := uri.Db.Save(&user).Error; err != nil {
		log.Logger.Warn(err.Error())
		return nil, model.InternalServerError(	)
	}

	return &user, nil
}
