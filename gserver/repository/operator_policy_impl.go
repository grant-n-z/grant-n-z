package repository

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"

	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

var oprInstance OperatorPolicyRepository

type OperatorPolicyRepositoryImpl struct {
	Db *gorm.DB
}

func GetOperatorPolicyRepositoryInstance(db *gorm.DB) OperatorPolicyRepository {
	if oprInstance == nil {
		oprInstance = NewOperatorPolicyRepository(db)
	}
	return oprInstance
}

func NewOperatorPolicyRepository(db *gorm.DB) OperatorPolicyRepository {
	log.Logger.Info("New `OperatorPolicyRepository` instance")
	log.Logger.Info("Inject `gorm.DB` to `OperatorPolicyRepository`")
	return OperatorPolicyRepositoryImpl{
		Db: db,
	}
}

func (opr OperatorPolicyRepositoryImpl) FindAll() ([]*entity.OperatorPolicy, *model.ErrorResponse) {
	var entities []*entity.OperatorPolicy
	if err := opr.Db.Find(&entities).Error; err != nil {
		log.Logger.Warn(err.Error())
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError(err.Error())
	}

	return entities, nil
}

func (opr OperatorPolicyRepositoryImpl) FindByUserId(userId int) ([]*entity.OperatorPolicy, *model.ErrorResponse) {
	var entities []*entity.OperatorPolicy
	if err := opr.Db.Where("user_id = ?", userId).Find(&entities).Error; err != nil {
		log.Logger.Warn(err.Error())
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError(err.Error())
	}

	return entities, nil
}

func (opr OperatorPolicyRepositoryImpl) FindByUserIdAndRoleId(userId int, roleId int) (*entity.OperatorPolicy, *model.ErrorResponse) {
	var operatorMemberRole entity.OperatorPolicy
	if err := opr.Db.Where("user_id = ? AND role_id = ?", userId, roleId).Find(&operatorMemberRole).Error; err != nil {
		log.Logger.Warn(err.Error())
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError(err.Error())
	}

	return &operatorMemberRole, nil
}

func (opr OperatorPolicyRepositoryImpl) FindRoleNameByUserId(userId int) ([]string, *model.ErrorResponse) {
	query := opr.Db.Table(entity.OperatorPolicy{}.TableName()).
		Select("name").
		Joins(fmt.Sprintf("LEFT JOIN %s ON %s.%s = %s.%s",
			entity.Role{}.TableName(),
			entity.OperatorPolicy{}.TableName(),
			entity.OperatorMemberRoleRoleId,
			entity.Role{}.TableName(),
			entity.RoleId)).
		Where(fmt.Sprintf("%s.%s = ?",
			entity.OperatorPolicy{}.TableName(),
			entity.OperatorMemberRoleUserId), userId)

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

func (opr OperatorPolicyRepositoryImpl) Save(entity entity.OperatorPolicy) (*entity.OperatorPolicy, *model.ErrorResponse) {
	if err := opr.Db.Create(&entity).Error; err != nil {
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
