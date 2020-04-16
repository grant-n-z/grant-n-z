package v1

import (
	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
)

// Less than stub struct
// TokenProcessor
type StubTokenProcessor struct {
}

func (tp StubTokenProcessor) Generate(userType string, groupIdStr string, userEntity entity.User) (string, *model.ErrorResBody) {
	return "", nil
}

func (tp StubTokenProcessor) ParseToken(token string) (map[string]string, bool) {
	resultMap := map[string]string{}
	return resultMap, true
}

func (tp StubTokenProcessor) VerifyOperatorToken(token string) (*model.AuthUser, *model.ErrorResBody) {
	return &model.AuthUser{}, nil
}

func (tp StubTokenProcessor) VerifyUserToken(token string, roleNames []string, permissionName string) (*model.AuthUser, *model.ErrorResBody) {
	return &model.AuthUser{}, nil
}

func (tp StubTokenProcessor) GetAuthUserInToken(token string) (*model.AuthUser, *model.ErrorResBody) {
	return &model.AuthUser{}, nil
}
