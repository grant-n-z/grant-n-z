package driver

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"

	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
)

var prInstance PermissionRepository

type PermissionRepository interface {
	// Find all permission
	FindAll() ([]*entity.Permission, error)

	// Find permission for offset and limit
	FindOffSetAndLimit(offsetCnt int, limitCnt int) ([]*entity.Permission, error)

	// Find permission by uuid
	FindByUuid(uuid string) (*entity.Permission, error)

	// Find permission by name
	FindByName(name string) (*entity.Permission, error)

	// Find permission by name array
	FindByNames(names []string) ([]entity.Permission, error)

	// Find permissions by group uuid
	// Join group_permission and permission
	FindByGroupUuid(groupUuid string) ([]*entity.Permission, error)

	// Find permission name by uuid
	FindNameByUuid(uuid string) *string

	// Save permission
	Save(permission entity.Permission) (*entity.Permission, error)

	// Save permission with relational data
	SaveWithRelationalData(groupUuid string, permission entity.Permission) (*entity.Permission, error)
}

type PermissionRepositoryImpl struct {
	Connection *gorm.DB
}

func GetPermissionRepositoryInstance() PermissionRepository {
	if prInstance == nil {
		prInstance = NewPermissionRepository()
	}
	return prInstance
}

func NewPermissionRepository() PermissionRepository {
	log.Logger.Info("New `PermissionRepository` instance")
	return PermissionRepositoryImpl{Connection: connection}
}

func (pri PermissionRepositoryImpl) FindAll() ([]*entity.Permission, error) {
	var permissions []*entity.Permission
	if err := pri.Connection.Find(&permissions).Error; err != nil {
		return nil, err
	}

	return permissions, nil
}

func (pri PermissionRepositoryImpl) FindOffSetAndLimit(offsetCnt int, limitCnt int) ([]*entity.Permission, error) {
	var permissions []*entity.Permission
	if err := pri.Connection.Limit(limitCnt).Offset(offsetCnt).Find(&permissions).Error; err != nil {
		return nil, err
	}

	return permissions, nil
}

func (pri PermissionRepositoryImpl) FindByUuid(uuid string) (*entity.Permission, error) {
	var permission entity.Permission
	if err := pri.Connection.Where("uuid = ?", uuid).Find(&permission).Error; err != nil {
		return nil, err
	}

	return &permission, nil
}

func (pri PermissionRepositoryImpl) FindByName(name string) (*entity.Permission, error) {
	var permission entity.Permission
	if err := pri.Connection.Where("name = ?", name).Find(&permission).Error; err != nil {
		return nil, err
	}

	return &permission, nil
}

func (pri PermissionRepositoryImpl) FindByNames(names []string) ([]entity.Permission, error) {
	var permissions []entity.Permission
	if err := pri.Connection.Where("name IN (?)", names).Find(&permissions).Error; err != nil {
		return nil, err
	}

	return permissions, nil
}

func (pri PermissionRepositoryImpl) FindByGroupUuid(groupUuid string) ([]*entity.Permission, error) {
	var permissions []*entity.Permission

	if err := pri.Connection.Table(entity.GroupPermissionTable.String()).
		Select("*").
		Joins(fmt.Sprintf("LEFT JOIN %s ON %s.%s = %s.%s",
			entity.PermissionTable.String(),
			entity.GroupPermissionTable.String(),
			entity.GroupPermissionPermissionUuid.String(),
			entity.PermissionTable.String(),
			entity.PermissionUuid.String())).
		Where(fmt.Sprintf("%s.%s = ?",
			entity.GroupPermissionTable.String(),
			entity.GroupPermissionGroupUuid.String()), groupUuid).
		Scan(&permissions).Error; err != nil {

			return nil, err
	}

	return permissions, nil
}

func (pri PermissionRepositoryImpl) FindNameByUuid(uuid string) *string {
	permission, err := pri.FindByUuid(uuid)
	if err != nil {
		return nil
	}
	return &permission.Name
}

func (pri PermissionRepositoryImpl) Save(permission entity.Permission) (*entity.Permission, error) {
	if err := pri.Connection.Create(&permission).Error; err != nil {
		return nil, err
	}

	return &permission, nil
}

func (pri PermissionRepositoryImpl) SaveWithRelationalData(gUuid string, permission entity.Permission) (*entity.Permission, error) {
	tx := pri.Connection.Begin()

	// Save permission
	if err := tx.Create(&permission).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Save group_permissions
	groupUuid, _ := uuid.FromBytes([]byte(gUuid))
	groupPermission := entity.GroupPermission{
		PermissionUuid: permission.Uuid,
		GroupUuid:      groupUuid,
	}
	if err := tx.Create(&groupPermission).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return &permission, nil
}
