package repository

import (
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

type UserGroupRepository interface {
	FindGroupsByUserId(userId int) ([]*entity.Group, *model.ErrorResBody)

	FindUsersByGroupId(groupId int) ([]*entity.User, *model.ErrorResBody)

	Save(userGroup entity.UserGroup) (*entity.UserGroup, *model.ErrorResBody)
}
