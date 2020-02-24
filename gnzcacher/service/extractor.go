package service

import (
	"github.com/tomoyane/grant-n-z/gnz/driver"
	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
)

type ExtractorService interface {
	// Get policies for offset and limit
	GetPolicies(offset int, limit int) []*entity.Policy

	// Get permissions for offset and limit
	GetPermissions(offset int, limit int) []*entity.Permission

	// Get roles for offset and limit
	GetRoles(offset int, limit int) []*entity.Role

	// Get services for offset and limit
	GetServices(offset int, limit int) []*entity.Service
}

type ExtractorServiceImpl struct {
	PolicyRepository     driver.PolicyRepository
	PermissionRepository driver.PermissionRepository
	RoleRepository       driver.RoleRepository
	ServiceRepository    driver.ServiceRepository
}

func NewExtractorService() ExtractorService {
	return ExtractorServiceImpl{
		PolicyRepository:     driver.NewPolicyRepository(),
		PermissionRepository: driver.NewPermissionRepository(),
		RoleRepository:       driver.NewRoleRepository(),
		ServiceRepository:    driver.NewServiceRepository(),
	}
}

func (us ExtractorServiceImpl) GetPolicies(offset int, limit int) []*entity.Policy {
	policies, err := us.PolicyRepository.FindOffSetAndLimit(offset, limit)
	if err != nil {
		log.Logger.Error("Get policy query is failed", err.Detail)
		return []*entity.Policy{}
	}

	return policies
}

func (us ExtractorServiceImpl) GetPermissions(offset int, limit int) []*entity.Permission {
	permissions, err := us.PermissionRepository.FindOffSetAndLimit(offset, limit)
	if err != nil {
		log.Logger.Error("Get permission query is failed", err.Detail)
		return []*entity.Permission{}
	}

	return permissions
}

func (us ExtractorServiceImpl) GetRoles(offset int, limit int) []*entity.Role {
	roles, err := us.RoleRepository.FindOffSetAndLimit(offset, limit)
	if err != nil {
		log.Logger.Error("Get role query is failed", err.Detail)
		return []*entity.Role{}
	}

	return roles
}

func (us ExtractorServiceImpl) GetServices(offset int, limit int) []*entity.Service {
	services, err := us.ServiceRepository.FindOffSetAndLimit(offset, limit)
	if err != nil {
		log.Logger.Error("Get service query is failed", err.Detail)
		return []*entity.Service{}
	}

	return services
}
