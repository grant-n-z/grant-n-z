package service

import (
	"github.com/tomoyane/grant-n-z/gserver/common/driver"
	"github.com/tomoyane/grant-n-z/gserver/data"
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

var ugsInstance UserGroupService

type UserGroupService interface {
	// Get UserGroup by user_id and group_id
	GetUserGroupByUserIdAndGroupId(userId int, groupId int) (*entity.UserGroup, *model.ErrorResBody)

	// Insert UserGroup
	InsertUserGroup(userGroup entity.UserGroup) (*entity.UserGroup, *model.ErrorResBody)
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
	return UserGroupServiceImpl{userGroupRepository: data.GetUserGroupRepositoryInstance(driver.Db)}
}

func (ugs UserGroupServiceImpl) GetUserGroupByUserIdAndGroupId(userId int, groupId int) (*entity.UserGroup, *model.ErrorResBody) {
	return ugs.userGroupRepository.FindGroupByUserIdAndGroupId(userId, groupId)
}

func (ugs UserGroupServiceImpl) InsertUserGroup(userGroupEntity entity.UserGroup) (*entity.UserGroup, *model.ErrorResBody) {
	userGroup, err := ugs.GetUserGroupByUserIdAndGroupId(userGroupEntity.UserId, userGroupEntity.GroupId)
	if err != nil || userGroup != nil {
		conflictErr := model.Conflict("This user already joins group")
		return nil, conflictErr
	}
	return ugs.userGroupRepository.Save(userGroupEntity)
}
