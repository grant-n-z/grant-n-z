package service

import (
	"strings"

	"github.com/google/uuid"

	"github.com/tomoyane/grant-n-z/gnz/cache"
	"github.com/tomoyane/grant-n-z/gnz/common"
	"github.com/tomoyane/grant-n-z/gnz/driver"
	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
)

var gsInstance GroupService

type GroupService interface {
	// Get all groups
	GetGroups() ([]*entity.Group, *model.ErrorResBody)

	// Get group by uuid
	GetGroupByUuid(uuid string) (*entity.Group, *model.ErrorResBody)

	// Get group that has the user
	GetGroupByUser(userUuid string) ([]*entity.Group, *model.ErrorResBody)

	// Insert group
	InsertGroupWithRelationalData(group entity.Group, userUuid string, secret string) (*entity.Group, *model.ErrorResBody)
}

// GroupService struct
type GroupServiceImpl struct {
	EtcdClient           cache.EtcdClient
	GroupRepository      driver.GroupRepository
	RoleRepository       driver.RoleRepository
	PermissionRepository driver.PermissionRepository
	ServiceRepository    driver.ServiceRepository
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
		EtcdClient:           cache.GetEtcdClientInstance(),
		GroupRepository:      driver.GetGroupRepositoryInstance(),
		RoleRepository:       driver.GetRoleRepositoryInstance(),
		PermissionRepository: driver.GetPermissionRepositoryInstance(),
		ServiceRepository:    driver.GetServiceRepositoryInstance(),
	}
}

func (gs GroupServiceImpl) GetGroups() ([]*entity.Group, *model.ErrorResBody) {
	groups, err := gs.GroupRepository.FindAll()
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}
		return nil, model.InternalServerError(err.Error())
	}

	return groups, nil
}

func (gs GroupServiceImpl) GetGroupByUuid(uuid string) (*entity.Group, *model.ErrorResBody) {
	group, err := gs.GroupRepository.FindByUuid(uuid)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, model.NotFound("Not found group")
		}
		return nil, model.InternalServerError(err.Error())
	}

	return group, nil
}

func (gs GroupServiceImpl) GetGroupByUser(userUuid string) ([]*entity.Group, *model.ErrorResBody) {
	groups, err := gs.GroupRepository.FindByUserUuid(userUuid)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}
		return nil, model.InternalServerError(err.Error())
	}

	return groups, nil
}

func (gs GroupServiceImpl) InsertGroupWithRelationalData(group entity.Group, uUuid string, secret string) (*entity.Group, *model.ErrorResBody) {
	group.Uuid = uuid.New()

	role, err := gs.RoleRepository.FindByName(common.AdminRole)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, model.NotFound("Not found relation role")
		}
		return nil, model.InternalServerError(err.Error())
	}

	permission, err := gs.PermissionRepository.FindByName(common.AdminPermission)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, model.NotFound("Not found relation permission")
		}
		return nil, model.InternalServerError(err.Error())
	}

	ser, err := gs.ServiceRepository.FindBySecret(secret)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, model.NotFound("Not found relation service")
		}
		return nil, model.InternalServerError(err.Error())
	}

	userUuid, _ := uuid.FromBytes([]byte(uUuid))

	// New ServiceGroup
	serviceGroup := entity.ServiceGroup{
		GroupUuid:   group.Uuid,
		ServiceUuid: ser.Uuid,
	}

	// New UserGroup
	userGroup := entity.UserGroup{
		UserUuid:  userUuid,
		GroupUuid: group.Uuid,
	}

	// New GroupPermission
	groupPermission := entity.GroupPermission{
		PermissionUuid: permission.Uuid,
		GroupUuid:      group.Uuid,
	}

	// New GroupRole
	groupRole := entity.GroupRole{
		RoleUuid:  role.Uuid,
		GroupUuid: group.Uuid,
	}

	// New Policy
	policy := entity.Policy{
		Name:           common.AdminPolicy,
		RoleUuid:       role.Uuid,
		PermissionUuid: permission.Uuid,
		ServiceUuid:    ser.Uuid,
		UserGroupUuid:  userGroup.Uuid,
	}

	savedData, err := gs.GroupRepository.SaveWithRelationalData(group, serviceGroup, userGroup, groupPermission, groupRole, policy)
	if err != nil {
		if strings.Contains(err.Error(), "1062") {
			return nil, model.Conflict("Already exit service group data.")
		}
		return nil, model.InternalServerError("Failed to save transaction")
	}

	return savedData, nil
}
