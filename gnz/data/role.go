package data

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"

	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnzserver/entity"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
)

var rrInstance RoleRepository

type RoleRepository interface {
	// Find all roles
	FindAll() ([]*entity.Role, *model.ErrorResBody)

	// Find role by id
	FindById(id int) (*entity.Role, *model.ErrorResBody)

	// FInd role by role name
	FindByName(name string) (*entity.Role, *model.ErrorResBody)

	// Find roles by role name array
	FindByNames(name []string) ([]*entity.Role, *model.ErrorResBody)

	// Find roles by group id
	// Join group_roles and roles
	FindByGroupId(groupId int) ([]*entity.Role, *model.ErrorResBody)

	// Find role name by id
	FindNameById(id int) *string

	// Save role
	Save(role entity.Role) (*entity.Role, *model.ErrorResBody)

	// Save role with relational data
	SaveWithRelationalData(groupId int, role entity.Role) (*entity.Role, *model.ErrorResBody)
}

type RoleRepositoryImpl struct {
	Db *gorm.DB
}

func GetRoleRepositoryInstance(db *gorm.DB) RoleRepository {
	if rrInstance == nil {
		rrInstance = NewRoleRepository(db)
	}
	return rrInstance
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	log.Logger.Info("New `RoleRepository` instance")
	return RoleRepositoryImpl{Db: db}
}

func (rri RoleRepositoryImpl) FindAll() ([]*entity.Role, *model.ErrorResBody) {
	var roles []*entity.Role
	if err := rri.Db.Find(&roles).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError(err.Error())
	}

	return roles, nil
}

func (rri RoleRepositoryImpl) FindById(id int) (*entity.Role, *model.ErrorResBody) {
	var role entity.Role
	if err := rri.Db.Where("id = ?", id).Find(&role).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError(err.Error())
	}

	return &role, nil
}

func (rri RoleRepositoryImpl) FindByName(name string) (*entity.Role, *model.ErrorResBody) {
	var role entity.Role
	if err := rri.Db.Where("name = ?", name).Find(&role).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError(err.Error())
	}

	return &role, nil
}

func (rri RoleRepositoryImpl) FindByNames(names []string) ([]*entity.Role, *model.ErrorResBody) {
	var roles []*entity.Role
	if err := rri.Db.Where("name IN (?)", names).Find(&roles).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError(err.Error())
	}

	return roles, nil
}

func (rri RoleRepositoryImpl) FindByGroupId(groupId int) ([]*entity.Role, *model.ErrorResBody) {
	var roles []*entity.Role

	if err := rri.Db.Table(entity.GroupRoleTable.String()).
		Select("*").
		Joins(fmt.Sprintf("LEFT JOIN %s ON %s.%s = %s.%s",
			entity.RoleTable.String(),
			entity.GroupRoleTable.String(),
			entity.GroupRoleRoleId,
			entity.RoleTable.String(),
			entity.RoleId)).
		Where(fmt.Sprintf("%s.%s = ?",
			entity.GroupRoleTable.String(),
			entity.GroupRoleGroupId), groupId).
		Scan(&roles).Error; err != nil {

		log.Logger.Warn(err.Error())
		if strings.Contains(err.Error(), "record not found") {
			return nil, model.NotFound("Not found role")
		}

		return nil, model.InternalServerError()
	}

	return roles, nil
}

func (rri RoleRepositoryImpl) FindNameById(id int) *string {
	if id == 0 {
		return nil
	}
	role, err := rri.FindById(id)
	if err != nil {
		return nil
	}
	return &role.Name
}

func (rri RoleRepositoryImpl) Save(role entity.Role) (*entity.Role, *model.ErrorResBody) {
	if err := rri.Db.Create(&role).Error; err != nil {
		log.Logger.Warn(err.Error())
		if strings.Contains(err.Error(), "1062") {
			return nil, model.Conflict("Already exit data.")
		}

		return nil, model.InternalServerError(err.Error())
	}

	return &role, nil
}

func (rri RoleRepositoryImpl) SaveWithRelationalData(groupId int, role entity.Role) (*entity.Role, *model.ErrorResBody) {
	tx := rri.Db.Begin()

	// Save role
	if err := tx.Create(&role).Error; err != nil {
		log.Logger.Warn("Failed to save roles at transaction process", err.Error())
		tx.Rollback()
		if strings.Contains(err.Error(), "1062") {
			return nil, model.Conflict("Already exit roles data.")
		}

		return nil, model.InternalServerError()
	}

	// Save group_roles
	groupRole := entity.GroupRole{
		RoleId:  role.Id,
		GroupId: groupId,
	}
	if err := tx.Create(&groupRole).Error; err != nil {
		log.Logger.Warn("Failed to save group_roles at transaction process", err.Error())
		tx.Rollback()
		if strings.Contains(err.Error(), "1062") {
			return nil, model.Conflict("Already exit group_roles data.")
		}

		return nil, model.InternalServerError()
	}

	tx.Commit()

	return &role, nil
}
