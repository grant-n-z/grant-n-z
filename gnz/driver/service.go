package driver

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"

	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
)

var srInstance ServiceRepository

type ServiceRepository interface {
	// Find all Service
	FindAll() ([]*entity.Service, *model.ErrorResBody)

	// Find Service for offset and limit
	FindOffSetAndLimit(offset int, limit int) ([]*entity.Service, *model.ErrorResBody)

	// Find Service by service id
	FindById(id int) (*entity.Service, *model.ErrorResBody)

	// Find Service by service name
	FindByName(name string) (*entity.Service, *model.ErrorResBody)

	// Find Service by service Api-Key
	FindByApiKey(apiKey string) (*entity.Service, *model.ErrorResBody)

	// Find Service name by service id
	FindNameById(id int) *string

	// Find Service name by Api-Key
	FindNameByApiKey(name string) *string

	// Fin Service by user_id
	FindServicesByUserId(userId int) ([]*entity.Service, *model.ErrorResBody)

	// Save Service
	Save(service entity.Service) (*entity.Service, *model.ErrorResBody)

	// Generate Service, ServicePermission, ServiceRole
	// When generate service, insert initialize permission and role data
	// Transaction mode
	SaveWithRelationalData(service entity.Service, roles []*entity.Role, permissions []*entity.Permission) (*entity.Service, *model.ErrorResBody)

	// Update Service
	Update(service entity.Service) *entity.Service
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

func (sri ServiceRepositoryImpl) FindAll() ([]*entity.Service, *model.ErrorResBody) {
	var services []*entity.Service
	if err := sri.Connection.Find(&services).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError(err.Error())
	}

	return services, nil
}

func (sri ServiceRepositoryImpl) FindOffSetAndLimit(offset int, limit int) ([]*entity.Service, *model.ErrorResBody) {
	var services []*entity.Service
	if err := sri.Connection.Limit(limit).Offset(offset).Find(&services).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError(err.Error())
	}

	return services, nil
}

func (sri ServiceRepositoryImpl) FindById(id int) (*entity.Service, *model.ErrorResBody) {
	var service entity.Service
	if err := sri.Connection.Where("id = ?", id).First(&service).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError(err.Error())
	}

	return &service, nil
}

func (sri ServiceRepositoryImpl) FindByName(name string) (*entity.Service, *model.ErrorResBody) {
	var service entity.Service
	if err := sri.Connection.Where("name = ?", name).First(&service).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError(err.Error())
	}

	return &service, nil
}

func (sri ServiceRepositoryImpl) FindByApiKey(apiKey string) (*entity.Service, *model.ErrorResBody) {
	var service entity.Service
	if err := sri.Connection.Where("api_key = ?", apiKey).First(&service).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError(err.Error())
	}

	return &service, nil
}

func (sri ServiceRepositoryImpl) FindNameById(id int) *string {
	service, err := sri.FindById(id)
	if err != nil {
		return nil
	}
	return &service.Name
}

func (sri ServiceRepositoryImpl) FindNameByApiKey(name string) *string {
	service, err := sri.FindByName(name)
	if err != nil {
		return nil
	}
	return &service.Name
}

func (sri ServiceRepositoryImpl) FindServicesByUserId(userId int) ([]*entity.Service, *model.ErrorResBody) {
	var services []*entity.Service

	if err := sri.Connection.Table(entity.ServiceTable.String()).
		Select("*").
		Joins(fmt.Sprintf("LEFT JOIN %s ON %s.%s = %s.%s",
			entity.UserServiceTable.String(),
			entity.ServiceTable.String(),
			entity.ServiceId,
			entity.UserServiceTable.String(),
			entity.UserServiceServiceId)).
		Where(fmt.Sprintf("%s.%s = ?",
			entity.UserServiceTable.String(),
			entity.UserServiceUserId), userId).
		Scan(&services).Error; err != nil {

		log.Logger.Warn(err.Error())
		if strings.Contains(err.Error(), "record not found") {
			return nil, model.NotFound("Not found service")
		}

		return nil, model.InternalServerError()
	}

	return services, nil
}

func (sri ServiceRepositoryImpl) Save(service entity.Service) (*entity.Service, *model.ErrorResBody) {
	if err := sri.Connection.Create(&service).Error; err != nil {
		log.Logger.Warn(err.Error())
		if strings.Contains(err.Error(), "1062") {
			return nil, model.Conflict("Already exit data.")
		}

		return nil, model.InternalServerError(err.Error())
	}

	return &service, nil
}

func (sri ServiceRepositoryImpl) SaveWithRelationalData(service entity.Service, roles []*entity.Role, permissions []*entity.Permission) (*entity.Service, *model.ErrorResBody) {
	tx := sri.Connection.Begin()

	// Save service
	if err := tx.Create(&service).Error; err != nil {
		log.Logger.Warn("Failed to save services at transaction process", err.Error())
		tx.Rollback()
		if strings.Contains(err.Error(), "1062") {
			return nil, model.Conflict("Already exit services data.")
		}

		return nil, model.InternalServerError()
	}

	// Save service_roles
	for _, role := range roles {
		serviceRole := entity.ServiceRole{
			RoleId:    role.Id,
			ServiceId: service.Id,
		}
		if err := tx.Create(&serviceRole).Error; err != nil {
			log.Logger.Warn("Failed to save service_roles at transaction process", err.Error())
			tx.Rollback()
			if strings.Contains(err.Error(), "1062") {
				return nil, model.Conflict("Already exit service_roles data.")
			}

			return nil, model.InternalServerError()
		}
	}

	// Save service_permissions
	for _, permission := range permissions {
		servicePermission := entity.ServicePermission{
			PermissionId: permission.Id,
			ServiceId:    service.Id,
		}
		if err := tx.Create(&servicePermission).Error; err != nil {
			log.Logger.Warn("Failed to save service_permissions at transaction process", err.Error())
			tx.Rollback()
			if strings.Contains(err.Error(), "1062") {
				return nil, model.Conflict("Already exit service_permissions data.")
			}

			return nil, model.InternalServerError()
		}
	}

	tx.Commit()

	return &service, nil
}

func (sri ServiceRepositoryImpl) Update(service entity.Service) *entity.Service {
	if err := sri.Connection.Update(&service).Error; err != nil {
		return nil
	}

	return &service
}
