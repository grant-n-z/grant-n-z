package service

import (
	"strings"

	"github.com/tomoyane/grant-n-z/gserver/common/property"
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

type tokenServiceImpl struct {
	userService               UserService
	operatorMemberRoleService OperatorMemberRoleService
}

func NewTokenService() TokenService {
	log.Logger.Info("Inject `UserService`, `OperatorMemberRoleService` to `TokenService`")
	return tokenServiceImpl{
		userService:               NewUserService(),
		operatorMemberRoleService: NewOperatorMemberRoleService(),
	}
}

func (tsi tokenServiceImpl) Generate(queryParam string, userEntity entity.User) (*string, *model.ErrorResponse) {
	if strings.EqualFold(queryParam, property.AuthOperator) {
		return tsi.operatorToken(userEntity)
	} else if strings.EqualFold(queryParam, "") {
		return tsi.userToken(userEntity)
	} else {
		return nil, model.BadRequest("Not support type of query parameter")
	}
}

func (tsi tokenServiceImpl) operatorToken(userEntity entity.User) (*string, *model.ErrorResponse) {
	user, err := tsi.userService.GetUserByEmail(userEntity.Email)
	if err != nil || user == nil {
		return nil, model.BadRequest("Failed to email or password")
	}

	if !tsi.userService.ComparePw(user.Password, userEntity.Password) {
		return nil, model.BadRequest("Failed to email or password")
	}

	memberRoles, err := tsi.operatorMemberRoleService.GetByUserId(user.Id)
	if err != nil {
		return nil, err
	}

	result := false
	for _, v := range memberRoles {
		if v.RoleId == property.OperatorRoleId {
			result = true
		}
	}

	if !result {
		return nil, model.BadRequest("Can not issue token this request")
	}

	return tsi.userService.GenerateJwt(user, property.OperatorRoleId), nil
}

func (tsi tokenServiceImpl) serviceToken(userEntity entity.User) (*string, *model.ErrorResponse) {
	return nil, nil
}

func (tsi tokenServiceImpl) userToken(userEntity entity.User) (*string, *model.ErrorResponse) {
	return nil, nil
}
