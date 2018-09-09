package repository

import (
	"github.com/tomoyane/grant-n-z/domain/entity"
	"github.com/tomoyane/grant-n-z/infra"
)

type RoleRepositoryImpl struct {
}

// Find role by roles.user_uuid
func (r RoleRepositoryImpl) FindByUserUuid(userUuid string) *entity.Role {
	role := entity.Role{}

	if err := infra.Db.Where("user_uuid = ?", userUuid).First(&role).Error; err != nil {
		if err.Error() == "record not found" {
			return &entity.Role{}
		}
		return nil
	}

	return &role
}

// Save to role
func (r RoleRepositoryImpl) Save(role entity.Role) *entity.Role {
	if err := infra.Db.Create(&role).Error; err != nil {
		return nil
	}

	return &role
}