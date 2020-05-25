package service

import (
	"github.com/tomoyane/grant-n-z/gnz/cache/structure"
	"github.com/tomoyane/grant-n-z/gnz/driver"
)

type ExtractorService interface {
	// Get policies for offset and limit
	GetPolicies(offset int, limit int) map[string][]structure.UserPolicy

	// Get permissions for offset and limit
	GetPermissions(offset int, limit int) []structure.Permission

	// Get roles for offset and limit
	GetRoles(offset int, limit int) []structure.Role

	// Get services for offset and limit
	GetServices(offset int, limit int) []structure.Service

	// Get user_services for offset and limit
	GetUserServices(offset int, limit int) map[string][]structure.UserService

	// Get user_groups for offset and limit
	GetUserGroups(offset int, limit int) map[string][]structure.UserGroup
}

type ExtractorServiceImpl struct {
	PolicyRepository     driver.PolicyRepository
	PermissionRepository driver.PermissionRepository
	RoleRepository       driver.RoleRepository
	ServiceRepository    driver.ServiceRepository
	UserRepository       driver.UserRepository
	GroupRepository      driver.GroupRepository
}

func NewExtractorService() ExtractorService {
	return ExtractorServiceImpl{
		PolicyRepository:     driver.NewPolicyRepository(),
		PermissionRepository: driver.NewPermissionRepository(),
		RoleRepository:       driver.NewRoleRepository(),
		ServiceRepository:    driver.NewServiceRepository(),
		UserRepository:       driver.GetUserRepositoryInstance(),
		GroupRepository:      driver.GetGroupRepositoryInstance(),
	}
}

func (es ExtractorServiceImpl) GetPolicies(offset int, limit int) map[string][]structure.UserPolicy {
	userServices, err := es.UserRepository.FindUserServicesOffSetAndLimit(offset, limit)
	if err != nil {
		return nil
	}

	userPolicyMap := make(map[string][]structure.UserPolicy)
	checkedUserUuid := ""
	var userPolicies []structure.UserPolicy
	for _, userService := range userServices {
		if checkedUserUuid == userService.UserUuid.String() {
			continue
		}

		policies, err := es.PolicyRepository.FindPolicyOfUserServiceByUserUuidAndServiceUuid(userService.UserUuid.String())
		if err != nil {
			return nil
		}

		for _, policy := range policies {
			userPolicies = append(userPolicies, structure.UserPolicy{
				ServiceUuid:    userService.ServiceUuid.String(),
				GroupUuid:      policy.GroupUuid,
				RoleName:       policy.RoleName,
				PermissionName: policy.PermissionName,
			})
		}

		userPolicyMap[userService.UserUuid.String()] = userPolicies
		checkedUserUuid = userService.UserUuid.String()
	}

	return userPolicyMap
}

func (es ExtractorServiceImpl) GetPermissions(offset int, limit int) []structure.Permission {
	permissions, err := es.PermissionRepository.FindOffSetAndLimit(offset, limit)
	if err != nil {
		return []structure.Permission{}
	}

	var stPermissions []structure.Permission
	for _, permission := range permissions {
		stPermissions = append(stPermissions, structure.Permission{
			Name: permission.Name,
			Uuid: permission.Uuid.String(),
		})
	}

	return stPermissions
}

func (es ExtractorServiceImpl) GetRoles(offset int, limit int) []structure.Role {
	roles, err := es.RoleRepository.FindOffSetAndLimit(offset, limit)
	if err != nil {
		return []structure.Role{}
	}

	var stRoles []structure.Role
	for _, role := range roles {
		stRoles = append(stRoles, structure.Role{
			Name: role.Name,
			Uuid: role.Uuid.String(),
		})
	}

	return stRoles
}

func (es ExtractorServiceImpl) GetServices(offset int, limit int) []structure.Service {
	services, err := es.ServiceRepository.FindOffSetAndLimit(offset, limit)
	if err != nil {
		return []structure.Service{}
	}

	var stServices []structure.Service
	for _, ser := range services {
		stServices = append(stServices, structure.Service{
			Name: ser.Name,
			Uuid: ser.Uuid.String(),
		})
	}

	return stServices
}

func (es ExtractorServiceImpl) GetUserServices(offset int, limit int) map[string][]structure.UserService {
	userServices, err := es.UserRepository.FindUserServicesOffSetAndLimit(offset, limit)
	if err != nil {
		return map[string][]structure.UserService{}
	}

	userServiceMap := make(map[string][]structure.UserService)
	checkedUserUuid := ""
	var stUserServices []structure.UserService
	for _, userService := range userServices {
		ser, err := es.ServiceRepository.FindByUuid(userService.ServiceUuid.String())
		if err != nil {
			return map[string][]structure.UserService{}
		}

		if checkedUserUuid != userService.UserUuid.String() {
			stUserServices = stUserServices[:0]
		}

		stUserServices = append(stUserServices, structure.UserService{
			ServiceUUid: userService.ServiceUuid.String(),
			ServiceName: ser.Name,
		})

		userServiceMap[userService.UserUuid.String()] = stUserServices
		checkedUserUuid = userService.UserUuid.String()
	}

	return userServiceMap
}

func (es ExtractorServiceImpl) GetUserGroups(offset int, limit int) map[string][]structure.UserGroup {
	userGroups, err := es.UserRepository.FindUserGroupsOffSetAndLimit(offset, limit)
	if err != nil {
		return map[string][]structure.UserGroup{}
	}

	userGroupMap := make(map[string][]structure.UserGroup)
	checkedUserUuid := ""
	var stUserGroups []structure.UserGroup
	for _, userGroup := range userGroups {
		group, err := es.GroupRepository.FindByUuid(userGroup.GroupUuid.String())
		if err != nil {
			return map[string][]structure.UserGroup{}
		}


		if checkedUserUuid != userGroup.UserUuid.String() {
			stUserGroups = stUserGroups[:0]
		}

		stUserGroups = append(stUserGroups, structure.UserGroup{
			GroupUuid: userGroup.GroupUuid.String(),
			GroupName: group.Name,
		})

		userGroupMap[userGroup.UserUuid.String()] = stUserGroups
		checkedUserUuid = userGroup.UserUuid.String()
	}

	return userGroupMap
}
