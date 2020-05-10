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

	// Get user_services for offset and limit
	GetUserServices(offset int, limit int) []*entity.UserService

	// Get user_groups for offset and limit
	GetUserGroups(offset int, limit int) []*entity.UserGroup
}

type ExtractorServiceImpl struct {
	PolicyRepository     driver.PolicyRepository
	PermissionRepository driver.PermissionRepository
	RoleRepository       driver.RoleRepository
	ServiceRepository    driver.ServiceRepository
	UserRepository       driver.UserRepository
}

func NewExtractorService() ExtractorService {
	return ExtractorServiceImpl{
		PolicyRepository:     driver.NewPolicyRepository(),
		PermissionRepository: driver.NewPermissionRepository(),
		RoleRepository:       driver.NewRoleRepository(),
		ServiceRepository:    driver.NewServiceRepository(),
		UserRepository:       driver.GetUserRepositoryInstance(),
	}
}

func (us ExtractorServiceImpl) GetPolicies(offset int, limit int) []*entity.Policy {
	policies, err := us.PolicyRepository.FindOffSetAndLimit(offset, limit)
	if err != nil {
		log.Logger.Error("Get policy query is failed", err.ToJson())
		return []*entity.Policy{}
	}

	return policies
}

func (us ExtractorServiceImpl) GetPermissions(offset int, limit int) []*entity.Permission {
	permissions, err := us.PermissionRepository.FindOffSetAndLimit(offset, limit)
	if err != nil {
		log.Logger.Error("Get permission query is failed", err.ToJson())
		return []*entity.Permission{}
	}

	return permissions
}

func (us ExtractorServiceImpl) GetRoles(offset int, limit int) []*entity.Role {
	roles, err := us.RoleRepository.FindOffSetAndLimit(offset, limit)
	if err != nil {
		log.Logger.Error("Get role query is failed", err.ToJson())
		return []*entity.Role{}
	}

	return roles
}

func (us ExtractorServiceImpl) GetServices(offset int, limit int) []*entity.Service {
	services, err := us.ServiceRepository.FindOffSetAndLimit(offset, limit)
	if err != nil {
		log.Logger.Error("Get service query is failed", err.ToJson())
		return []*entity.Service{}
	}

	return services
}

func (us ExtractorServiceImpl) GetUserServices(offset int, limit int) []*entity.UserService {
	userServices, err := us.UserRepository.FindUserServicesOffSetAndLimit(offset, limit)
	if err != nil {
		log.Logger.Error("Get user_service query is failed", err.ToJson())
		return []*entity.UserService{}
	}

	return userServices
}

func (us ExtractorServiceImpl) GetUserGroups(offset int, limit int) []*entity.UserGroup {
	userGroups, err := us.UserRepository.FindUserGroupsOffSetAndLimit(offset, limit)
	if err != nil {
		log.Logger.Error("Get user_group query is failed", err.ToJson())
		return []*entity.UserGroup{}
	}

	return userGroups
}
