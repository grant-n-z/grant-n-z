package service

import (
	"github.com/tomoyane/grant-n-z/gserver/model"
)

type AuthService interface {
	VerifyOperatorMember(token string) (*model.AuthUser, *model.ErrorResponse)

	VerifyServiceMember(token string) (*model.AuthUser, *model.ErrorResponse)

	verifyToken(token string) (*model.AuthUser, *model.ErrorResponse)
}
