package repository

import (
	"strings"

	"github.com/jinzhu/gorm"

	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

var ugrInstance UserGroupRepository

type UserGroupRepositoryImpl struct {
	Db *gorm.DB
}

func GetUserGroupRepositoryInstance(db *gorm.DB) UserGroupRepository {
	if ugrInstance == nil {
		ugrInstance = NewUserGroupRepository(db)
	}
	return ugrInstance
}

func NewUserGroupRepository(db *gorm.DB) UserGroupRepository {
	log.Logger.Info("New `UserGroupRepository` instance")
	log.Logger.Info("Inject `gorm.DB` to `UserGroupRepository`")
	return UserGroupRepositoryImpl{Db: db}
}

func (ugr UserGroupRepositoryImpl) FindGroupsByUserId(userId int) ([]*entity.Group, *model.ErrorResponse) {
	return nil, nil
}

func (ugr UserGroupRepositoryImpl) FindUsersByGroupId(groupId int) ([]*entity.User, *model.ErrorResponse) {
	return nil, nil
}

func (ugr UserGroupRepositoryImpl) Save(userGroup entity.UserGroup) (*entity.UserGroup, *model.ErrorResponse) {
	if err := ugr.Db.Create(&userGroup).Error; err != nil {
		log.Logger.Warn(err.Error())
		if strings.Contains(err.Error(), "1062") {
			return nil, model.Conflict("Already exit data.")
		} else if strings.Contains(err.Error(), "1452") {
			return nil, model.BadRequest("Not register relational id.")
		}

		return nil, model.InternalServerError()
	}

	return &userGroup, nil
}
