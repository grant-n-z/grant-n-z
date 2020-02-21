package service

import (
	"github.com/tomoyane/grant-n-z/gnz/data"
	"github.com/tomoyane/grant-n-z/gnz/driver"
)

var baseLimit = 500

type ExtractorService interface {
	GetPolicy()

	GetPermission()

	GetRole()

	GetService()
}

type ExtractorServiceImpl struct {
	PolicyRepository     data.PolicyRepository
	PermissionRepository data.PermissionRepository
	RoleRepository       data.RoleRepository
	ServiceRepository    data.ServiceRepository
}

func NewExtractorService() ExtractorService {
	return ExtractorServiceImpl{
		PolicyRepository: data.NewPolicyRepository(driver.Rdbms),
		PermissionRepository: data.NewPermissionRepository(driver.Rdbms),
		RoleRepository: data.NewRoleRepository(driver.Rdbms),
		ServiceRepository: data.NewServiceRepository(driver.Rdbms),
	}
}

func (us ExtractorServiceImpl) GetPolicy() {
	us.PolicyRepository.FindLimit(baseLimit)
}

func (us ExtractorServiceImpl) GetPermission() {
	us.PermissionRepository.FindLimit(baseLimit)
}

func (us ExtractorServiceImpl) GetRole() {
	us.RoleRepository.FindLimit(baseLimit)
}

func (us ExtractorServiceImpl) GetService() {
	us.ServiceRepository.FindLimit(baseLimit)
}
