package driver

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"

	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
)

var srInstance ServiceRepository

type ServiceRepository interface {
	// FindAll
	// Find all Service
	FindAll() ([]*entity.Service, error)

	// FindOffSetAndLimit
	// Find Service for offset and limit
	FindOffSetAndLimit(offset int, limit int) ([]*entity.Service, error)

	// FindByUuid
	// Find Service by service uuid
	FindByUuid(uuid string) (*entity.Service, error)

	// FindByName
	// Find Service by service name
	FindByName(name string) (*entity.Service, error)

	// FindBySecret
	// Find Service by service Client-Secret
	FindBySecret(apiKey string) (*entity.Service, error)

	// FindNameByUuid
	// Find Service name by service uuid
	FindNameByUuid(uuid string) *string

	// FindServicesByUserUuid
	// Fin Service by user uuid
	FindServicesByUserUuid(userUuid string) ([]*entity.Service, error)

	// Save Service
	Save(service entity.Service) (*entity.Service, error)

	// SaveWithRelationalData
	// Generate Service, ServicePermission, ServiceRole
	// When generate service, insert initialize permission and role data
	// Transaction mode
	SaveWithRelationalData(service entity.Service, roles []entity.Role, permissions []entity.Permission) (*entity.Service, error)

	// Update Service
	Update(service entity.Service) (*entity.Service, error)
}

// RdbmsServiceRepository
// ServiceRepository struct
type RdbmsServiceRepository struct {
	Connection *gorm.DB
}

// GetServiceRepositoryInstance Get Policy instance.
// If use singleton pattern, call this instance method
func GetServiceRepositoryInstance() ServiceRepository {
	if srInstance == nil {
		srInstance = NewServiceRepository()
	}
	return srInstance
}

// NewServiceRepository
// Constructor
func NewServiceRepository() ServiceRepository {
	log.Logger.Info("New `ServiceRepository` instance")
	return RdbmsServiceRepository{Connection: connection}
}

func (sri RdbmsServiceRepository) FindAll() ([]*entity.Service, error) {
	var services []*entity.Service
	if err := sri.Connection.Find(&services).Error; err != nil {
		return nil, err
	}

	return services, nil
}

func (sri RdbmsServiceRepository) FindOffSetAndLimit(offset int, limit int) ([]*entity.Service, error) {
	var services []*entity.Service
	if err := sri.Connection.Limit(limit).Offset(offset).Find(&services).Error; err != nil {
		return nil, err
	}

	return services, nil
}

func (sri RdbmsServiceRepository) FindByUuid(uuid string) (*entity.Service, error) {
	var service entity.Service
	if err := sri.Connection.Where("uuid = ?", uuid).First(&service).Error; err != nil {
		return nil, err
	}

	return &service, nil
}

func (sri RdbmsServiceRepository) FindByName(name string) (*entity.Service, error) {
	var service entity.Service
	if err := sri.Connection.Where("name = ?", name).First(&service).Error; err != nil {
		return nil, err
	}

	return &service, nil
}

func (sri RdbmsServiceRepository) FindBySecret(secret string) (*entity.Service, error) {
	var service entity.Service
	if err := sri.Connection.Where("secret = ?", secret).First(&service).Error; err != nil {
		return nil, err
	}

	return &service, nil
}

func (sri RdbmsServiceRepository) FindNameByUuid(uuid string) *string {
	service, err := sri.FindByUuid(uuid)
	if err != nil {
		return nil
	}
	return &service.Name
}

func (sri RdbmsServiceRepository) FindServicesByUserUuid(userUuid string) ([]*entity.Service, error) {
	var services []*entity.Service

	if err := sri.Connection.Table(entity.ServiceTable.String()).
		Select("*").
		Joins(fmt.Sprintf("LEFT JOIN %s ON %s.%s = %s.%s",
			entity.UserServiceTable.String(),
			entity.ServiceTable.String(),
			entity.ServiceUuid.String(),
			entity.UserServiceTable.String(),
			entity.UserServiceServiceUuid.String())).
		Where(fmt.Sprintf("%s.%s = ?",
			entity.UserServiceTable.String(),
			entity.UserServiceUserUuid.String()), userUuid).
		Scan(&services).Error; err != nil {

		return nil, err
	}

	return services, nil
}

func (sri RdbmsServiceRepository) Save(service entity.Service) (*entity.Service, error) {
	if err := sri.Connection.Create(&service).Error; err != nil {
		return nil, err
	}

	return &service, nil
}

func (sri RdbmsServiceRepository) SaveWithRelationalData(service entity.Service, roles []entity.Role, permissions []entity.Permission) (*entity.Service, error) {
	tx := sri.Connection.Begin()

	// Save service
	if err := tx.Create(&service).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Save service_roles
	for _, role := range roles {
		serviceRoleMd5 := md5.Sum(uuid.New().NodeID())
		serviceRole := entity.ServiceRole{
			InternalId:  hex.EncodeToString(serviceRoleMd5[:]),
			RoleUuid:    role.Uuid,
			ServiceUuid: service.Uuid,
		}
		if err := tx.Create(&serviceRole).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	// Save service_permissions
	for _, permission := range permissions {
		servicePermissionMd5 := md5.Sum(uuid.New().NodeID())
		servicePermission := entity.ServicePermission{
			InternalId:     hex.EncodeToString(servicePermissionMd5[:]),
			PermissionUuid: permission.Uuid,
			ServiceUuid:    service.Uuid,
		}
		if err := tx.Create(&servicePermission).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	tx.Commit()
	return &service, nil
}

func (sri RdbmsServiceRepository) Update(service entity.Service) (*entity.Service, error) {
	if err := sri.Connection.Save(&service).Error; err != nil {
		return nil, err
	}

	return &service, nil
}
