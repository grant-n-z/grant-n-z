package service

import (
	"strconv"
	"strings"

	"github.com/satori/go.uuid"

	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
	"github.com/tomoyane/grant-n-z/gserver/usecase/cache"
)

type AuthServiceImpl struct {
	UserService               UserService
	OperatorMemberRoleService OperatorMemberRoleService
	RedisClient               cache.RedisClient
}

func NewAuthService() AuthService {
	return AuthServiceImpl{
		UserService:               NewUserService(),
		OperatorMemberRoleService: NewOperatorMemberRoleService(),
		RedisClient:               cache.NewRedisClient(),
	}
}

func (as AuthServiceImpl) VerifyOperatorMember(token string) (*model.AuthUser, *model.ErrorResponse) {
	authUser, err := as.verifyToken(token)
	if err != nil {
		return nil, err
	}

	operatorRole, err := as.OperatorMemberRoleService.GetByUserIdAndRoleId(authUser.UserId, authUser.RoleId)
	if err != nil {
		return nil, err
	}

	if operatorRole == nil {
		log.Logger.Info("OperatorMemberRole data is null")
		return nil, model.Unauthorized("Failed to token.")
	}

	return authUser, nil
}

func (as AuthServiceImpl) VerifyServiceMember(token string) (*model.AuthUser, *model.ErrorResponse) {
	authUser, err := as.verifyToken(token)
	if err != nil {
		return nil, err
	}

	return authUser, nil
}

func (as AuthServiceImpl) verifyToken(token string) (*model.AuthUser, *model.ErrorResponse) {
	if !strings.Contains(token, "Bearer") {
		log.Logger.Info("Not contain `Bearer` authorization header")
		return nil, model.Unauthorized("Not contain `Bearer` authorization header.")
	}

	jwt := strings.Replace(token, "Bearer ", "", 1)
	userData, result := as.UserService.ParseJwt(jwt)
	if !result {
		log.Logger.Info("Failed to parse token")
		return nil, model.Unauthorized("Failed to token.")
	}

	// TODO: Read cache
	id, _ := strconv.Atoi(userData["user_id"])
	user, err := as.UserService.GetUserById(id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		log.Logger.Info("User data is null")
		return nil, model.Unauthorized("Failed to token.")
	}

	roleId, _ := strconv.Atoi(userData["role"])
	uid, _ := uuid.FromString(userData["user_uuid"])
	authUser := model.AuthUser{
		Username: userData["user_name"],
		UserUuid: uid,
		UserId:   id,
		Expires:  userData["expires"],
		RoleId:   roleId,
	}
	return &authUser, nil
}
