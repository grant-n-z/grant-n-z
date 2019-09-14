package service

import (
	"github.com/tomoyane/grant-n-z/gserver/common/driver"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/repository"
)

var gsInstance GroupService

type GroupServiceImpl struct {
	groupRepository repository.GroupRepository
}

func GetGroupServiceInstance() GroupService {
	if gsInstance == nil {
		gsInstance = NewGroupService()
	}
	return gsInstance
}

func NewGroupService() GroupService {
	log.Logger.Info("New `GroupService` instance")
	log.Logger.Info("Inject `GroupRepository` to `GroupService`")
	return GroupServiceImpl{groupRepository: repository.GetGroupRepositoryInstance(driver.Db)}
}
