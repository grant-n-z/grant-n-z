package service

import (
	"strconv"
	"strings"

	"github.com/satori/go.uuid"

	"github.com/tomoyane/grant-n-z/gserver/cache"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

var asInstance AuthService

type AuthServiceImpl struct {
	userService               UserService
	operatorMemberRoleService OperatorMemberRoleService
	redisClient               cache.RedisClient
}

func GetAuthServiceInstance() AuthService {
	if asInstance == nil {
		asInstance = NewAuthService()
	}
	return asInstance
}

func NewAuthService() AuthService {
	log.Logger.Info("New `AuthService` instance")
	log.Logger.Info("Inject `UserService`, `OperatorMemberRoleService`, `RedisClient` to `AuthService`")
	return AuthServiceImpl{
		userService:               GetUserServiceInstance(),
		operatorMemberRoleService: GetOperatorMemberRoleServiceInstance(),
		redisClient:               cache.GetRedisClientInstance(),
	}
}

func (as AuthServiceImpl) VerifyOperatorMember(token string) (*model.AuthUser, *model.ErrorResponse) {
	authUser, err := as.verifyToken(token)
	if err != nil {
		return nil, err
	}

	operatorRole, err := as.operatorMemberRoleService.GetByUserIdAndRoleId(authUser.UserId, authUser.RoleId)
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
	userData, result := as.userService.ParseJwt(jwt)
	if !result {
		log.Logger.Info("Failed to parse token")
		return nil, model.Unauthorized("Failed to token.")
	}

	// TODO: Read cache
	id, _ := strconv.Atoi(userData["user_id"])
	user, err := as.userService.GetUserById(id)
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
