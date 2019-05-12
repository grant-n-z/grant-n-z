package service

import (
	"strconv"
	"strings"

	"github.com/tomoyane/grant-n-z/server/config"
	"github.com/tomoyane/grant-n-z/server/entity"
	"github.com/tomoyane/grant-n-z/server/log"
	"github.com/tomoyane/grant-n-z/server/usecase/repository"
)

type policyServiceImpl struct {
	policyRepository     repository.PolicyRepository
	permissionRepository repository.PermissionRepository
	roleRepository       repository.RoleRepository
}

func NewPolicyService() PolicyService {
	log.Logger.Info("Inject `roleRepository`, `permissionRepository`, `roleRepository` to `PolicyService`")
	return policyServiceImpl{
		policyRepository:     repository.PolicyRepositoryImpl{Db: config.Db},
		permissionRepository: repository.PermissionRepositoryImpl{Db: config.Db},
		roleRepository:       repository.RoleRepositoryImpl{Db: config.Db},
	}
}

func (ps policyServiceImpl) Get(queryParam string) ([]*entity.Policy, *entity.ErrorResponse) {
	if strings.EqualFold(queryParam, "") {
		return ps.GetPolicies()
	}

	i, castErr := strconv.Atoi(queryParam)
	if castErr != nil {
		log.Logger.Warn("The role_id is only integer")
		return nil, entity.BadRequest(castErr.Error())
	}

	policyEntities, err := ps.GetPoliciesByRoleId(i)
	if err != nil {
		return nil, err
	}

	if policyEntities == nil {
		return []*entity.Policy{}, nil
	}

	return policyEntities, nil
}

func (ps policyServiceImpl) GetPolicies() ([]*entity.Policy, *entity.ErrorResponse) {
	return ps.policyRepository.FindAll()
}

func (ps policyServiceImpl) GetPoliciesByRoleId(roleId int) ([]*entity.Policy, *entity.ErrorResponse) {
	return ps.policyRepository.FindByRoleId(roleId)
}

func (ps policyServiceImpl) InsertPolicy(policy *entity.Policy) (*entity.Policy, *entity.ErrorResponse) {
	if permissionEntity, _ := ps.permissionRepository.FindById(policy.PermissionId); permissionEntity == nil {
		return nil, entity.BadRequest("Not found permission id")
	}

	if roleEntity, _ := ps.roleRepository.FindById(policy.RoleId); roleEntity == nil {
		return nil, entity.BadRequest("Not found role id")
	}

	return ps.policyRepository.Save(*policy)
}
