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
type policyServiceImpl struct {
	etcdClient           cache.EtcdClient
	policyRepository     driver.PolicyRepository
	permissionRepository driver.PermissionRepository
	roleRepository       driver.RoleRepository
	serviceRepository    driver.ServiceRepository
	groupRepository      driver.GroupRepository
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
	return policyServiceImpl{
		etcdClient:           cache.GetEtcdClientInstance(),
		policyRepository:     driver.GetPolicyRepositoryInstance(),
		permissionRepository: driver.GetPermissionRepositoryInstance(),
		roleRepository:       driver.GetRoleRepositoryInstance(),
		serviceRepository:    driver.GetServiceRepositoryInstance(),
		groupRepository:      driver.GetGroupRepositoryInstance(),
	}
}

func (ps policyServiceImpl) GetPolicies() ([]*entity.Policy, *model.ErrorResBody) {
	return ps.policyRepository.FindAll()
}

func (ps policyServiceImpl) GetPoliciesByRoleId(roleId int) ([]*entity.Policy, *model.ErrorResBody) {
	return ps.policyRepository.FindByRoleId(roleId)
}

func (ps policyServiceImpl) GetPoliciesOfUser() ([]model.PolicyResponse, *model.ErrorResBody) {
	userGroupPolicies, err := ps.groupRepository.FindGroupWithUserWithPolicyGroupsByUserId(ctx.GetUserId().(int))
	if err != nil {
		return nil, err
	}

	var policyResponses []model.PolicyResponse
	for _, ugp := range userGroupPolicies {
		if ugp.ServiceId == ctx.GetServiceId() {
			role := ps.etcdClient.GetRole(ugp.Policy.RoleId)
			if role == nil {
				masterRole, err := ps.roleRepository.FindById(ugp.Policy.RoleId)
				if err != nil {
					return nil, err
				}
				role = masterRole
			}

			permission := ps.etcdClient.GetPermission(ugp.Policy.PermissionId)
			if permission == nil {
				masterPermission, err := ps.permissionRepository.FindById(ugp.Policy.PermissionId)
				if err != nil {
					return nil, err
				}
				permission = masterPermission
			}

			service := ps.etcdClient.GetService(ugp.Policy.ServiceId)
			if service == nil {
				masterService, err := ps.serviceRepository.FindById(ugp.Policy.ServiceId)
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

func (ps policyServiceImpl) GetPolicyByUserGroup(userId int, groupId int) (*entity.Policy, *model.ErrorResBody) {
	groupWithPolicy, err := ps.groupRepository.FindGroupWithPolicyByUserIdAndGroupId(userId, groupId)
	if err != nil {
		return nil, err
	}

	return &groupWithPolicy.Policy, nil
}

func (ps policyServiceImpl) GetPolicyById(id int) (entity.Policy, *model.ErrorResBody) {
	if id == 0 {
		return entity.Policy{}, nil
	}

	cachePolicy := ps.etcdClient.GetPolicy(id)
	if cachePolicy != nil {
		return *cachePolicy, nil
	}

	policy, err := ps.policyRepository.FindById(id)
	if err != nil {
		return policy, err
	}

	return policy, nil
}

func (ps policyServiceImpl) UpdatePolicy(policy entity.Policy) (*entity.Policy, *model.ErrorResBody) {
	return ps.policyRepository.Update(policy)
}
