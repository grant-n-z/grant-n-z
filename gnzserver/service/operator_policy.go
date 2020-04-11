package service

import (
	"strconv"
	"strings"

	"github.com/tomoyane/grant-n-z/gnz/driver"
	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
)

var opsInstance OperatorPolicyService

type OperatorPolicyService interface {
	Get(queryParam string) ([]*entity.OperatorPolicy, *model.ErrorResBody)

	GetAll() ([]*entity.OperatorPolicy, *model.ErrorResBody)

	GetByUserId(userId int) ([]*entity.OperatorPolicy, *model.ErrorResBody)

	GetByUserIdAndRoleId(userId int, roleId int) (*entity.OperatorPolicy, *model.ErrorResBody)

	Insert(operatorPolicy *entity.OperatorPolicy) (*entity.OperatorPolicy, *model.ErrorResBody)
}

type OperatorPolicyServiceImpl struct {
	OperatorPolicyRepository driver.OperatorPolicyRepository
	UserRepository           driver.UserRepository
	RoleRepository           driver.RoleRepository
}

func GetOperatorPolicyServiceInstance() OperatorPolicyService {
	if opsInstance == nil {
		opsInstance = NewOperatorPolicyServiceService()
	}
	return opsInstance
}

func NewOperatorPolicyServiceService() OperatorPolicyService {
	log.Logger.Info("New `OperatorPolicyService` instance")
	return OperatorPolicyServiceImpl{
		OperatorPolicyRepository: driver.GetOperatorPolicyRepositoryInstance(),
		UserRepository:           driver.GetUserRepositoryInstance(),
		RoleRepository:           driver.GetRoleRepositoryInstance(),
	}
}

func (ops OperatorPolicyServiceImpl) Get(queryParam string) ([]*entity.OperatorPolicy, *model.ErrorResBody) {
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

func (ops OperatorPolicyServiceImpl) GetAll() ([]*entity.OperatorPolicy, *model.ErrorResBody) {
	return ops.OperatorPolicyRepository.FindAll()
}

func (ops OperatorPolicyServiceImpl) GetByUserId(userId int) ([]*entity.OperatorPolicy, *model.ErrorResBody) {
	return ops.OperatorPolicyRepository.FindByUserId(userId)
}

func (ops OperatorPolicyServiceImpl) GetByUserIdAndRoleId(userId int, roleId int) (*entity.OperatorPolicy, *model.ErrorResBody) {
	return ops.OperatorPolicyRepository.FindByUserIdAndRoleId(userId, roleId)
}

func (ops OperatorPolicyServiceImpl) Insert(entity *entity.OperatorPolicy) (*entity.OperatorPolicy, *model.ErrorResBody) {
	if userEntity, _ := ops.UserRepository.FindById(entity.UserId); userEntity == nil {
		log.Logger.Warn("Not found user id")
		return nil, model.BadRequest("Not found user id")
	}

	if roleEntity, _ := ops.RoleRepository.FindById(entity.RoleId); roleEntity == nil {
		log.Logger.Warn("Not found role id")
		return nil, model.BadRequest("Not found role id")
	}

	return ops.OperatorPolicyRepository.Save(*entity)
}
