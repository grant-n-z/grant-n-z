package service

import (
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

type TokenService interface {
	Generate(queryParam string, userEntity entity.User) (*string, *model.ErrorResponse)

	operatorToken(userEntity entity.User) (*string, *model.ErrorResponse)

	serviceToken(userEntity entity.User) (*string, *model.ErrorResponse)

	userToken(userEntity entity.User) (*string, *model.ErrorResponse)
}
