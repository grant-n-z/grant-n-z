package repository

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"

	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

type OperatorMemberRoleRepositoryImpl struct {
	Db *gorm.DB
}

func NewOperatorMemberRoleRepository(db *gorm.DB) OperatorMemberRoleRepository {
	return OperatorMemberRoleRepositoryImpl{
		Db: db,
	}
}

func (omrri OperatorMemberRoleRepositoryImpl) FindAll() ([]*entity.OperatorMemberRole, *model.ErrorResponse) {
	var entities []*entity.OperatorMemberRole
	if err := omrri.Db.Find(&entities).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError(err.Error())
	}

	return entities, nil
}

func (omrri OperatorMemberRoleRepositoryImpl) FindByUserId(userId int) ([]*entity.OperatorMemberRole, *model.ErrorResponse) {
	var entities []*entity.OperatorMemberRole
	if err := omrri.Db.Where("user_id = ?", userId).Find(&entities).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError(err.Error())
	}

	return entities, nil
}

func (omrri OperatorMemberRoleRepositoryImpl) FindRoleNameByUserId(userId int) ([]string, *model.ErrorResponse) {
	query := omrri.Db.Table(entity.OperatorMemberRole{}.TableName()).
		Select("name").
		Joins(fmt.Sprintf("LEFT JOIN %s ON %s.%s = %s.%s",
			entity.Role{}.TableName(),
			entity.OperatorMemberRole{}.TableName(),
			entity.OperatorMemberRoleRoleId,
			entity.Role{}.TableName(),
			entity.RoleId)).
		Where(fmt.Sprintf("%s.%s = ?",
			entity.OperatorMemberRole{}.TableName(),
			entity.OperatorMemberRoleUserId), userId)

	rows, err := query.Rows()
	if err != nil {
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

func (omrri OperatorMemberRoleRepositoryImpl) Save(entity entity.OperatorMemberRole) (*entity.OperatorMemberRole, *model.ErrorResponse) {
	if err := omrri.Db.Create(&entity).Error; err != nil {
		log.Logger.Warn(err.Error())
		if strings.Contains(err.Error(), "1062") {
			return nil, model.Conflict("Already exit data.")
		} else if strings.Contains(err.Error(), "1452") {
			return nil, model.BadRequest("Not register relational id.")
		}

		return nil, model.InternalServerError("Error internal processing.")
	}

	return &entity, nil
}
