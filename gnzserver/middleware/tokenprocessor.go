package middleware

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"

	"github.com/tomoyane/grant-n-z/gnz/common"
	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
	"github.com/tomoyane/grant-n-z/gnzserver/service"
)

var tpInstance TokenProcessor

type TokenProcessor interface {
	// Generate jwt token
	Generate(userType string, groupIdStr string, userEntity entity.User) (string, *model.ErrorResBody)

	// Parse and check token
	ParseToken(token string) (map[string]string, bool)

	// Verify operator token
	VerifyOperatorToken(token string) (*model.AuthUser, *model.ErrorResBody)

	// Verify user token
	VerifyUserToken(token string, roleNames []string, permissionName string) (*model.AuthUser, *model.ErrorResBody)

	// Get auth user data in token
	// If invalid token, return 401
	GetAuthUserInToken(token string) (*model.AuthUser, *model.ErrorResBody)
}

// TokenProcessor struct
type TokenProcessorImpl struct {
	UserService           service.UserService
	OperatorPolicyService service.OperatorPolicyService
	Service               service.Service
	PolicyService         service.PolicyService
	RoleService           service.RoleService
	PermissionService     service.PermissionService
	ServerConfig          common.ServerConfig
	Token                 *jwt.Token
}

// Get TokenProcessor instance.
// If use singleton pattern, call this instance method
func GetTokenProcessorInstance() TokenProcessor {
	if tpInstance == nil {
		tpInstance = NewTokenProcessor()
	}
	return tpInstance
}

// Constructor
func NewTokenProcessor() TokenProcessor {
	log.Logger.Info("New `TokenProcessor` instance")
	serverConfig := common.GServer
	return TokenProcessorImpl{
		UserService:           service.GetUserServiceInstance(),
		OperatorPolicyService: service.GetOperatorPolicyServiceInstance(),
		Service:               service.GetServiceInstance(),
		PolicyService:         service.GetPolicyServiceInstance(),
		RoleService:           service.GetRoleServiceInstance(),
		PermissionService:     service.GetPermissionServiceInstance(),
		ServerConfig:          serverConfig,
		Token:                 jwt.New(serverConfig.SigningMethod),
	}
}

func (tp TokenProcessorImpl) Generate(userType string, groupIdStr string, userEntity entity.User) (string, *model.ErrorResBody) {
	if strings.EqualFold(groupIdStr, "") {
		groupIdStr = "0"
	}

	groupId, err := strconv.Atoi(groupIdStr)
	if err != nil {
		return "", model.BadRequest("Group id is only integer of query parameter")
	}

	switch userType {
	case common.AuthOperator:
		return tp.generateOperatorToken(userEntity)
	case common.AuthUser:
		return tp.generateUserToken(userEntity, groupId)
	case "":
		return tp.generateUserToken(userEntity, groupId)
	default:
		return "", model.BadRequest("Not support type of query parameter")
	}
}

func (tp TokenProcessorImpl) ParseToken(token string) (map[string]string, bool) {
	resultMap := map[string]string{}

	parseToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return "", errors.New("unexpected signing method")
		}
		return tp.ServerConfig.ValidatePublicKey, nil
	})

	if err != nil || !parseToken.Valid {
		log.Logger.Info("Failed to token validation.", err.Error())
		return resultMap, false
	}

	claims := parseToken.Claims.(jwt.MapClaims)
	if _, ok := claims["iss"].(string); !ok {
		return resultMap, false
	}
	if _, ok := claims["sub"].(string); !ok {
		return resultMap, false
	}
	if _, ok := claims["exp"].(string); !ok {
		return resultMap, false
	}
	if _, ok := claims["iat"].(string); !ok {
		return resultMap, false
	}
	if _, ok := claims["role_id"].(string); !ok {
		return resultMap, false
	}
	if _, ok := claims["service_id"].(string); !ok {
		return resultMap, false
	}
	if _, ok := claims["policy_id"].(string); !ok {
		return resultMap, false
	}

	resultMap["exp"] = claims["exp"].(string)
	resultMap["iat"] = claims["iat"].(string)
	resultMap["sub"] = claims["sub"].(string)
	resultMap["iss"] = claims["iss"].(string)
	resultMap["role_id"] = claims["role_id"].(string)
	resultMap["service_id"] = claims["service_id"].(string)
	resultMap["policy_id"] = claims["policy_id"].(string)

	return resultMap, true
}

func (tp TokenProcessorImpl) VerifyOperatorToken(token string) (*model.AuthUser, *model.ErrorResBody) {
	authUser, err := tp.GetAuthUserInToken(token)
	if err != nil {
		return nil, err
	}

	operatorRole, err := tp.OperatorPolicyService.GetByUserIdAndRoleId(authUser.UserId, authUser.RoleId)
	if operatorRole == nil || err != nil {
		log.Logger.Info("Not contain operator role or failed to query")
		return nil, model.Forbidden("Forbidden this token")
	}

	return authUser, nil
}

