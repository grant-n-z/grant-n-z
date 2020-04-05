package service

import (
	"github.com/jinzhu/gorm"
	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
)

// Less than stub struct
// Permission repository
type StubPermissionRepositoryImpl struct {
	Connection *gorm.DB
}

func (pri StubPermissionRepositoryImpl) FindAll() ([]*entity.Permission, *model.ErrorResBody) {
	var permissions []*entity.Permission
	return permissions, nil
}

func (pri StubPermissionRepositoryImpl) FindOffSetAndLimit(offsetCnt int, limitCnt int) ([]*entity.Permission, *model.ErrorResBody) {
	var permissions []*entity.Permission
	return permissions, nil
}

func (pri StubPermissionRepositoryImpl) FindById(id int) (*entity.Permission, *model.ErrorResBody) {
	var permission entity.Permission
	return &permission, nil
}

func (pri StubPermissionRepositoryImpl) FindByName(name string) (*entity.Permission, *model.ErrorResBody) {
	var permission entity.Permission
	return &permission, nil
}

func (pri StubPermissionRepositoryImpl) FindByNames(names []string) ([]entity.Permission, *model.ErrorResBody) {
	var permissions []entity.Permission
	return permissions, nil
}

func (pri StubPermissionRepositoryImpl) FindByGroupId(groupId int) ([]*entity.Permission, *model.ErrorResBody) {
	var permissions []*entity.Permission
	return permissions, nil
}

func (pri StubPermissionRepositoryImpl) FindNameById(id int) *string {
	permission, _ := pri.FindById(id)
	return &permission.Name
}

func (pri StubPermissionRepositoryImpl) Save(permission entity.Permission) (*entity.Permission, *model.ErrorResBody) {
	return &permission, nil
}

func (pri StubPermissionRepositoryImpl) SaveWithRelationalData(groupId int, permission entity.Permission) (*entity.Permission, *model.ErrorResBody) {
	return &permission, nil
}
