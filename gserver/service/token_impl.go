package service

import (
	"strings"

	"github.com/tomoyane/grant-n-z/gserver/common/property"
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

var tsInstance TokenService

type tokenServiceImpl struct {
	userService               UserService
	operatorMemberRoleService OperatorPolicyService
}

func GetTokenServiceInstance() TokenService {
	if tsInstance == nil {
		tsInstance = NewTokenService()
	}
	return tsInstance
}

func NewTokenService() TokenService {
	log.Logger.Info("New `TokenService` instance")
	log.Logger.Info("Inject `UserGroup`, `OperatorPolicyService` to `TokenService`")
	return tokenServiceImpl{
		userService:               NewUserService(),
		operatorMemberRoleService: NewOperatorPolicyServiceService(),
	}
}

func (tsi tokenServiceImpl) Generate(queryParam string, userEntity entity.User) (*string, *model.ErrorResBody) {
	if strings.EqualFold(queryParam, property.AuthOperator) {
		return tsi.operatorToken(userEntity)
	} else if strings.EqualFold(queryParam, "") {
		return tsi.userToken(userEntity)
	} else {
		return nil, model.BadRequest("Not support type of query parameter")
	}
}

func (tsi tokenServiceImpl) operatorToken(userEntity entity.User) (*string, *model.ErrorResBody) {
	// TODO: Cache

	userWithRole, err := tsi.userService.GetUserWithRoleByEmail(userEntity.Email)
	if err != nil || userWithRole == nil {
		return nil, model.BadRequest("Failed to email or password")
	}
	if !tsi.userService.ComparePw(userWithRole.Password, userEntity.Password) {
		return nil, model.BadRequest("Failed to email or password")
	}
	if userWithRole.RoleId != property.OperatorRoleId {
		return nil, model.BadRequest("Can not issue token")
	}

	user := entity.User{
		Id:       userWithRole.UserId,
		Username: userWithRole.Username,
		Uuid:     userWithRole.Uuid,
	}
	return tsi.userService.GenerateJwt(&user, property.OperatorRoleId), nil
}

func (tsi tokenServiceImpl) serviceToken(userEntity entity.User) (*string, *model.ErrorResBody) {
	return nil, nil
}

func (tsi tokenServiceImpl) userToken(userEntity entity.User) (*string, *model.ErrorResBody) {
	// TODO: Cache
	// TODO: Set user policy

	userData, err := tsi.userService.GetUserByEmail(userEntity.Email)
	if err != nil || userData == nil {
		return nil, model.BadRequest("Failed to email or password")
	}

	if !tsi.userService.ComparePw(userData.Password, userEntity.Password) {
		return nil, model.BadRequest("Failed to email or password")
	}

	user := entity.User{
		Id:       userData.Id,
		Username: userData.Username,
		Uuid:     userData.Uuid,
	}
	return tsi.userService.GenerateJwt(&user, property.UserRoleId), nil
}
