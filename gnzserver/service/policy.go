package service

import (
	"strings"

	"github.com/tomoyane/grant-n-z/gnz/cache"
	"github.com/tomoyane/grant-n-z/gnz/cache/structure"
	"github.com/tomoyane/grant-n-z/gnz/driver"
	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
)

var (
	plsInstance PolicyService
)

type PolicyService interface {
	// Get all policy
	GetPolicies() ([]*entity.Policy, *model.ErrorResBody)

	// Get policy by role uuid
	GetPoliciesByRoleUuid(roleUuid string) ([]*entity.Policy, *model.ErrorResBody)

	// Get policies of user
	// The method uses request scope user id
	GetPoliciesByUser(userUuid string) ([]model.PolicyResponse, *model.ErrorResBody)

	// Get policy by user_groups data
	GetPolicyByUserGroup(userUuid string, groupUuid string) (*entity.Policy, *model.ErrorResBody)

	// Get policies by group uuid
	GetPoliciesByUserGroup(groupUuid string) ([]model.UserPolicyOnGroupResponse, *model.ErrorResBody)

	// Get policy by uuid
	GetPolicyByUuid(uuid string) (entity.Policy, *model.ErrorResBody)

	// Insert or update policy
	UpdatePolicy(policyRequest model.PolicyRequest, secret string, groupUuid string) (*entity.Policy, *model.ErrorResBody)
}

// PolicyService struct
type PolicyServiceImpl struct {
	EtcdClient           cache.EtcdClient
	PolicyRepository     driver.PolicyRepository
	PermissionRepository driver.PermissionRepository
	RoleRepository       driver.RoleRepository
	ServiceRepository    driver.ServiceRepository
	GroupRepository      driver.GroupRepository
	UserRepository       driver.UserRepository
}

// Get PolicyService instance.
// If use singleton pattern, call this instance method
func GetPolicyServiceInstance() PolicyService {
	if plsInstance == nil {
		plsInstance = NewPolicyService()
	}
	return plsInstance
}

// Constructor
func NewPolicyService() PolicyService {
	log.Logger.Info("New `PolicyService` instance")
	return PolicyServiceImpl{
		EtcdClient:           cache.GetEtcdClientInstance(),
		PolicyRepository:     driver.GetPolicyRepositoryInstance(),
		PermissionRepository: driver.GetPermissionRepositoryInstance(),
		RoleRepository:       driver.GetRoleRepositoryInstance(),
		ServiceRepository:    driver.GetServiceRepositoryInstance(),
		GroupRepository:      driver.GetGroupRepositoryInstance(),
		UserRepository:       driver.GetUserRepositoryInstance(),
	}
}

func (ps PolicyServiceImpl) GetPolicies() ([]*entity.Policy, *model.ErrorResBody) {
	policies, err := ps.PolicyRepository.FindAll()
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}
		return nil, model.InternalServerError(err.Error())
	}

	return policies, nil
}

func (ps PolicyServiceImpl) GetPoliciesByRoleUuid(roleUuid string) ([]*entity.Policy, *model.ErrorResBody) {
	policies, err := ps.PolicyRepository.FindByRoleUuid(roleUuid)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}
		return nil, model.InternalServerError(err.Error())
	}

	return policies, nil
}

func (ps PolicyServiceImpl) GetPoliciesByUser(userUuid string) ([]model.PolicyResponse, *model.ErrorResBody) {
	userGroupPolicies, err := ps.GroupRepository.FindGroupWithUserWithPolicyGroupsByUserUuid(userUuid)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil,nil
		}
		return nil, model.InternalServerError(err.Error())
	}

	var policyResponses []model.PolicyResponse
	for _, ugp := range userGroupPolicies {
		role, err := ps.RoleRepository.FindByUuid(ugp.Policy.RoleUuid.String())
		if err != nil {
			if strings.Contains(err.Error(), "record not found") {
				return nil, model.NotFound("Not found role that have policy")
			}
			return nil, model.InternalServerError(err.Error())
		}

		permission, err := ps.PermissionRepository.FindByUuid(ugp.Policy.PermissionUuid.String())
		if err != nil {
			if strings.Contains(err.Error(), "record not found") {
				return nil, model.NotFound("Not found permission that have policy")
			}
			return nil, model.InternalServerError(err.Error())
		}

		service, err := ps.ServiceRepository.FindByUuid(ugp.Policy.ServiceUuid.String())
		if err != nil {
			if strings.Contains(err.Error(), "record not found") {
				return nil, model.NotFound("Not found service that have policy")
			}
			return nil, model.InternalServerError(err.Error())
		}

		policyResponse := model.NewPolicyResponse().
			SetName(&ugp.Policy.Name).
			SetRoleName(&role.Name).
			SetPermissionName(&permission.Name).
			SetServiceName(&service.Name).
			SetGroupName(&ugp.Group.Name).
			Build()

		policyResponses = append(policyResponses, policyResponse)
	}

	return policyResponses, nil
}

