package service

import (
	"strings"

	"github.com/tomoyane/grant-n-z/gnz/driver"
	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
)

var opsInstance OperatorPolicyService

type OperatorPolicyService interface {
	// Get
	Get(queryParam string) ([]*entity.OperatorPolicy, *model.ErrorResBody)

	// Get all operator policy
	GetAll() ([]*entity.OperatorPolicy, *model.ErrorResBody)

	// Get by user uuid
	GetByUserUuid(userUuid string) ([]*entity.OperatorPolicy, *model.ErrorResBody)

	// Get user uuid and role uuid
	GetByUserUuidAndRoleUuid(userUuid string, roleUuid string) (*entity.OperatorPolicy, *model.ErrorResBody)

	// Insert policy
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

	entities, err := ops.GetByUserUuid(queryParam)
	if err != nil {
		return nil, err
	}

	if entities == nil {
		return []*entity.OperatorPolicy{}, nil
	}

	return entities, nil
}

func (ops OperatorPolicyServiceImpl) GetAll() ([]*entity.OperatorPolicy, *model.ErrorResBody) {
	operatorPolicies, err := ops.OperatorPolicyRepository.FindAll()
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}
		return nil, model.InternalServerError(err.Error())
	}

	return operatorPolicies, nil
}

func (ops OperatorPolicyServiceImpl) GetByUserUuid(userUuid string) ([]*entity.OperatorPolicy, *model.ErrorResBody) {
	operatorPolicies, err := ops.OperatorPolicyRepository.FindByUserUuid(userUuid)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}
		return nil, model.InternalServerError(err.Error())
	}

	return operatorPolicies, nil
}

func (ops OperatorPolicyServiceImpl) GetByUserUuidAndRoleUuid(userUuid string, roleUuid string) (*entity.OperatorPolicy, *model.ErrorResBody) {
	operatorPolicy, err := ops.OperatorPolicyRepository.FindByUserUuidAndRoleUuid(userUuid, roleUuid)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}
		return nil, model.InternalServerError(err.Error())
	}

	return operatorPolicy, nil
}

func (ops OperatorPolicyServiceImpl) Insert(entity *entity.OperatorPolicy) (*entity.OperatorPolicy, *model.ErrorResBody) {
	if _, err := ops.UserRepository.FindByUuid(entity.UserUuid.String()); err != nil {
		if !strings.Contains(err.Error(), "record not found") {
			return nil, model.InternalServerError(err.Error())
		}
	}

	if _, err := ops.RoleRepository.FindByUuid(entity.RoleUuid.String()); err != nil {
		if !strings.Contains(err.Error(), "record not found") {
			return nil, model.InternalServerError(err.Error())
		}
	}

	savedOperatorPolicy, err := ops.OperatorPolicyRepository.Save(*entity)
	if err != nil {
		if strings.Contains(err.Error(), "1062") {
			return nil, model.Conflict("Already exit data.")
		} else if strings.Contains(err.Error(), "1452") {
			return nil, model.BadRequest("Not register relational id.")
		} else {
			return nil, model.InternalServerError(err.Error())
		}
	}

	return savedOperatorPolicy, nil
}
