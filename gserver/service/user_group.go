package service

import (
	"github.com/tomoyane/grant-n-z/gserver/common/driver"
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
	"github.com/tomoyane/grant-n-z/gserver/data"
)

var ugsInstance UserGroupService

type UserGroupService interface {
	InsertUserGroup(userGroup *entity.UserGroup) (*entity.UserGroup, *model.ErrorResBody)
}

type UserGroupServiceImpl struct {
	userGroupRepository data.UserGroupRepository
}

func GetUserGroupServiceInstance() UserGroupService {
	if ugsInstance == nil {
		ugsInstance = NewUserGroupService()
	}
	return ugsInstance
}

func NewUserGroupService() UserGroupService {
	log.Logger.Info("New `UserGroupService` instance")
	log.Logger.Info("Inject `UserGroupService` to `UserGroupService`")
	return UserGroupServiceImpl{userGroupRepository: data.GetUserGroupRepositoryInstance(driver.Db)}
}

func (ugs UserGroupServiceImpl) InsertUserGroup(userGroup *entity.UserGroup) (*entity.UserGroup, *model.ErrorResBody) {
	return ugs.userGroupRepository.Save(*userGroup)
}
