package service

import (
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

type TokenService interface {
	Generate(queryParam string, userEntity entity.User) (*string, *model.ErrorResBody)

	operatorToken(userEntity entity.User) (*string, *model.ErrorResBody)

	serviceToken(userEntity entity.User) (*string, *model.ErrorResBody)

	userToken(userEntity entity.User) (*string, *model.ErrorResBody)
}
