package service

import (
	"github.com/tomoyane/grant-n-z/gnz/driver"
	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
)

var baseLimit = 500

type ExtractorService interface {
	// Get policies for offset and limit 500
	GetPolicy(offset int) []*entity.Policy

	GetPermission() []*entity.Permission

	GetRole() []*entity.Role

	GetService() []*entity.Service
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

func (us ExtractorServiceImpl) GetPolicy(offset int) []*entity.Policy {
	policies, err := us.PolicyRepository.FindOffSetAndLimit(offset, baseLimit)
	if err != nil {
		log.Logger.Error("Get policy query is failed", err.Detail)
		return []*entity.Policy{}
	}

	return policies
}

func (us ExtractorServiceImpl) GetPermission() []*entity.Permission {
	permissions, err := us.PermissionRepository.FindLimit(baseLimit)
	if err != nil {
		log.Logger.Error("Get permission query is failed", err.Detail)
		return []*entity.Permission{}
	}

	return permissions
}

func (us ExtractorServiceImpl) GetRole() []*entity.Role {
	roles, err := us.RoleRepository.FindLimit(baseLimit)
	if err != nil {
		log.Logger.Error("Get role query is failed", err.Detail)
		return []*entity.Role{}
	}

	return roles
}

func (us ExtractorServiceImpl) GetService() []*entity.Service {
	services, err := us.ServiceRepository.FindLimit(baseLimit)
	if err != nil {
		log.Logger.Error("Get service query is failed", err.Detail)
		return []*entity.Service{}
	}

	return services
}