func (tp TokenProcessorImpl) VerifyUserToken(token string, roleNames []string, permissionName string) (*model.AuthUser, *model.ErrorResBody) {
	authUser, err := tp.GetAuthUserInToken(token)
	if err != nil {
		return nil, err
	}

	policy, err := tp.PolicyService.GetPolicyById(authUser.PolicyId)
	if err != nil {
		return nil, model.BadRequest("You don't join this group")
	}

	if len(roleNames) > 0 && !strings.EqualFold(roleNames[0], "") {
		roles, err := tp.RoleService.GetRoleByNames(roleNames)
		if err != nil {
			return nil, model.Forbidden("Forbidden the user has not role")
		}
		result := false
		for _, role := range roles {
			if role.Id == policy.RoleId {
				result = true
				break
			}
		}
		if !result {
			return nil, model.Forbidden("Forbidden the user policy does not match role")
		}
	}

	if !strings.EqualFold(permissionName, "") {
		permission, err := tp.PermissionService.GetPermissionByName(permissionName)
		if permission == nil || err != nil || permission.Id != policy.PermissionId {
			return nil, model.Forbidden("Forbidden the user has not permission")
		}
	}

	userService, err := tp.UserService.GetUserServiceByUserIdAndServiceId(authUser.UserId, authUser.ServiceId)
	if userService == nil || err != nil {
		return nil, model.Forbidden("Forbidden the user cannot access Service")
	}

	return authUser, nil
}

func (tp TokenProcessorImpl) GetAuthUserInToken(token string) (*model.AuthUser, *model.ErrorResBody) {
	if !strings.Contains(token, "Bearer") {
		log.Logger.Info("Not found authorization header or not contain `Bearer` in authorization header")
		return nil, model.Unauthorized("Unauthorized.")
	}

	userData, result := tp.ParseToken(strings.Replace(token, "Bearer ", "", 1))
	if !result {
		return nil, model.Unauthorized("Token is invalid.")
	}

	userId, _ := strconv.Atoi(userData["sub"])
	userUuid, _ := uuid.Parse(userData["iss"])
	roleId, _ := strconv.Atoi(userData["role_id"])
	serviceId, _ := strconv.Atoi(userData["service_id"])
	policyId, _ := strconv.Atoi(userData["policy_id"])

	authUser := &model.AuthUser{
		UserId:    userId,
		UserUuid:  userUuid,
		ServiceId: serviceId,
		Expires:   userData["exp"],
		IssueDate: userData["iss"],
		RoleId:    roleId,
		PolicyId:  policyId,
	}

	return authUser, nil
}

func (tp TokenProcessorImpl) generateOperatorToken(userEntity entity.User) (string, *model.ErrorResBody) {
	targetUser, err := tp.UserService.GetUserWithOperatorPolicyByEmail(userEntity.Email)
	if err != nil || targetUser == nil {
		return "", model.BadRequest("Failed to email or password")
	}

	if !tp.UserService.ComparePw(targetUser.Password, userEntity.Password) {
		return "", model.BadRequest("Failed to email or password")
	}

	if targetUser.OperatorPolicy.RoleId != common.OperatorRoleId {
		return "", model.BadRequest("Can not issue token")
	}

	// OperatorRole token is not required Service id, policy id
	serviceId := 0
	policyId := 0
	return tp.signedInToken(targetUser.UserId, targetUser.Uuid.String(), targetUser.OperatorPolicy.RoleId, serviceId, policyId), nil
}

func (tp TokenProcessorImpl) generateUserToken(userEntity entity.User, groupId int) (string, *model.ErrorResBody) {
	service, err := tp.Service.GetServiceOfApiKey()
	if err != nil || service == nil {
		return "", model.BadRequest("Not found registered services by Api-Key")
	}

	targetUser, err := tp.UserService.GetUserByEmail(userEntity.Email)
	if err != nil || targetUser == nil {
		return "", model.BadRequest("Failed to email or password")
	}

	if !tp.UserService.ComparePw(targetUser.Password, userEntity.Password) {
		return "", model.BadRequest("Failed to email or password")
	}

	// Case of group_id is zero, not using policy.
	if groupId == 0 {
		roleId := 0
		policyId := 0
		return tp.signedInToken(targetUser.Id, targetUser.Uuid.String(), roleId, service.Id, policyId), nil
	}

	policy, err := tp.PolicyService.GetPolicyByUserGroup(targetUser.Id, groupId)
	if err != nil {
		return "", model.Forbidden("Can't issue token for group")
	}

	return tp.signedInToken(targetUser.Id, targetUser.Uuid.String(), 0, service.Id, policy.Id), nil
}

func (tp TokenProcessorImpl) signedInToken(userId int, userUuid string, roleId int, serviceId int, policyId int) string {
	claims := tp.Token.Claims.(jwt.MapClaims)
	claims["exp"] = strconv.FormatInt(time.Now().Add(time.Hour * time.Duration(tp.ServerConfig.TokenExpireHour)).Unix(), 10)
	claims["iat"] = strconv.FormatInt(time.Now().UnixNano(), 10)
	claims["sub"] = strconv.Itoa(userId)
	claims["iss"] = userUuid
	claims["role_id"] = strconv.Itoa(roleId)
	claims["service_id"] = strconv.Itoa(serviceId)
	claims["policy_id"] = strconv.Itoa(policyId)

	signedToken, err := tp.Token.SignedString(tp.ServerConfig.SignedInPrivateKey)
	if err != nil {
		log.Logger.Error("Failed to issue signed token", err.Error())
		return ""
	}

	return signedToken
}
