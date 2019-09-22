package service

import (
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

type UserGroupService interface {
	InsertUserGroup(userGroup *entity.UserGroup) (*entity.UserGroup, *model.ErrorResBody)
}
