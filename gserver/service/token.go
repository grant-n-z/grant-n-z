package service

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/tomoyane/grant-n-z/gserver/common/config"
	"github.com/tomoyane/grant-n-z/gserver/common/ctx"
	"github.com/tomoyane/grant-n-z/gserver/common/property"
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

var tsInstance TokenService

type TokenService interface {
	Generate(queryParam string, userEntity entity.User) (*string, *model.ErrorResBody)

	ParseToken(token string) (map[string]string, bool)

	VerifyToken(w http.ResponseWriter, r *http.Request, authType string) (*model.AuthUser, *model.ErrorResBody)

	generateSignedInToken(user *entity.User, roleId int, serviceId int) *string

	generateOperatorToken(userEntity entity.User) (*string, *model.ErrorResBody)

	generateUserToken(userEntity entity.User) (*string, *model.ErrorResBody)

	verifyOperatorToken(token string) (*model.AuthUser, *model.ErrorResBody)

	verifyUserToken(token string) (*model.AuthUser, *model.ErrorResBody)

	getAuthUserInToken(token string) (*model.AuthUser, *model.ErrorResBody)
}

type tokenServiceImpl struct {
	userService           UserService
	operatorPolicyService OperatorPolicyService
	userServiceService    UserServiceService
	appConfig             config.AppConfig
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
		userService:           GetUserServiceInstance(),
		operatorPolicyService: GetOperatorPolicyServiceInstance(),
		userServiceService:    GetUserServiceServiceInstance(),
		appConfig:             config.App,
	}
}

func (tsi tokenServiceImpl) Generate(queryParam string, userEntity entity.User) (*string, *model.ErrorResBody) {
	if strings.EqualFold(queryParam, property.AuthOperator) {
		return tsi.generateOperatorToken(userEntity)
	} else if strings.EqualFold(queryParam, "") {
		return tsi.generateUserToken(userEntity)
	} else {
		return nil, model.BadRequest("Not support type of query parameter")
	}
}

func (tsi tokenServiceImpl) ParseToken(token string) (map[string]string, bool) {
	resultMap := map[string]string{}

	parseToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(tsi.appConfig.PrivateKeyBase64), nil
	})

	if err != nil || !parseToken.Valid {
		log.Logger.Error("Error parse token validation", err.Error())
		return resultMap, false
	}

	claims := parseToken.Claims.(jwt.MapClaims)
	if _, ok := claims["username"].(string); !ok {
		log.Logger.Info("Can not get username from token")
		return resultMap, false
	}
	if _, ok := claims["user_uuid"].(string); !ok {
		log.Logger.Info("Can not get user_uuid from token")
		return resultMap, false
	}
	if _, ok := claims["user_id"].(string); !ok {
		log.Logger.Info("Can not get user_id from token")
		return resultMap, false
	}
	if _, ok := claims["expires"].(string); !ok {
		log.Logger.Info("Can not get expires from token")
		return resultMap, false
	}
	if _, ok := claims["role_id"].(string); !ok {
		log.Logger.Info("Can not get role from token")
		return resultMap, false
	}

	resultMap["username"] = claims["username"].(string)
	resultMap["user_uuid"] = claims["user_uuid"].(string)
	resultMap["user_id"] = claims["user_id"].(string)
	resultMap["expires"] = claims["expires"].(string)
	resultMap["role_id"] = claims["role"].(string)

	return resultMap, true
}

func (tsi tokenServiceImpl) VerifyToken(w http.ResponseWriter, r *http.Request, authType string) (*model.AuthUser, *model.ErrorResBody) {
	var authUser *model.AuthUser
	var err *model.ErrorResBody
	if strings.EqualFold(authType, property.AuthOperator) {
		authUser, err = tsi.verifyOperatorToken(ctx.GetToken().(string))
	} else {
		authUser, err = tsi.verifyUserToken(ctx.GetToken().(string))
	}
	if err != nil {
		return nil, err
	}

	return authUser, nil
}

func (tsi tokenServiceImpl) generateOperatorToken(userEntity entity.User) (*string, *model.ErrorResBody) {
	// TODO: Cache user data, operator_policy data
	uwo, err := tsi.userService.GetUserWithOperatorPolicyByEmail(userEntity.Email)
	if err != nil || uwo == nil {
		return nil, model.BadRequest("Failed to email or password")
	}

	if !tsi.userService.ComparePw(uwo.Password, userEntity.Password) {
		return nil, model.BadRequest("Failed to email or password")
	}

	if uwo.OperatorPolicy.RoleId != property.OperatorRoleId {
		return nil, model.BadRequest("Can not issue token")
	}

	user := entity.User{
		Id:       uwo.UserId,
		Username: uwo.Username,
		Uuid:     uwo.Uuid,
	}
	return tsi.generateSignedInToken(&user, uwo.OperatorPolicy.RoleId, 0), nil
}

