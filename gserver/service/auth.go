package service

import (
	"github.com/tomoyane/grant-n-z/gserver/model"
)

type AuthService interface {
	VerifyOperatorMember(token string) (*model.AuthUser, *model.ErrorResBody)

	VerifyServiceMember(token string) (*model.AuthUser, *model.ErrorResBody)

	verifyToken(token string) (*model.AuthUser, *model.ErrorResBody)
}
