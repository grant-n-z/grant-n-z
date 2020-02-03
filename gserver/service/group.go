package service

import (
	"github.com/google/uuid"

	"github.com/tomoyane/grant-n-z/gserver/common/constant"
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

	// Get group by id
	GetGroupById(id int) (*entity.Group, *model.ErrorResBody)

	// Get group that has the user
	GetGroupOfUser() ([]*entity.Group, *model.ErrorResBody)

	// Insert group
	InsertGroupWithRelationalData(group entity.Group) (*entity.Group, *model.ErrorResBody)
}

// GroupService struct
type GroupServiceImpl struct {
	groupRepository      data.GroupRepository
	roleRepository       data.RoleRepository
	permissionRepository data.PermissionRepository
}

// Get Policy instance.
// If use singleton pattern, call this instance method
func GetGroupServiceInstance() GroupService {
	if gsInstance == nil {
		gsInstance = NewGroupService()
	}
	return gsInstance
}

// Constructor
func NewGroupService() GroupService {
	log.Logger.Info("New `GroupService` instance")
	return GroupServiceImpl{
		groupRepository:      data.GetGroupRepositoryInstance(driver.Db),
		roleRepository:       data.GetRoleRepositoryInstance(driver.Db),
		permissionRepository: data.GetPermissionRepositoryInstance(driver.Db),
	}
}

func (gs GroupServiceImpl) GetGroups() ([]*entity.Group, *model.ErrorResBody) {
	return gs.groupRepository.FindAll()
}

func (gs GroupServiceImpl) GetGroupById(id int) (*entity.Group, *model.ErrorResBody) {
	return gs.groupRepository.FindById(id)
}

func (gs GroupServiceImpl) GetGroupOfUser() ([]*entity.Group, *model.ErrorResBody) {
	return gs.groupRepository.FindGroupsByUserId(ctx.GetUserId().(int))
}

func (gs GroupServiceImpl) InsertGroupWithRelationalData(group entity.Group) (*entity.Group, *model.ErrorResBody) {
	group.Uuid = uuid.New()

	// TODO: Cache role
	role, err := gs.roleRepository.FindByName(constant.AdminRole)
	if err != nil {
		log.Logger.Info("Failed to get role for insert groups process")
		return nil, model.InternalServerError()
	}

	// TODO: Cache permission
	permission, err := gs.permissionRepository.FindByName(constant.AdminRole)
	if err != nil {
		log.Logger.Info("Failed to get permission for insert groups process")
		return nil, model.InternalServerError()
	}

	serviceId := ctx.GetServiceId().(int)
	// New ServiceGroup
	serviceGroup := entity.ServiceGroup{
		GroupId:   group.Id,
		ServiceId: serviceId,
	}

	// New UserGroup
	userId := ctx.GetUserId().(int)
	userGroup := entity.UserGroup{
		UserId:  userId,
		GroupId: group.Id,
	}

	// New GroupPermission
	groupPermission := entity.GroupPermission{
		PermissionId: permission.Id,
		GroupId:      group.Id,
	}

	// New GroupRole
	groupRole := entity.GroupRole{
		RoleId:  role.Id,
		GroupId: group.Id,
	}

	// New Policy
	policy := entity.Policy{
		Name:         constant.AdminPolicy,
		RoleId:       role.Id,
		PermissionId: permission.Id,
		ServiceId:    serviceId,
		UserGroupId:  userGroup.Id,
	}

	return gs.groupRepository.SaveWithRelationalData(group, *role, *permission, serviceGroup, userGroup, groupPermission, groupRole, policy)
}
