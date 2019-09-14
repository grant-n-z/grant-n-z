package repository

import (
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

}

func (ugr UserGroupRepositoryImpl) FindUsersByGroupId(groupId int) ([]*entity.User, *model.ErrorResponse) {

}

func (ugr UserGroupRepositoryImpl) Save(userGroup entity.UserGroup) (*entity.Group, *model.ErrorResponse) {

}
