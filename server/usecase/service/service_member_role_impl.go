package service

import (
	"github.com/tomoyane/grant-n-z/server/config"
	"github.com/tomoyane/grant-n-z/server/entity"
	"github.com/tomoyane/grant-n-z/server/log"
	"github.com/tomoyane/grant-n-z/server/model"
	"github.com/tomoyane/grant-n-z/server/usecase/repository"
)

type serviceMemberRoleServiceImpl struct {
	serviceMemberRoleRepository repository.ServiceMemberRoleRepository
	roleRepository repository.RoleRepository
	userServiceRepository repository.UserServiceRepository
}

func NewServiceMemberRoleService() ServiceMemberRoleService {
	log.Logger.Info("Inject `serviceMemberRoleRepository` to `OperatorMemberRoleService`")
	return serviceMemberRoleServiceImpl{
		serviceMemberRoleRepository: repository.NewServiceMemberRoleRepository(config.Db),
		roleRepository: repository.NewRoleRepository(config.Db),
		userServiceRepository: repository.NewUserServiceRepository(config.Db),
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
