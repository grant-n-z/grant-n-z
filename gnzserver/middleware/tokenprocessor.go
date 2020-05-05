package middleware

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"

	"github.com/tomoyane/grant-n-z/gnz/common"
	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
	"github.com/tomoyane/grant-n-z/gnzserver/service"
)

var tpInstance TokenProcessor

type TokenProcessor interface {
	// Generate jwt token
	Generate(userType string, groupIdStr string, tokenRequest model.TokenRequest) (*model.TokenResponse, *model.ErrorResBody)

	// Parse and check token
	ParseToken(token string) (map[string]string, bool)

	// Verify operator token
	VerifyOperatorToken(token string) (*model.AuthUser, *model.ErrorResBody)

	// Verify user token
	VerifyUserToken(token string, roleNames []string, permissionName string) (*model.AuthUser, *model.ErrorResBody)

	// Get auth user data in token
	// If invalid token, return 401
	GetAuthUserInToken(token string, isRefresh bool) (*model.AuthUser, *model.ErrorResBody)
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

func (tp TokenProcessorImpl) Generate(userType string, groupIdStr string, tokenRequest model.TokenRequest) (*model.TokenResponse, *model.ErrorResBody) {
	if strings.EqualFold(groupIdStr, "") {
		groupIdStr = "0"
	}

	groupId, err := strconv.Atoi(groupIdStr)
	if err != nil {
		return nil, model.BadRequest("Group id is only integer of query parameter")
	}

	if strings.EqualFold(tokenRequest.GrantType, model.GrantPassword.String()) {
		switch userType {
		case common.AuthOperator:
			return tp.generateOperatorToken(tokenRequest)
		case common.AuthUser:
			return tp.generateUserToken(tokenRequest, groupId)
		case "":
			return tp.generateUserToken(tokenRequest, groupId)
		default:
			return nil, model.BadRequest("Not support type of query parameter")
		}
	} else {
		return tp.generateTokenByRefreshToken(tokenRequest.RefreshToken)
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
	if _, ok := claims["username"].(string); !ok {
		return resultMap, false
	}

	resultMap["exp"] = claims["exp"].(string)
	resultMap["iat"] = claims["iat"].(string)
	resultMap["sub"] = claims["sub"].(string)
	resultMap["iss"] = claims["iss"].(string)
	resultMap["role_id"] = claims["role_id"].(string)
	resultMap["service_id"] = claims["service_id"].(string)
	resultMap["policy_id"] = claims["policy_id"].(string)
	resultMap["username"] = claims["username"].(string)

	return resultMap, true
}

func (tp TokenProcessorImpl) VerifyOperatorToken(token string) (*model.AuthUser, *model.ErrorResBody) {
	authUser, err := tp.GetAuthUserInToken(token, false)
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
	authUser, err := tp.GetAuthUserInToken(token, false)
	if err != nil {
		return nil, err
	}

	policy, err := tp.PolicyService.GetPolicyById(authUser.PolicyId)
	if err != nil {
		return nil, model.Forbidden("You don't join this group")
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

func (tp TokenProcessorImpl) GetAuthUserInToken(token string, isRefresh bool) (*model.AuthUser, *model.ErrorResBody) {
	if !isRefresh && !strings.Contains(token, "Bearer") {
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
	userName, _ := userData["username"]

	authUser := &model.AuthUser{
		UserId:    userId,
		UserUuid:  userUuid,
		UserName:  userName,
		ServiceId: serviceId,
		Expires:   userData["exp"],
		IssueDate: userData["iss"],
		RoleId:    roleId,
		PolicyId:  policyId,
	}

	return authUser, nil
}

func (tp TokenProcessorImpl) generateOperatorToken(tokenRequest model.TokenRequest) (*model.TokenResponse, *model.ErrorResBody) {
	targetUser, err := tp.UserService.GetUserWithOperatorPolicyByEmail(tokenRequest.Email)
	if err != nil || targetUser == nil {
		return nil, model.Unauthorized("Failed to email or password")
	}

	if !tp.UserService.ComparePw(targetUser.Password, tokenRequest.Password) {
		return nil, model.Unauthorized("Failed to email or password")
	}

	if targetUser.OperatorPolicy.RoleId != common.OperatorRoleId {
		return nil, model.Unauthorized("You don't have operator role")
	}

	// OperatorRole token is not required Service id, policy id
	serviceId := 0
	policyId := 0

	tokenExp := time.Now().Add(time.Hour * time.Duration(tp.ServerConfig.TokenExpireHour))
	token := tp.signedInToken(targetUser.UserId, targetUser.Uuid.String(), targetUser.Username, targetUser.OperatorPolicy.RoleId, serviceId, policyId, tokenExp)

	refreshTokenExp := time.Now().Add(time.Hour * time.Duration(tp.ServerConfig.TokenExpireHour) * 2)
	refreshToken := tp.signedInToken(targetUser.UserId, targetUser.Uuid.String(), targetUser.Username, targetUser.OperatorPolicy.RoleId, serviceId, policyId, refreshTokenExp)

	return &model.TokenResponse{Token: token, RefreshToken: refreshToken}, nil
}

func (tp TokenProcessorImpl) generateUserToken(tokenRequest model.TokenRequest, groupId int) (*model.TokenResponse, *model.ErrorResBody) {
	service, err := tp.Service.GetServiceOfSecret()
	if err != nil || service == nil {
		return nil, model.Unauthorized("Client-Secret is invalid")
	}

	targetUser, err := tp.UserService.GetUserByEmail(tokenRequest.Email)
	if err != nil || targetUser == nil {
		return nil, model.Unauthorized("Failed to email or password")
	}

	if !tp.UserService.ComparePw(targetUser.Password, tokenRequest.Password) {
		return nil, model.Unauthorized("Failed to email or password")
	}

	// Case of group_id is zero, not using policy.
	tokenExp := time.Now().Add(time.Hour * time.Duration(tp.ServerConfig.TokenExpireHour))
	if groupId == 0 {
		roleId := 0
		policyId := 0

		token := tp.signedInToken(targetUser.Id, targetUser.Uuid.String(), targetUser.Username, roleId, service.Id, policyId, tokenExp)
		refreshTokenExp := time.Now().Add(time.Hour * time.Duration(tp.ServerConfig.TokenExpireHour) * 2)
		refreshToken := tp.signedInToken(targetUser.Id, targetUser.Uuid.String(), targetUser.Username, roleId, service.Id, policyId, refreshTokenExp)
		return &model.TokenResponse{Token: token, RefreshToken: refreshToken}, nil
	}

	policy, err := tp.PolicyService.GetPolicyByUserGroup(targetUser.Id, groupId)
	if err != nil {
		return nil, model.Forbidden("Can't issue token for group")
	}

	token := tp.signedInToken(targetUser.Id, targetUser.Uuid.String(), targetUser.Username, 0, service.Id, policy.Id, tokenExp)
	refreshTokenExp := time.Now().Add(time.Hour * time.Duration(tp.ServerConfig.TokenExpireHour) * 2)
	refreshToken := tp.signedInToken(targetUser.Id, targetUser.Uuid.String(), targetUser.Username, 0, service.Id, policy.Id, refreshTokenExp)

	return &model.TokenResponse{Token: token, RefreshToken: refreshToken}, nil
}

func (tp TokenProcessorImpl) generateTokenByRefreshToken(refreshToken string) (*model.TokenResponse, *model.ErrorResBody) {
	authUser, err := tp.GetAuthUserInToken(refreshToken , true)
	if err != nil {
		return nil, err
	}

	tokenExp := time.Now().Add(time.Hour * time.Duration(tp.ServerConfig.TokenExpireHour))
	token := tp.signedInToken(authUser.UserId, authUser.UserUuid.String(), authUser.UserName, authUser.RoleId, authUser.ServiceId, authUser.PolicyId, tokenExp)

	refreshTokenExp := time.Now().Add(time.Hour * time.Duration(tp.ServerConfig.TokenExpireHour) * 2)
	newRefreshToken := tp.signedInToken(authUser.UserId, authUser.UserUuid.String(), authUser.UserName, authUser.RoleId, authUser.ServiceId, authUser.PolicyId, refreshTokenExp)
	return &model.TokenResponse{Token: token, RefreshToken: newRefreshToken}, nil
}

func (tp TokenProcessorImpl) signedInToken(userId int, userUuid string, userName string, roleId int, serviceId int, policyId int, exp time.Time) string {
	claims := tp.Token.Claims.(jwt.MapClaims)
	claims["exp"] = strconv.FormatInt(exp.Unix(), 10)
	claims["iat"] = strconv.FormatInt(time.Now().UnixNano(), 10)
	claims["sub"] = strconv.Itoa(userId)
	claims["iss"] = userUuid
	claims["role_id"] = strconv.Itoa(roleId)
	claims["service_id"] = strconv.Itoa(serviceId)
	claims["policy_id"] = strconv.Itoa(policyId)
	claims["username"] = userName

	signedToken, err := tp.Token.SignedString(tp.ServerConfig.SignedInPrivateKey)
	if err != nil {
		log.Logger.Error("Failed to issue signed token", err.Error())
		return ""
	}

	return signedToken
}