func (ps PolicyServiceImpl) GetPolicyByUserGroup(userUuid string, groupUuid string) (*entity.Policy, *model.ErrorResBody) {
	groupWithPolicy, err := ps.GroupRepository.FindGroupWithPolicyByUserUuidAndGroupUuid(userUuid, groupUuid)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}
		return nil, model.InternalServerError(err.Error())
	}

	return &groupWithPolicy.Policy, nil
}

func (ps PolicyServiceImpl) GetPoliciesByUserGroup(groupUuid string) ([]model.UserPolicyOnGroupResponse, *model.ErrorResBody) {
	users, err := ps.UserRepository.FindByGroupUuid(groupUuid)
	if err != nil {
		return nil, model.InternalServerError(err.Error())
	}

	var userPolicies []model.UserPolicyOnGroupResponse
	for _, user := range users {
		policyResponse, err := ps.PolicyRepository.FindPolicyOfUserGroupByUserUuidAndGroupUuid(user.Uuid.String(), groupUuid)
		if err != nil {
			return nil, model.InternalServerError()
		}
		userPolicies = append(userPolicies, policyResponse)
	}

	return userPolicies, nil
}

func (ps PolicyServiceImpl) GetPolicyByUuid(uuid string) (entity.Policy, *model.ErrorResBody) {
	policy, err := ps.PolicyRepository.FindByUuid(uuid)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return policy, nil
		}

		return policy, model.InternalServerError(err.Error())
	}

	return policy, nil
}

func (ps PolicyServiceImpl) UpdatePolicy(policyRequest model.PolicyRequest, secret string, groupUuid string) (*entity.Policy, *model.ErrorResBody) {
	user, errUser := ps.UserRepository.FindByEmail(policyRequest.ToUserEmail)
	if errUser != nil {
		if strings.Contains(errUser.Error(), "record not found") {
			return nil, model.BadRequest("Not exist this user")
		}
		return nil, model.InternalServerError(errUser.Error())
	}

	userGroup, errGroup := ps.UserRepository.FindUserGroupByUserUuidAndGroupUuid(user.Uuid.String(), groupUuid)
	if errGroup != nil {
		if strings.Contains(errGroup.Error(), "record not found") {
			return nil, model.BadRequest("Not exist this user in group")
		}
		return nil, model.InternalServerError(errGroup.Error())
	}

	role, errRole := ps.RoleRepository.FindByUuid(policyRequest.RoleUuid)
	if errRole != nil {
		if strings.Contains(errRole.Error(), "record not found") {
			return nil, model.BadRequest("Not exist role")
		}
		return nil, model.InternalServerError(errRole.Error())
	}

	permission, errPermission := ps.PermissionRepository.FindByUuid(policyRequest.PermissionUuid)
	if errPermission != nil {
		if strings.Contains(errPermission.Error(), "record not found") {
			return nil, model.BadRequest("Not exist permission")
		}
		return nil, model.InternalServerError(errPermission.Error())
	}

	ser, errSer := ps.ServiceRepository.FindBySecret(secret)
	if errSer != nil {
		if strings.Contains(errSer.Error(), "record not found") {
			return nil, model.BadRequest("Not exist service")
		}
		return nil, model.InternalServerError(errSer.Error())
	}

	policy := entity.Policy{
		Name:           policyRequest.Name,
		RoleUuid:       role.Uuid,
		PermissionUuid: permission.Uuid,
		ServiceUuid:    ser.Uuid,
		UserGroupUuid:  userGroup.Uuid,
	}

	// Update RDBMS
	updatedPolicy, err := ps.PolicyRepository.Update(policy)
	if err != nil {
		if strings.Contains(err.Error(), "1062") {
			return nil, model.Conflict("Already exit data.")
		} else if strings.Contains(err.Error(), "1452") {
			return nil, model.BadRequest("Not register relational id.")
		} else {
			return nil, model.InternalServerError()
		}
	}

	// Update etcd
	userPolicies := ps.EtcdClient.GetUserPolicy(user.Uuid.String())
	var updatePolicies []structure.UserPolicy
	for _, userPolicy := range userPolicies {
		if updatedPolicy.ServiceUuid.String() == userPolicy.ServiceUuid {
			policy := structure.UserPolicy{
				ServiceUuid: updatedPolicy.ServiceUuid.String(),
				GroupUuid: groupUuid,
				RoleName: role.Name,
				PermissionName: permission.Name,
			}
			updatePolicies = append(updatePolicies, policy)
		} else {
			updatePolicies = append(updatePolicies, userPolicy)
		}
	}

	ps.EtcdClient.SetUserPolicy(user.Uuid.String(), updatePolicies)
	return updatedPolicy, nil
}
