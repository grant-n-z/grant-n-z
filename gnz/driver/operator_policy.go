package driver

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"

	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
)

var oprInstance OperatorPolicyRepository

type OperatorPolicyRepository interface {
	FindAll() ([]*entity.OperatorPolicy, *model.ErrorResBody)

	FindByUserId(userId int) ([]*entity.OperatorPolicy, *model.ErrorResBody)

	FindByUserIdAndRoleId(userId int, roleId int) (*entity.OperatorPolicy, *model.ErrorResBody)

	FindRoleNameByUserId(userId int) ([]string, *model.ErrorResBody)

	Save(role entity.OperatorPolicy) (*entity.OperatorPolicy, *model.ErrorResBody)
}

type OperatorPolicyRepositoryImpl struct {
	Connection *gorm.DB
}

func GetOperatorPolicyRepositoryInstance() OperatorPolicyRepository {
	if oprInstance == nil {
		oprInstance = NewOperatorPolicyRepository()
	}
	return oprInstance
}

func NewOperatorPolicyRepository() OperatorPolicyRepository {
	log.Logger.Info("New `OperatorPolicyRepository` instance")
	return OperatorPolicyRepositoryImpl{Connection: connection}
}

func (opr OperatorPolicyRepositoryImpl) FindAll() ([]*entity.OperatorPolicy, *model.ErrorResBody) {
	var entities []*entity.OperatorPolicy
	if err := opr.Connection.Find(&entities).Error; err != nil {
		log.Logger.Warn(err.Error())
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError(err.Error())
	}

	return entities, nil
}

func (opr OperatorPolicyRepositoryImpl) FindByUserId(userId int) ([]*entity.OperatorPolicy, *model.ErrorResBody) {
	var entities []*entity.OperatorPolicy
	if err := opr.Connection.Where("user_id = ?", userId).Find(&entities).Error; err != nil {
		log.Logger.Warn(err.Error())
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError(err.Error())
	}

	return entities, nil
}

func (opr OperatorPolicyRepositoryImpl) FindByUserIdAndRoleId(userId int, roleId int) (*entity.OperatorPolicy, *model.ErrorResBody) {
	var operatorMemberRole entity.OperatorPolicy
	if err := opr.Connection.Where("user_id = ? AND role_id = ?", userId, roleId).Find(&operatorMemberRole).Error; err != nil {
		log.Logger.Warn(err.Error())
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError(err.Error())
	}

	return &operatorMemberRole, nil
}

func (opr OperatorPolicyRepositoryImpl) FindRoleNameByUserId(userId int) ([]string, *model.ErrorResBody) {
	query := opr.Connection.Table(entity.OperatorPolicyTable.String()).
		Select("name").
		Joins(fmt.Sprintf("LEFT JOIN %s ON %s.%s = %s.%s",
			entity.RoleTable.String(),
			entity.OperatorPolicyTable.String(),
			entity.OperatorPolicyId,
			entity.RoleTable.String(),
			entity.RoleId)).
		Where(fmt.Sprintf("%s.%s = ?",
			entity.OperatorPolicyTable.String(),
			entity.OperatorPolicyUserId), userId)

	rows, err := query.Rows()
	if err != nil {
		log.Logger.Warn(err.Error())
		return nil, model.InternalServerError(err.Error())
	}

	var result struct {
		name *string
	}
	var names []string

	for rows.Next() {
		err := query.ScanRows(rows, &result)
		if err != nil {
			return nil, model.InternalServerError(err.Error())
		}
		if result.name != nil {
			names = append(names, *result.name)
		}
	}

	return names, nil
}

func (opr OperatorPolicyRepositoryImpl) Save(entity entity.OperatorPolicy) (*entity.OperatorPolicy, *model.ErrorResBody) {
	if err := opr.Connection.Create(&entity).Error; err != nil {
		log.Logger.Warn(err.Error())
		if strings.Contains(err.Error(), "1062") {
			return nil, model.Conflict("Already exit data.")
		} else if strings.Contains(err.Error(), "1452") {
			return nil, model.BadRequest("Not register relational id.")
		}

		return nil, model.InternalServerError()
	}

	return &entity, nil
}
