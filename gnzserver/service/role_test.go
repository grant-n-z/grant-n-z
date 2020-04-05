package service

import (
	"github.com/jinzhu/gorm"
	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
)

// Less than stub struct
// Role repository
type StubRoleRepositoryImpl struct {
	Connection *gorm.DB
}

func (rri StubRoleRepositoryImpl) FindAll() ([]*entity.Role, *model.ErrorResBody) {
	var roles []*entity.Role
	return roles, nil
}

func (rri StubRoleRepositoryImpl) FindOffSetAndLimit(offset int, limit int) ([]*entity.Role, *model.ErrorResBody) {
	var roles []*entity.Role
	return roles, nil
}

func (rri StubRoleRepositoryImpl) FindById(id int) (*entity.Role, *model.ErrorResBody) {
	var role entity.Role
	return &role, nil
}

func (rri StubRoleRepositoryImpl) FindByName(name string) (*entity.Role, *model.ErrorResBody) {
	var role entity.Role
	return &role, nil
}

func (rri StubRoleRepositoryImpl) FindByNames(names []string) ([]entity.Role, *model.ErrorResBody) {
	var roles []entity.Role
	return roles, nil
}

func (rri StubRoleRepositoryImpl) FindByGroupId(groupId int) ([]*entity.Role, *model.ErrorResBody) {
	var roles []*entity.Role
	return roles, nil
}

func (rri StubRoleRepositoryImpl) FindNameById(id int) *string {
	role, _ := rri.FindById(id)
	return &role.Name
}

func (rri StubRoleRepositoryImpl) Save(role entity.Role) (*entity.Role, *model.ErrorResBody) {
	return &role, nil
}

func (rri StubRoleRepositoryImpl) SaveWithRelationalData(groupId int, role entity.Role) (*entity.Role, *model.ErrorResBody) {
	return &role, nil
}

