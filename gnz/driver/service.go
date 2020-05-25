package driver

import (
	"fmt"
	"github.com/jinzhu/gorm"

	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
)

var srInstance ServiceRepository

type ServiceRepository interface {
	// Find all Service
	FindAll() ([]*entity.Service, error)

	// Find Service for offset and limit
	FindOffSetAndLimit(offset int, limit int) ([]*entity.Service, error)

	// Find Service by service uuid
	FindByUuid(uuid string) (*entity.Service, error)

	// Find Service by service name
	FindByName(name string) (*entity.Service, error)

	// Find Service by service Client-Secret
	FindBySecret(apiKey string) (*entity.Service, error)

	// Find Service name by service uuid
	FindNameByUuid(uuid string) *string

	// Fin Service by user uuid
	FindServicesByUserUuid(userUuid string) ([]*entity.Service, error)

	// Save Service
	Save(service entity.Service) (*entity.Service, error)

	// Generate Service, ServicePermission, ServiceRole
	// When generate service, insert initialize permission and role data
	// Transaction mode
	SaveWithRelationalData(service entity.Service, roles []entity.Role, permissions []entity.Permission) (*entity.Service, error)

	// Update Service
	Update(service entity.Service) (*entity.Service, error)
}

// ServiceRepository struct
type ServiceRepositoryImpl struct {
	Connection *gorm.DB
}

// Get Policy instance.
// If use singleton pattern, call this instance method
func GetServiceRepositoryInstance() ServiceRepository {
	if srInstance == nil {
		srInstance = NewServiceRepository()
	}
	return srInstance
}

// Constructor
func NewServiceRepository() ServiceRepository {
	log.Logger.Info("New `ServiceRepository` instance")
	return ServiceRepositoryImpl{Connection: connection}
}

func (sri ServiceRepositoryImpl) FindAll() ([]*entity.Service, error) {
	var services []*entity.Service
	if err := sri.Connection.Find(&services).Error; err != nil {
		return nil, err
	}

	return services, nil
}

func (sri ServiceRepositoryImpl) FindOffSetAndLimit(offset int, limit int) ([]*entity.Service, error) {
	var services []*entity.Service
	if err := sri.Connection.Limit(limit).Offset(offset).Find(&services).Error; err != nil {
		return nil, err
	}

	return services, nil
}

func (sri ServiceRepositoryImpl) FindByUuid(uuid string) (*entity.Service, error) {
	var service entity.Service
	if err := sri.Connection.Where("uuid = ?", uuid).First(&service).Error; err != nil {
		return nil, err
	}

	return &service, nil
}

func (sri ServiceRepositoryImpl) FindByName(name string) (*entity.Service, error) {
	var service entity.Service
	if err := sri.Connection.Where("name = ?", name).First(&service).Error; err != nil {
		return nil, err
	}

	return &service, nil
}

func (sri ServiceRepositoryImpl) FindBySecret(secret string) (*entity.Service, error) {
	var service entity.Service
	if err := sri.Connection.Where("secret = ?", secret).First(&service).Error; err != nil {
		return nil, err
	}

	return &service, nil
}

func (sri ServiceRepositoryImpl) FindNameByUuid(uuid string) *string {
	service, err := sri.FindByUuid(uuid)
	if err != nil {
		return nil
	}
	return &service.Name
}

func (sri ServiceRepositoryImpl) FindServicesByUserUuid(userUuid string) ([]*entity.Service, error) {
	var services []*entity.Service

	if err := sri.Connection.Table(entity.ServiceTable.String()).
		Select("*").
		Joins(fmt.Sprintf("LEFT JOIN %s ON %s.%s = %s.%s",
			entity.UserServiceTable.String(),
			entity.ServiceTable.String(),
			entity.ServiceUuid.String(),
			entity.UserServiceTable.String(),
			entity.UserServiceServiceUuid.String())).
		Where(fmt.Sprintf(	"%s.%s = ?",
			entity.UserServiceTable.String(),
			entity.UserServiceUserUuid.String()), userUuid).
		Scan(&services).Error; err != nil {

			return nil, err
	}

	return services, nil
}

func (sri ServiceRepositoryImpl) Save(service entity.Service) (*entity.Service, error) {
	if err := sri.Connection.Create(&service).Error; err != nil {
		return nil, err
	}

	return &service, nil
}

func (sri ServiceRepositoryImpl) SaveWithRelationalData(service entity.Service, roles []entity.Role, permissions []entity.Permission) (*entity.Service, error) {
	tx := sri.Connection.Begin()

	// Save service
	if err := tx.Create(&service).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Save service_roles
	for _, role := range roles {
		serviceRole := entity.ServiceRole{
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
		servicePermission := entity.ServicePermission{
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

func (sri ServiceRepositoryImpl) Update(service entity.Service) (*entity.Service, error) {
	if err := sri.Connection.Save(&service).Error; err != nil {
		return nil, err
	}

	return &service, nil
}
