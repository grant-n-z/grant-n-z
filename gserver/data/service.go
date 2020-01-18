package data

import (
	"strings"

	"github.com/jinzhu/gorm"

	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

var srInstance ServiceRepository

type ServiceRepository interface {
	// Find all service
	FindAll() ([]*entity.Service, *model.ErrorResBody)

	// Find service by service id
	FindById(id int) (*entity.Service, *model.ErrorResBody)

	// Find service by service name
	FindByName(name string) (*entity.Service, *model.ErrorResBody)

	// Find service by service Api-Key
	FindByApiKey(apiKey string) (*entity.Service, *model.ErrorResBody)

	// Find service name by service id
	FindNameById(id int) *string

	// Find service name by Api-Key
	FindNameByApiKey(name string) *string

	// Save service
	Save(service entity.Service) (*entity.Service, *model.ErrorResBody)

	// Generate service, service_permissions, service_roles
	// When generate service, insert initialize permission and role data
	// Transaction mode
	SaveWithRelationalData(service entity.Service, roleId int, permissionId int) (*entity.Service, *model.ErrorResBody)

	// Update service
	Update(service entity.Service) *entity.Service
}

// ServiceRepository struct
type ServiceRepositoryImpl struct {
	Db *gorm.DB
}

// Get Policy instance.
// If use singleton pattern, call this instance method
func GetServiceRepositoryInstance(db *gorm.DB) ServiceRepository {
	if srInstance == nil {
		srInstance = NewServiceRepository(db)
	}
	return srInstance
}

// Constructor
func NewServiceRepository(db *gorm.DB) ServiceRepository {
	log.Logger.Info("New `ServiceRepository` instance")
	return ServiceRepositoryImpl{
		Db: db,
	}
}

func (sri ServiceRepositoryImpl) FindAll() ([]*entity.Service, *model.ErrorResBody) {
	var services []*entity.Service
	if err := sri.Db.Find(&services).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError(err.Error())
	}

	return services, nil
}

func (sri ServiceRepositoryImpl) FindById(id int) (*entity.Service, *model.ErrorResBody) {
	var service entity.Service
	if err := sri.Db.Where("id = ?", id).First(&service).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError(err.Error())
	}

	return &service, nil
}

func (sri ServiceRepositoryImpl) FindByName(name string) (*entity.Service, *model.ErrorResBody) {
	var service entity.Service
	if err := sri.Db.Where("name = ?", name).First(&service).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError(err.Error())
	}

	return &service, nil
}

func (sri ServiceRepositoryImpl) FindByApiKey(apiKey string) (*entity.Service, *model.ErrorResBody) {
	var service entity.Service
	if err := sri.Db.Where("api_key = ?", apiKey).First(&service).Error; err != nil {
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

func (sri ServiceRepositoryImpl) Save(service entity.Service) (*entity.Service, *model.ErrorResBody) {
	if err := sri.Db.Create(&service).Error; err != nil {
		log.Logger.Warn(err.Error())
		if strings.Contains(err.Error(), "1062") {
			return nil, model.Conflict("Already exit data.")
		}

		return nil, model.InternalServerError(err.Error())
	}

	return &service, nil
}

func (sri ServiceRepositoryImpl) SaveWithRelationalData(service entity.Service, roleId int, permissionId int) (*entity.Service, *model.ErrorResBody) {
	tx := sri.Db.Begin()

	// Save service
	if err := tx.Create(&service).Error; err != nil {
		log.Logger.Warn("Failed to save services at transaction process", err.Error())
		tx.Rollback()
		if strings.Contains(err.Error(), "1062") {
			return nil, model.Conflict("Already exit services data.")
		}

		return nil, model.InternalServerError()
	}

	// Save service_permissions
	servicePermission := entity.ServicePermission{
		PermissionId: permissionId,
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

	// Save service_roles
	serviceRole := entity.ServiceRole{
		RoleId:    roleId,
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

	tx.Commit()

	return &service, nil
}

func (sri ServiceRepositoryImpl) Update(service entity.Service) *entity.Service {
	if err := sri.Db.Update(&service).Error; err != nil {
		return nil
	}

	return &service
}
