package service

import (
	"github.com/tomoyane/grant-n-z/gnz/cache"
	"github.com/tomoyane/grant-n-z/gnz/ctx"
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

	// Get policy by role idr
	GetPoliciesByRoleId(roleId int) ([]*entity.Policy, *model.ErrorResBody)

	// Get policies of user
	// The method uses request scope user id
	GetPoliciesOfUser() ([]model.PolicyResponse, *model.ErrorResBody)

	// Get policy by user_groups data
	GetPolicyByUserGroup(userId int, groupId int) (*entity.Policy, *model.ErrorResBody)

	// Get policy by id
	GetPolicyById(id int) (entity.Policy, *model.ErrorResBody)

	// Insert or update policy
	UpdatePolicy(policy entity.Policy) (*entity.Policy, *model.ErrorResBody)
}

// PolicyService struct
type PolicyServiceImpl struct {
	EtcdClient           cache.EtcdClient
	PolicyRepository     driver.PolicyRepository
	PermissionRepository driver.PermissionRepository
	RoleRepository       driver.RoleRepository
	ServiceRepository    driver.ServiceRepository
	GroupRepository      driver.GroupRepository
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
	}
}

func (ps PolicyServiceImpl) GetPolicies() ([]*entity.Policy, *model.ErrorResBody) {
	return ps.PolicyRepository.FindAll()
}

func (ps PolicyServiceImpl) GetPoliciesByRoleId(roleId int) ([]*entity.Policy, *model.ErrorResBody) {
	return ps.PolicyRepository.FindByRoleId(roleId)
}

func (ps PolicyServiceImpl) GetPoliciesOfUser() ([]model.PolicyResponse, *model.ErrorResBody) {
	userGroupPolicies, err := ps.GroupRepository.FindGroupWithUserWithPolicyGroupsByUserId(ctx.GetUserId().(int))
	if err != nil {
		return nil, err
	}

	var policyResponses []model.PolicyResponse
	for _, ugp := range userGroupPolicies {
		if ugp.ServiceId == ctx.GetServiceId() {
			role := ps.EtcdClient.GetRole(ugp.Policy.RoleId)
			if role == nil {
				masterRole, err := ps.RoleRepository.FindById(ugp.Policy.RoleId)
				if err != nil {
					return nil, err
				}
				role = masterRole
			}

			permission := ps.EtcdClient.GetPermission(ugp.Policy.PermissionId)
			if permission == nil {
				masterPermission, err := ps.PermissionRepository.FindById(ugp.Policy.PermissionId)
				if err != nil {
					return nil, err
				}
				permission = masterPermission
			}

			service := ps.EtcdClient.GetService(ugp.Policy.ServiceId)
			if service == nil {
				masterService, err := ps.ServiceRepository.FindById(ugp.Policy.ServiceId)
				if err != nil {
					return nil, err
				}
				service = masterService
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
	}

	return policyResponses, nil
}

func (ps PolicyServiceImpl) GetPolicyByUserGroup(userId int, groupId int) (*entity.Policy, *model.ErrorResBody) {
	groupWithPolicy, err := ps.GroupRepository.FindGroupWithPolicyByUserIdAndGroupId(userId, groupId)
	if err != nil {
		return nil, err
	}

	return &groupWithPolicy.Policy, nil
}

func (ps PolicyServiceImpl) GetPolicyById(id int) (entity.Policy, *model.ErrorResBody) {
	if id == 0 {
		return entity.Policy{}, nil
	}

	cachePolicy := ps.EtcdClient.GetPolicy(id)
	if cachePolicy != nil {
		return *cachePolicy, nil
	}

	policy, err := ps.PolicyRepository.FindById(id)
	if err != nil {
		return policy, err
	}

	return policy, nil
}

func (ps PolicyServiceImpl) UpdatePolicy(policy entity.Policy) (*entity.Policy, *model.ErrorResBody) {
	updatedPolicy, err := ps.PolicyRepository.Update(policy)
	if err != nil {
		return nil, err
	}
	ps.EtcdClient.SetPolicy(*updatedPolicy, 0)
	return updatedPolicy, nil
}
