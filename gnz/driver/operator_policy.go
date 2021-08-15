package driver

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
)

var oprInstance OperatorPolicyRepository

type OperatorPolicyRepository interface {
	FindAll() ([]*entity.OperatorPolicy, error)

	FindByUserUuid(userUuid string) ([]*entity.OperatorPolicy, error)

	FindByUserUuidAndRoleUuid(userUuid string, roleUuid string) (*entity.OperatorPolicy, error)

	FindRoleNameByUserUuid(userUuid string) ([]string, error)

	Save(role entity.OperatorPolicy) (*entity.OperatorPolicy, error)
}

type RdbmsOperatorPolicyRepository struct {
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
	return RdbmsOperatorPolicyRepository{Connection: connection}
}

func (opr RdbmsOperatorPolicyRepository) FindAll() ([]*entity.OperatorPolicy, error) {
	var entities []*entity.OperatorPolicy
	if err := opr.Connection.Find(&entities).Error; err != nil {
		return nil, err
	}

	return entities, nil
}

func (opr RdbmsOperatorPolicyRepository) FindByUserUuid(userUuid string) ([]*entity.OperatorPolicy, error) {
	var entities []*entity.OperatorPolicy
	if err := opr.Connection.Where("user_uuid = ?", userUuid).Find(&entities).Error; err != nil {
		return nil, err
	}

	return entities, nil
}

func (opr RdbmsOperatorPolicyRepository) FindByUserUuidAndRoleUuid(userUuid string, roleUuid string) (*entity.OperatorPolicy, error) {
	var operatorMemberRole entity.OperatorPolicy
	if err := opr.Connection.Where("user_uuid = ? AND role_uuid = ?", userUuid, roleUuid).Find(&operatorMemberRole).Error; err != nil {
		return nil, err
	}

	return &operatorMemberRole, nil
}

func (opr RdbmsOperatorPolicyRepository) FindRoleNameByUserUuid(userUuid string) ([]string, error) {
	query := opr.Connection.Table(entity.OperatorPolicyTable.String()).
		Select("name").
		Joins(fmt.Sprintf("LEFT JOIN %s ON %s.%s = %s.%s",
			entity.RoleTable.String(),
			entity.OperatorPolicyTable.String(),
			entity.OperatorPolicyRoleUuid.String(),
			entity.RoleTable.String(),
			entity.RoleUuid.String())).
		Where(fmt.Sprintf("%s.%s = ?",
			entity.OperatorPolicyTable.String(),
			entity.OperatorPolicyUserUuid.String()), userUuid)

	rows, err := query.Rows()
	if err != nil {
		return nil, err
	}

	var result struct {
		name *string
	}
	var names []string

	for rows.Next() {
		err := query.ScanRows(rows, &result)
		if err != nil {
			return nil, err
		}
		if result.name != nil {
			names = append(names, *result.name)
		}
	}

	return names, nil
}

func (opr RdbmsOperatorPolicyRepository) Save(entity entity.OperatorPolicy) (*entity.OperatorPolicy, error) {
	if err := opr.Connection.Create(&entity).Error; err != nil {
		return nil, err
	}

	return &entity, nil
}
