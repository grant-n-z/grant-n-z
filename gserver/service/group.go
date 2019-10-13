package service

import (
	"strings"

	"github.com/tomoyane/grant-n-z/gserver/common/driver"
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
	"github.com/tomoyane/grant-n-z/gserver/data"
)

var gsInstance GroupService

type GroupService interface {
	Get(queryParam string) (interface{}, *model.ErrorResBody)

	GetGroups() ([]*entity.Group, *model.ErrorResBody)

	GetGroupByName(name string) (*entity.Group, *model.ErrorResBody)

	InsertGroup(group *entity.Group) (*entity.Group, *model.ErrorResBody)
}

type GroupServiceImpl struct {
	groupRepository data.GroupRepository
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
	return GroupServiceImpl{groupRepository: data.GetGroupRepositoryInstance(driver.Db)}
}

func (gs GroupServiceImpl) Get(queryParam string) (interface{}, *model.ErrorResBody) {
	if strings.EqualFold(queryParam, "") {
		return gs.GetGroups()
	}

	groupEntities, err := gs.GetGroupByName(queryParam)
	if err != nil {
		return nil, err
	}

	if groupEntities == nil {
		return []*entity.Group{}, nil
	}

	return groupEntities, nil
}

func (gs GroupServiceImpl) GetGroups() ([]*entity.Group, *model.ErrorResBody) {
	return gs.groupRepository.FindAll()
}

func (gs GroupServiceImpl) GetGroupByName(name string) (*entity.Group, *model.ErrorResBody) {
	return gs.groupRepository.FindByName(name)
}

func (gs GroupServiceImpl) InsertGroup(group *entity.Group) (*entity.Group, *model.ErrorResBody) {
	return gs.groupRepository.Save(*group)
}