func (tsi tokenServiceImpl) generateUserToken(userEntity entity.User) (*string, *model.ErrorResBody) {
	// TODO: Cache user data, user_service, service data
	uus, err := tsi.userService.GetUserWithUserServiceWithServiceByEmail(userEntity.Email)
	if err != nil || uus == nil {
		return nil, model.BadRequest("Failed to email or password")
	}

	if !tsi.userService.ComparePw(uus.User.Password, userEntity.Password) {
		return nil, model.BadRequest("Failed to email or password")
	}

	apiKey := ctx.GetApiKey().(string)
	if !strings.EqualFold(uus.Service.ApiKey, apiKey) {
		return nil, model.BadRequest("Can not issue token")
	}

	user := entity.User{
		Id:       uus.User.Id,
		Username: uus.User.Username,
		Uuid:     uus.User.Uuid,
	}
	return tsi.generateSignedInToken(&user, 0, uus.Service.Id), nil
}

func (tsi tokenServiceImpl) generateSignedInToken(user *entity.User, roleId int, serviceId int) *string {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["user_uuid"] = user.Uuid
	claims["user_id"] = strconv.Itoa(user.Id)
	claims["service_id"] = serviceId
	claims["expires"] = time.Now().Add(time.Hour * 1).String()
	claims["role_id"] = strconv.Itoa(roleId)

	signedToken, err := token.SignedString([]byte(tsi.appConfig.PrivateKeyBase64))
	if err != nil {
		log.Logger.Error("Error signed token", err.Error())
		return nil
	}

	return &signedToken
}

func (tsi tokenServiceImpl) verifyOperatorToken(token string) (*model.AuthUser, *model.ErrorResBody) {
	authUser, err := tsi.getAuthUserInToken(token)
	if err != nil {
		return nil, err
	}

	operatorRole, err := tsi.operatorPolicyService.GetByUserIdAndRoleId(authUser.UserId, authUser.RoleId)
	if operatorRole == nil || err != nil {
		log.Logger.Info("Not contain operator role or failed to query", err.ToJson())
		return nil, model.Unauthorized("Invalid token")
	}

	return authUser, nil
}

func (tsi tokenServiceImpl) verifyUserToken(token string) (*model.AuthUser, *model.ErrorResBody) {
	authUser, err := tsi.getAuthUserInToken(token)
	if err != nil {
		return nil, err
	}

	userService, err := tsi.userServiceService.GetUserServiceByUserIdAndServiceId(authUser.UserId, authUser.ServiceId)
	if userService == nil || err != nil {
		log.Logger.Info("Not contain service of user or failed to query", err.ToJson())
		return nil, model.Unauthorized("Invalid token")
	}

	return authUser, nil
}

func (tsi tokenServiceImpl) getAuthUserInToken(token string) (*model.AuthUser, *model.ErrorResBody) {
	if !strings.Contains(token, "Bearer") {
		log.Logger.Info("Not found authorization header or not contain `Bearer` in authorization header")
		return nil, model.Unauthorized("Unauthorized.")
	}

	userData, result := tsi.ParseToken(strings.Replace(token, "Bearer ", "", 1))
	if !result {
		log.Logger.Info("Failed to parse token")
		return nil, model.Unauthorized("Failed to token.")
	}

	// TODO: Cache user data
	id, _ := strconv.Atoi(userData["user_id"])
	user, err := tsi.userService.GetUserById(id)
	if err != nil {
		return nil, model.Unauthorized("Failed to token.")
	}

	if user == nil {
		log.Logger.Info("User data is null")
		return nil, model.Unauthorized("Failed to token.")
	}

	roleId, _ := strconv.Atoi(userData["role"])
	serviceId, _ := strconv.Atoi(userData["service_id"])
	return &model.AuthUser{
		Username:  user.Username,
		UserUuid:  user.Uuid,
		UserId:    user.Id,
		UserEmail: user.Email,
		ServiceId: serviceId,
		Expires:   userData["expires"],
		RoleId:    roleId,
	}, nil
}
