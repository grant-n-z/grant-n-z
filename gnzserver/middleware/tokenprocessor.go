package middleware

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"encoding/json"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/tomoyane/grant-n-z/gnz/cache/structure"
	"github.com/tomoyane/grant-n-z/gnz/common"
	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
	"github.com/tomoyane/grant-n-z/gnzserver/service"
)

var tpInstance TokenProcessor

type TokenProcessor interface {
	// Generate jwt token
	Generate(userType string, tokenRequest model.TokenRequest) (*model.TokenResponse, *model.ErrorResBody)

	// Verify operator token
	VerifyOperatorToken(token string) (*model.JwtPayload, *model.ErrorResBody)

	// Verify user token
	VerifyUserToken(token string, roleNames string, permissionNames string, groupUuid string) (*model.JwtPayload, *model.ErrorResBody)

	// Get auth user data in token
	// If invalid token, return 401
	GetJwtPayload(token string, isRefresh bool) (*model.JwtPayload, *model.ErrorResBody)
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

func (tp TokenProcessorImpl) Generate(userType string, tokenRequest model.TokenRequest) (*model.TokenResponse, *model.ErrorResBody) {
	if strings.EqualFold(tokenRequest.GrantType, model.GrantPassword.String()) {
		switch userType {
		case common.AuthOperator:
			return tp.generateOperatorToken(tokenRequest)
		case common.AuthUser:
			return tp.generateUserToken(tokenRequest)
		default:
			return nil, model.BadRequest("Not support type of query parameter")
		}
	} else {
		return tp.generateTokenByRefreshToken(tokenRequest.RefreshToken)
	}
}

func (tp TokenProcessorImpl) VerifyOperatorToken(token string) (*model.JwtPayload, *model.ErrorResBody) {
	jwtPayload, err := tp.GetJwtPayload(token, false)
	if err != nil || jwtPayload.IsRefresh {
		return nil, err
	}

	isOperator := false
	for _, policy := range jwtPayload.UserPolicies {
		if policy.RoleName == common.OperatorRole {
			isOperator = true
		}
	}
	if !isOperator {
		return nil, model.Forbidden("Forbidden this token")
	}

	return jwtPayload, nil
}

func (tp TokenProcessorImpl) VerifyUserToken(token string, roleNames string, permissionNames string, groupUuid string) (*model.JwtPayload, *model.ErrorResBody) {
	jwtPayload, err := tp.GetJwtPayload(token, false)
	if err != nil || jwtPayload.IsRefresh {
		return nil, err
	}

	if groupUuid != "" {
		hasGroup := false
		userGroups := tp.UserService.GetUserGroupsByUserUuid(jwtPayload.UserUuid)
		for _, group := range userGroups {
			if strings.EqualFold(groupUuid, group.GroupUuid) {
				hasGroup = true
			}
		}
		if !hasGroup {
			return nil, model.Forbidden("Forbidden the user not join this group")
		}
	}

	hasRole := false
	hasPermission := false
	if roleNames != "" || permissionNames != "" {
		userPolicies := tp.UserService.GetUserPoliciesByUserUuid(jwtPayload.UserUuid)
		if userPolicies == nil {
			return nil, model.Forbidden("Forbidden the user has not policy")
		}
		for _, policy := range userPolicies {
			if policy.RoleName != "" && strings.Contains(roleNames, policy.RoleName) {
				hasRole = true
			}
			if policy.PermissionName != "" && strings.Contains(permissionNames, policy.PermissionName) {
				hasPermission = true
			}
		}
	}

	if roleNames != "" && !hasRole {
		return nil, model.Forbidden("Forbidden the user has not role")
	}
	if permissionNames != "" && !hasPermission {
		return nil, model.Forbidden("Forbidden the user has not permission")
	}

	return jwtPayload, nil
}

func (tp TokenProcessorImpl) GetJwtPayload(token string, isRefresh bool) (*model.JwtPayload, *model.ErrorResBody) {
	if !isRefresh && !strings.Contains(token, "Bearer") {
		log.Logger.Info("Not found authorization header or not contain `Bearer` in authorization header")
		return nil, model.Unauthorized("Unauthorized.")
	}

	payload, result := tp.parseToken(strings.Replace(token, "Bearer ", "", 1))
	if !result {
		return nil, model.Unauthorized("Token is invalid.")
	}

	err := tp.checkExpired(payload.Expires)
	if err != nil {
		return nil, err
	}

	return &payload, nil
}

func (tp TokenProcessorImpl) generateOperatorToken(tokenRequest model.TokenRequest) (*model.TokenResponse, *model.ErrorResBody) {
	targetUser, err := tp.UserService.GetUserWithOperatorPolicyByEmail(tokenRequest.Email)
	if err != nil || targetUser == nil {
		return nil, model.Unauthorized("Failed to email or password")
	}

	if !tp.UserService.ComparePw(targetUser.Password, tokenRequest.Password) {
		return nil, model.Unauthorized("Failed to email or password")
	}

	if targetUser.Role.Name != common.OperatorRole {
		return nil, model.Unauthorized("You don't have operator role")
	}

	// OperatorRole token is not required Service id, policy id
	tokenExp := time.Now().Add(time.Hour * time.Duration(tp.ServerConfig.TokenExpireHour))
	refreshTokenExp := time.Now().Add(time.Hour * time.Duration(tp.ServerConfig.TokenExpireHour) * 200)
	userPolicies := []structure.UserPolicy{{
		RoleName:       common.OperatorRole,
		PermissionName: common.AdminPermission,
	}}
	token := tp.generateTokenResponse(tokenExp, refreshTokenExp, userPolicies, targetUser.UserUuid.String(), targetUser.Username)
	return token, nil
}

func (tp TokenProcessorImpl) generateUserToken(tokenRequest model.TokenRequest) (*model.TokenResponse, *model.ErrorResBody) {
	user, err := tp.UserService.GetUserByEmail(tokenRequest.Email)
	if err != nil {
		return nil, model.Unauthorized("Failed to email or password")
	}

	if !tp.UserService.ComparePw(user.Password, tokenRequest.Password) {
		return nil, model.Unauthorized("Failed to email or password")
	}

	policies, err := tp.PolicyService.GetPoliciesByUser(user.Uuid.String())
	if err != nil {
		return nil, model.Forbidden("Can't issue token for group")
	}

	var userPolicies []structure.UserPolicy
	for _, policyRes := range policies {
		userPolicy := structure.UserPolicy{
			ServiceUuid:    policyRes.ServiceUuid.String(),
			GroupUuid:      policyRes.GroupUuid.String(),
			RoleName:       policyRes.RoleName,
			PermissionName: policyRes.PermissionName,
		}
		userPolicies = append(userPolicies, userPolicy)
	}

	exp := time.Now().Add(time.Hour * time.Duration(tp.ServerConfig.TokenExpireHour))
	rExp := time.Now().Add((time.Hour * time.Duration(tp.ServerConfig.TokenExpireHour)) * 200)
	return tp.generateTokenResponse(exp, rExp, userPolicies, user.Uuid.String(), user.Username), nil
}

func (tp TokenProcessorImpl) generateTokenByRefreshToken(refreshToken string) (*model.TokenResponse, *model.ErrorResBody) {
	jwtPayload, err := tp.GetJwtPayload(refreshToken, true)
	if err != nil {
		return nil, err
	}

	exp := time.Now().Add(time.Hour * time.Duration(tp.ServerConfig.TokenExpireHour))
	rExp := time.Now().Add(time.Hour * time.Duration(tp.ServerConfig.TokenExpireHour) * 200)

	policies := tp.UserService.GetUserPoliciesByUserUuid(jwtPayload.UserUuid)
	if policies == nil {
		policies = []structure.UserPolicy{}
	}

	return tp.generateTokenResponse(exp, rExp, policies, jwtPayload.UserUuid, jwtPayload.Username), nil
}

func (tp TokenProcessorImpl) generateTokenResponse(exp time.Time, rExp time.Time, userPolicy []structure.UserPolicy, userUuid string, username string) *model.TokenResponse {
	token := tp.signedInToken(userUuid, username, userPolicy, exp, false)
	refreshToken := tp.signedInToken(userUuid, username, userPolicy, rExp, true)
	return &model.TokenResponse{
		Token:        token,
		RefreshToken: refreshToken,
	}
}

func (tp TokenProcessorImpl) signedInToken(userUuid string, username string, userPolicies []structure.UserPolicy, exp time.Time, isRefresh bool) string {
	userPolicyJson, _ := json.Marshal(userPolicies)

	claims := tp.Token.Claims.(jwt.MapClaims)
	claims["exp"] = strconv.FormatInt(exp.Unix(), 10)
	claims["iat"] = strconv.FormatInt(time.Now().Unix(), 10)
	claims["sub"] = uuid.New().String() // TODO: grant nz server uuid
	claims["iss"] = userUuid
	claims["user_policies"] = string(userPolicyJson)
	claims["username"] = username
	if isRefresh {
		claims["is_refresh"] = true
	} else {
		claims["is_refresh"] = false
	}
	signedToken, err := tp.Token.SignedString(tp.ServerConfig.SignedInPrivateKey)
	if err != nil {
		log.Logger.Error("Failed to issue signed token", err.Error())
		return ""
	}

	return signedToken
}

func (tp TokenProcessorImpl) checkExpired(exp string) *model.ErrorResBody {
	i, _ := strconv.ParseInt(exp, 10, 64)
	expiredAt := time.Unix(i, 0).Unix()
	now := time.Now().Unix()
	if now > expiredAt {
		return model.Unauthorized("The access token provided has expired.")
	}
	return nil
}

func (tp TokenProcessorImpl) parseToken(token string) (model.JwtPayload, bool) {
	parseToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return "", errors.New("unexpected signing method")
		}
		return tp.ServerConfig.ValidatePublicKey, nil
	})

	if err != nil {
		e, ok := err.(*jwt.ValidationError)
		if !ok || ok && e.Errors&jwt.ValidationErrorIssuedAt == 0 {
			log.Logger.Info("Failed to token parse. ", err.Error())
			return model.JwtPayload{}, false
		}
	}

	//if !parseToken.Valid {
	//	log.Logger.Info("Failed to token validation.")
	//	return model.JwtPayload{}, false
	//}

	claims := parseToken.Claims.(jwt.MapClaims)

	var userPolicies []structure.UserPolicy
	err = json.Unmarshal([]byte(claims["user_policies"].(string)), &userPolicies)
	if err != nil {
		return model.JwtPayload{}, false
	}

	jwtPayload := model.JwtPayload{
		ServerId:     claims["sub"].(string),
		UserUuid:     claims["iss"].(string),
		Username:     claims["username"].(string),
		UserPolicies: userPolicies,
		Expires:      claims["exp"].(string),
		IssueDate:    claims["iat"].(string),
		IsRefresh:    claims["is_refresh"].(bool),
	}

	return jwtPayload, true
}
