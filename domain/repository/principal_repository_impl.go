package repository

import (
	"github.com/tomoyane/grant-n-z/domain/entity"
	"github.com/tomoyane/grant-n-z/infra"
)

type PrincipalRepositoryImpl struct {
}

// Find principal by principals.name
func (r PrincipalRepositoryImpl) FindByName(name string) *entity.Principal {
	principal := entity.Principal{}

	if err := infra.Db.Where("name = ?", name).First(&principal).Error; err != nil {
		if err.Error() == "record not found" {
			return &entity.Principal{}
		}
		return nil
	}

	return &principal
}

// Save to principal
func (r PrincipalRepositoryImpl) Save(principal entity.Principal) *entity.Principal {
	if err := infra.Db.Create(&principal).Error; err != nil {
		return nil
	}

	return &principal
}