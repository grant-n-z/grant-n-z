package service

import (
	"github.com/tomoyane/grant-n-z/gserver/common/driver"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/repository"
)

var ugsInstance UserGroupService

type UserGroupServiceImpl struct {
	userGroupRepository repository.UserGroupRepository
}

func GetUSerServiceInstance() UserGroupService {
	if ugsInstance == nil {
		ugsInstance = NewUserGroupService()
	}
	return ugsInstance
}

func NewUserGroupService() UserGroupService {
	log.Logger.Info("New `UserGroupService` instance")
	log.Logger.Info("Inject `UserGroupService` to `UserGroupService`")
	return UserGroupServiceImpl{userGroupRepository: repository.GetUserGroupRepositoryInstance(driver.Db)}
}
