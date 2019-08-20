package service

import (
	"github.com/tomoyane/grant-n-z/gserver/common/driver"
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
	"github.com/tomoyane/grant-n-z/gserver/repository"
)

type serviceMemberRoleServiceImpl struct {
	serviceMemberRoleRepository repository.ServiceMemberRoleRepository
	roleRepository repository.RoleRepository
	userServiceRepository repository.UserServiceRepository
}

func NewServiceMemberRoleService() ServiceMemberRoleService {
	log.Logger.Info("Inject `serviceMemberRoleRepository` to `OperatorMemberRoleService`")
	return serviceMemberRoleServiceImpl{
		serviceMemberRoleRepository: repository.NewServiceMemberRoleRepository(driver.Db),
		roleRepository: repository.NewRoleRepository(driver.Db),
		userServiceRepository: repository.NewUserServiceRepository(driver.Db),
	}
}

func (smrs serviceMemberRoleServiceImpl) GetAll() ([]*entity.ServiceMemberRole, *model.ErrorResponse) {
	return smrs.serviceMemberRoleRepository.FindAll()
}

func (smrs serviceMemberRoleServiceImpl) GetByRoleId(roleId int) ([]*entity.ServiceMemberRole, *model.ErrorResponse) {
	return smrs.serviceMemberRoleRepository.FindByRoleId(roleId)
}

func (smrs serviceMemberRoleServiceImpl) GetByUserServiceId(userServiceId int) ([]*entity.ServiceMemberRole, *model.ErrorResponse) {
	return smrs.serviceMemberRoleRepository.FindByUserServiceId(userServiceId)
}

func (smrs serviceMemberRoleServiceImpl) Insert(entity *entity.ServiceMemberRole) (*entity.ServiceMemberRole, *model.ErrorResponse) {
	if roleEntity, _ := smrs.roleRepository.FindById(entity.RoleId); roleEntity == nil {
		log.Logger.Warn("Not found role id")
		return nil, model.BadRequest("Not found role id")
	}

	if userServiceEntity, _ := smrs.userServiceRepository.FindById(entity.UserServiceId); userServiceEntity == nil {
		log.Logger.Warn("Not found role id")
		return nil, model.BadRequest("Not found role id")
	}

	return smrs.serviceMemberRoleRepository.Save(*entity)
}
