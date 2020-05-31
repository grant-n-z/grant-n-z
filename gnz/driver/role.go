package driver

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"

	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
)

var rrInstance RoleRepository

type RoleRepository interface {
	// Find all roles
	FindAll() ([]*entity.Role, error)

	// Find role for offset and limit
	FindOffSetAndLimit(offset int, limit int) ([]*entity.Role, error)

	// Find role by uuid
	FindByUuid(uuid string) (*entity.Role, error)

	// FInd role by role name
	FindByName(name string) (*entity.Role, error)

	// Find roles by role name array
	FindByNames(name []string) ([]entity.Role, error)

	// Find roles by group uuid
	// Join group_roles and roles
	FindByGroupUuid(groupUuid string) ([]*entity.Role, error)

	// Find role name by uuid
	FindNameByUuid(uuid string) *string

	// Save role
	Save(role entity.Role) (*entity.Role, error)

	// Save role with relational data
	SaveWithRelationalData(groupUuid string, role entity.Role) (*entity.Role, error)
}

type RoleRepositoryImpl struct {
	Connection *gorm.DB
}

func GetRoleRepositoryInstance() RoleRepository {
	if rrInstance == nil {
		rrInstance = NewRoleRepository()
	}
	return rrInstance
}

func NewRoleRepository() RoleRepository {
	log.Logger.Info("New `RoleRepository` instance")
	return RoleRepositoryImpl{Connection: connection}
}

func (rri RoleRepositoryImpl) FindAll() ([]*entity.Role, error) {
	var roles []*entity.Role
	if err := rri.Connection.Find(&roles).Error; err != nil {
		return nil, err
	}

	return roles, nil
}

func (rri RoleRepositoryImpl) FindOffSetAndLimit(offset int, limit int) ([]*entity.Role, error) {
	var roles []*entity.Role
	if err := rri.Connection.Limit(limit).Offset(offset).Find(&roles).Error; err != nil {
		return nil, err
	}

	return roles, nil
}

func (rri RoleRepositoryImpl) FindByUuid(uuid string) (*entity.Role, error) {
	var role entity.Role
	if err := rri.Connection.Where("uuid = ?", uuid).Find(&role).Error; err != nil {
		return nil, err
	}

	return &role, nil
}

func (rri RoleRepositoryImpl) FindByName(name string) (*entity.Role, error) {
	var role entity.Role
	if err := rri.Connection.Where("name = ?", name).Find(&role).Error; err != nil {
		return nil, err
	}

	return &role, nil
}

func (rri RoleRepositoryImpl) FindByNames(names []string) ([]entity.Role, error) {
	var roles []entity.Role
	if err := rri.Connection.Where("name IN (?)", names).Find(&roles).Error; err != nil {
		return nil, err
	}

	return roles, nil
}

func (rri RoleRepositoryImpl) FindByGroupUuid(groupUuid string) ([]*entity.Role, error) {
	var roles []*entity.Role

	if err := rri.Connection.Table(entity.GroupRoleTable.String()).
		Select("*").
		Joins(fmt.Sprintf("LEFT JOIN %s ON %s.%s = %s.%s",
			entity.RoleTable.String(),
			entity.GroupRoleTable.String(),
			entity.GroupRoleRoleUuid.String(),
			entity.RoleTable.String(),
			entity.RoleUuid.String())).
		Where(fmt.Sprintf("%s.%s = ?",
			entity.GroupRoleTable.String(),
			entity.GroupRoleGroupUuid.String()), groupUuid).
		Scan(&roles).Error; err != nil {

		return nil, err
	}

	return roles, nil
}

func (rri RoleRepositoryImpl) FindNameByUuid(uuid string) *string {
	role, err := rri.FindByUuid(uuid)
	if err != nil {
		return nil
	}
	return &role.Name
}

func (rri RoleRepositoryImpl) Save(role entity.Role) (*entity.Role, error) {
	if err := rri.Connection.Create(&role).Error; err != nil {
		return nil, err
	}

	return &role, nil
}

func (rri RoleRepositoryImpl) SaveWithRelationalData(gUuid string, role entity.Role) (*entity.Role, error) {
	tx := rri.Connection.Begin()

	// Save role
	if err := tx.Create(&role).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Save group_roles
	groupUuid, _ := uuid.FromBytes([]byte(gUuid))
	groupRole := entity.GroupRole{
		RoleUuid:  role.Uuid,
		GroupUuid: groupUuid,
	}
	if err := tx.Create(&groupRole).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return &role, nil
}
