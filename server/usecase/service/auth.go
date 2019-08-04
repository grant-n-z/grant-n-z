package service

import "github.com/tomoyane/grant-n-z/server/model"

type AuthService interface {
	VerifyOperatorMember(token string) *model.ErrorResponse

	VerifyServiceMember(token string) *model.ErrorResponse
}
