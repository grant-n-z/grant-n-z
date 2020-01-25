package service

import (
	"crypto/rsa"

	"github.com/tomoyane/grant-n-z/gserver/common/ctx"
	"github.com/tomoyane/grant-n-z/gserver/common/driver"
	"github.com/tomoyane/grant-n-z/gserver/data"
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

var (
	PrivateKey  *rsa.PrivateKey
	PublicKey   *rsa.PublicKey
	plsInstance PolicyService
)

type PolicyService interface {
	// Get all policy
	GetPolicies() ([]*entity.Policy, *model.ErrorResBody)

	// Get policy by role id
	GetPoliciesByRoleId(roleId int) ([]*entity.Policy, *model.ErrorResBody)

	// Get policies of user
	// The method uses request scope user id
	GetPoliciesOfUser() ([]entity.PolicyResponse, *model.ErrorResBody)

	// Get policy by user_groups data
	GetPolicyByUserGroup(userId int, groupId int) (*entity.Policy, *model.ErrorResBody)

	// Get policy by id
	GetPolicyById(id int) (entity.Policy, *model.ErrorResBody)

	// Insert policy
	InsertPolicy(policy *entity.Policy) (*entity.Policy, *model.ErrorResBody)
}

// PolicyService struct
type policyServiceImpl struct {
	policyRepository     data.PolicyRepository
	permissionRepository data.PermissionRepository
	roleRepository       data.RoleRepository
	serviceRepository    data.ServiceRepository
	userGroupRepository  data.UserGroupRepository
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
		policyRepository:     data.GetPolicyRepositoryInstance(driver.Db),
		permissionRepository: data.GetPermissionRepositoryInstance(driver.Db),
		roleRepository:       data.GetRoleRepositoryInstance(driver.Db),
		serviceRepository:    data.GetServiceRepositoryInstance(driver.Db),
		userGroupRepository:  data.GetUserGroupRepositoryInstance(driver.Db),
	}
}

func (ps policyServiceImpl) GetPolicies() ([]*entity.Policy, *model.ErrorResBody) {
	return ps.policyRepository.FindAll()
}

func (ps policyServiceImpl) GetPoliciesByRoleId(roleId int) ([]*entity.Policy, *model.ErrorResBody) {
	return ps.policyRepository.FindByRoleId(roleId)
}

func (ps policyServiceImpl) GetPoliciesOfUser() ([]entity.PolicyResponse, *model.ErrorResBody) {
	userGroupPolicies, err := ps.userGroupRepository.FindGroupWithUserWithPolicyGroupsByUserId(ctx.GetUserId().(int))
	if err != nil {
		return nil, err
	}

	var policyResponses []entity.PolicyResponse
	for _, ugp := range userGroupPolicies {
		if ugp.ServiceId == ctx.GetServiceId() {
			// TODO: Cache role, permission, service
			policyResponse := entity.NewPolicyResponse().
				SetName(&ugp.Policy.Name).
				SetRoleName(ps.roleRepository.FindNameById(ugp.Policy.RoleId)).
				SetPermissionName(ps.permissionRepository.FindNameById(ugp.Policy.PermissionId)).
				SetServiceName(ps.serviceRepository.FindNameById(ugp.Policy.ServiceId)).
				SetGroupName(&ugp.Group.Name).
				Build()

			policyResponses = append(policyResponses, policyResponse)
		}
	}

	return policyResponses, nil
}

func (ps policyServiceImpl) GetPolicyByUserGroup(userId int, groupId int) (*entity.Policy, *model.ErrorResBody) {
	groupWithPolicy, err := ps.userGroupRepository.FindGroupWithPolicyByUserIdAndGroupId(userId, groupId)
	if err != nil {
		return nil, err
	}

	return &groupWithPolicy.Policy, nil
}

func (ps policyServiceImpl) GetPolicyById(id int) (entity.Policy, *model.ErrorResBody) {
	policy, err := ps.policyRepository.FindById(id)
	if err != nil {
		return policy, err
	}

	return policy, nil
}

func (ps policyServiceImpl) InsertPolicy(policy *entity.Policy) (*entity.Policy, *model.ErrorResBody) {
	return ps.policyRepository.Save(*policy)
}
