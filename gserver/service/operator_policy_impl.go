package service

import (
	"strconv"
	"strings"

	"github.com/tomoyane/grant-n-z/gserver/common/driver"
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
	"github.com/tomoyane/grant-n-z/gserver/repository"
)

var opsInstance OperatorPolicyService

type operatorPolicyServiceImpl struct {
	operatorPolicyRepository repository.OperatorPolicyRepository
	userRepository           repository.UserRepository
	roleRepository           repository.RoleRepository
}

func GetOperatorPolicyServiceInstance() OperatorPolicyService {
	if opsInstance == nil {
		opsInstance = NewOperatorPolicyServiceService()
	}
	return opsInstance
}

func NewOperatorPolicyServiceService() OperatorPolicyService {
	log.Logger.Info("New `OperatorPolicyService` instance")
	log.Logger.Info("Inject `OperatorPolicyRepository`, `UserRepository`, `RoleRepository` to `OperatorPolicyService`")
	return operatorPolicyServiceImpl{
		operatorPolicyRepository: repository.GetOperatorPolicyRepositoryInstance(driver.Db),
		userRepository:           repository.GetUserRepositoryInstance(driver.Db),
		roleRepository:           repository.GetRoleRepositoryInstance(driver.Db),
	}
}

func (ops operatorPolicyServiceImpl) Get(queryParam string) ([]*entity.OperatorPolicy, *model.ErrorResBody) {
	if strings.EqualFold(queryParam, "") {
		return ops.GetAll()
	}

	i, castErr := strconv.Atoi(queryParam)
	if castErr != nil {
		log.Logger.Warn("The user_id is only integer")
		return nil, model.BadRequest(castErr.Error())
	}

	entities, err := ops.GetByUserId(i)
	if err != nil {
		return nil, err
	}

	if entities == nil {
		return []*entity.OperatorPolicy{}, nil
	}

	return entities, nil
}

func (ops operatorPolicyServiceImpl) GetAll() ([]*entity.OperatorPolicy, *model.ErrorResBody) {
	return ops.operatorPolicyRepository.FindAll()
}

func (ops operatorPolicyServiceImpl) GetByUserId(userId int) ([]*entity.OperatorPolicy, *model.ErrorResBody) {
	return ops.operatorPolicyRepository.FindByUserId(userId)
}

func (ops operatorPolicyServiceImpl) GetByUserIdAndRoleId(userId int, roleId int) (*entity.OperatorPolicy, *model.ErrorResBody) {
	return ops.operatorPolicyRepository.FindByUserIdAndRoleId(userId, roleId)
}

func (ops operatorPolicyServiceImpl) GetRoleNameByUserId(userId int) ([]string, *model.ErrorResBody) {
	return nil, nil
}

func (ops operatorPolicyServiceImpl) Insert(entity *entity.OperatorPolicy) (*entity.OperatorPolicy, *model.ErrorResBody) {
	if userEntity, _ := ops.userRepository.FindById(entity.UserId); userEntity == nil {
		log.Logger.Warn("Not found user id")
		return nil, model.BadRequest("Not found user id")
	}

	if roleEntity, _ := ops.roleRepository.FindById(entity.RoleId); roleEntity == nil {
		log.Logger.Warn("Not found role id")
		return nil, model.BadRequest("Not found role id")
	}

	return ops.operatorPolicyRepository.Save(*entity)
}
