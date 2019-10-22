package service

import (
	"github.com/tomoyane/grant-n-z/gserver/common/ctx"
	"github.com/tomoyane/grant-n-z/gserver/common/driver"
	"github.com/tomoyane/grant-n-z/gserver/data"
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

var gsInstance GroupService

type GroupService interface {
	// Get all groups
	GetGroups() ([]*entity.Group, *model.ErrorResBody)

	// Get group that has the user
	GetGroupOfUser() ([]*entity.Group, *model.ErrorResBody)

	// Insert group
	InsertGroup(group *entity.Group) (*entity.Group, *model.ErrorResBody)
}

type GroupServiceImpl struct {
	groupRepository     data.GroupRepository
	userGroupRepository data.UserGroupRepository
}

func GetGroupServiceInstance() GroupService {
	if gsInstance == nil {
		gsInstance = NewGroupService()
	}
	return gsInstance
}

func NewGroupService() GroupService {
	log.Logger.Info("New `GroupService` instance")
	log.Logger.Info("Inject `GroupRepository`, `UserGroupRepository` to `GroupService`")
	return GroupServiceImpl{
		groupRepository:     data.GetGroupRepositoryInstance(driver.Db),
		userGroupRepository: data.GetUserGroupRepositoryInstance(driver.Db),
	}
}

func (gs GroupServiceImpl) GetGroups() ([]*entity.Group, *model.ErrorResBody) {
	return gs.groupRepository.FindAll()
}

func (gs GroupServiceImpl) GetGroupOfUser() ([]*entity.Group, *model.ErrorResBody) {
	if ctx.GetUserId().(int) == 0 {
		return nil, model.BadRequest("Required user id")
	}
	return gs.userGroupRepository.FindGroupsByUserId(ctx.GetUserId().(int))
}

func (gs GroupServiceImpl) InsertGroup(group *entity.Group) (*entity.Group, *model.ErrorResBody) {
	return gs.groupRepository.SaveWithUserGroupWithServiceGroup(*group)
}
