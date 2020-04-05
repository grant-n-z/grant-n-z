package service

import (
	"github.com/jinzhu/gorm"
	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
)

// Less than stub struct
// OperatorPolicy repository
type StubUserRepositoryImpl struct {
	Connection *gorm.DB
}

func (uri StubUserRepositoryImpl) FindById(id int) (*entity.User, *model.ErrorResBody) {
	var user entity.User
	return &user, nil
}

func (uri StubUserRepositoryImpl) FindByEmail(email string) (*entity.User, *model.ErrorResBody) {
	var user entity.User
	return &user, nil
}

func (uri StubUserRepositoryImpl) FindWithOperatorPolicyByEmail(email string) (*model.UserWithOperatorPolicy, *model.ErrorResBody) {
	var uwo model.UserWithOperatorPolicy
	return &uwo, nil
}

func (uri StubUserRepositoryImpl) FindWithUserServiceWithServiceByEmail(email string) (*model.UserWithUserServiceWithService, *model.ErrorResBody) {
	var uus model.UserWithUserServiceWithService
	return &uus, nil
}

func (uri StubUserRepositoryImpl) FindUserGroupByUserIdAndGroupId(userId int, groupId int) (*entity.UserGroup, *model.ErrorResBody) {
	var userGroup entity.UserGroup
	return &userGroup, nil
}

func (uri StubUserRepositoryImpl) FindUserServices() ([]*entity.UserService, *model.ErrorResBody) {
	var userServices []*entity.UserService
	return userServices, nil
}

func (uri StubUserRepositoryImpl) FindUserServicesOffSetAndLimit(offset int, limit int) ([]*entity.UserService, *model.ErrorResBody) {
	var userServices []*entity.UserService
	return userServices, nil
}

func (uri StubUserRepositoryImpl) FindUserServiceByUserIdAndServiceId(userId int, serviceId int) (*entity.UserService, *model.ErrorResBody) {
	var userService entity.UserService
	return &userService, nil
}

func (uri StubUserRepositoryImpl) SaveUserGroup(userGroup entity.UserGroup) (*entity.UserGroup, *model.ErrorResBody) {
	return &userGroup, nil
}

func (uri StubUserRepositoryImpl) SaveUser(user entity.User) (*entity.User, *model.ErrorResBody) {
	return &user, nil
}

func (uri StubUserRepositoryImpl) SaveWithUserService(user entity.User, userService entity.UserService) (*entity.User, *model.ErrorResBody) {
	return &user, nil
}

func (uri StubUserRepositoryImpl) UpdateUser(user entity.User) (*entity.User, *model.ErrorResBody) {
	return &user, nil
}

func (uri StubUserRepositoryImpl) SaveUserService(userService entity.UserService) (*entity.UserService, *model.ErrorResBody) {
	return &userService, nil
}
