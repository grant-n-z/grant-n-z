package service

import (
	"github.com/tomoyane/grant-n-z/gserver/usecase/cache"
	"strconv"
	"strings"

	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

type AuthServiceImpl struct {
	UserService UserService
	RedisClient cache.RedisClient
}

func NewAuthService() AuthService {
	return AuthServiceImpl{
		UserService:   NewUserService(),
		RedisClient: cache.NewRedisClient(),
	}
}

func (as AuthServiceImpl) VerifyOperatorMember(token string) *model.ErrorResponse {
	if !strings.Contains(token, "Bearer") {
		log.Logger.Info("Not contain `Bearer` authorization header")
		return model.Unauthorized("Not contain `Bearer` authorization header.")
	}

	jwt := strings.Replace(token, "Bearer ", "", 1)
	userData, result := as.UserService.ParseJwt(jwt)
	if !result {
		log.Logger.Info("Failed to parse token")
		return model.Unauthorized("Failed to token.")
	}

	id, _ := strconv.Atoi(userData["user_id"])
	user, err := as.UserService.GetUserById(id)
	if err != nil {
		return err
	}

	if user == nil {
		log.Logger.Info("User data is null")
		return model.Unauthorized("Failed to token.")
	}
	return nil
}

func (as AuthServiceImpl) VerifyServiceMember(token string) *model.ErrorResponse {

	return nil
}
