package service

import (
	"crypto/md5"
	"encoding/hex"
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
			return []*entity.Group{}, nil
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
			return []*entity.Group{}, nil
		}
		return nil, model.InternalServerError(err.Error())
	}

	return groups, nil
}

func (gs GroupServiceImpl) InsertGroupWithRelationalData(group entity.Group, uUuid string, secret string) (*entity.Group, *model.ErrorResBody) {
	gid := uuid.New()
	gidMd5 := md5.Sum(gid.NodeID())
	group.Uuid = gid
	group.InternalId = hex.EncodeToString(gidMd5[:])

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
			return nil, model.BadRequest("Invalid secret")
		}
		return nil, model.InternalServerError(err.Error())
	}

	userUuid, _ := uuid.FromBytes([]byte(uUuid))

	// New ServiceGroup
	serviceGroupIdMd5 := md5.Sum(uuid.New().NodeID())
	serviceGroup := entity.ServiceGroup{
		InternalId:  hex.EncodeToString(serviceGroupIdMd5[:]),
		GroupUuid:   group.Uuid,
		ServiceUuid: ser.Uuid,
	}

	// New UserGroup
	userGroupIdMd5 := md5.Sum(uuid.New().NodeID())
	userGroup := entity.UserGroup{
		InternalId: hex.EncodeToString(userGroupIdMd5[:]),
		UserUuid:   userUuid,
		GroupUuid:  group.Uuid,
	}

	// New GroupPermission
	groupIdMd5 := md5.Sum(uuid.New().NodeID())
	groupPermission := entity.GroupPermission{
		InternalId:     hex.EncodeToString(groupIdMd5[:]),
		PermissionUuid: permission.Uuid,
		GroupUuid:      group.Uuid,
	}

	// New GroupRole
	groupRoleIdMd5 := md5.Sum(uuid.New().NodeID())
	groupRole := entity.GroupRole{
		InternalId: hex.EncodeToString(groupRoleIdMd5[:]),
		RoleUuid:   role.Uuid,
		GroupUuid:  group.Uuid,
	}

	// New Policy
	policyIdMd5 := md5.Sum(uuid.New().NodeID())
	policy := entity.Policy{
		InternalId:     hex.EncodeToString(policyIdMd5[:]),
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
