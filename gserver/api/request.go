package api

import (
	"strings"

	"io/ioutil"
	"net/http"

	"gopkg.in/go-playground/validator.v9"

	"github.com/tomoyane/grant-n-z/gserver/cache"
	"github.com/tomoyane/grant-n-z/gserver/common/ctx"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
	"github.com/tomoyane/grant-n-z/gserver/service"
)

var rhInstance Request

type Request interface {
	// Http interceptor
	// Set request scope
	Intercept(w http.ResponseWriter, r *http.Request, authType string) ([]byte, *model.ErrorResBody)

	// Validate http request
	ValidateBody(w http.ResponseWriter, i interface{}) *model.ErrorResBody

	// Validate http header
	validateHeader(r *http.Request) *model.ErrorResBody

	// Bind http request body
	bindBody(r *http.Request) ([]byte, *model.ErrorResBody)
}

type RequestImpl struct {
	tokenService service.TokenService
	userService  service.UserService
	redisClient  cache.RedisClient
}

func GetRequestInstance() Request {
	if rhInstance == nil {
		rhInstance = NewRequest()
	}
	return rhInstance
}

func NewRequest() Request {
	log.Logger.Info("New `request` instance")
	log.Logger.Info("Inject `AuthService` to `request`")
	return RequestImpl{
		tokenService: service.GetTokenServiceInstance(),
		userService:  service.GetUserServiceInstance(),
		redisClient:  cache.GetRedisClientInstance(),
	}
}

func (rh RequestImpl) Intercept(w http.ResponseWriter, r *http.Request, authType string) ([]byte, *model.ErrorResBody) {
	var authUser *model.AuthUser
	var err *model.ErrorResBody
	if !strings.EqualFold(authType, "") {
		token := r.Header.Get("Authorization")
		authUser, err = rh.tokenService.VerifyToken(w, r, authType, token)
		if err != nil {
			model.WriteError(w, err.ToJson(), err.Code)
			return nil, err
		}

		// Set user id request scope
		ctx.SetUserId(authUser.UserId)
		// Set user uuid request scope
		ctx.SetUserUuid(authUser.UserUuid.String())
		// Set service id request scope
		ctx.SetServiceId(authUser.ServiceId)
	}

	if err := rh.validateHeader(r); err != nil {
		model.WriteError(w, err.ToJson(), err.Code)
		return nil, err
	}

	// Set api key request scope
	ctx.SetApiKey(r.Header.Get("Api-Key"))

	bodyBytes, err := rh.bindBody(r)
	if err != nil {
		model.WriteError(w, err.ToJson(), err.Code)
		return nil, err
	}

	return bodyBytes, nil
}

func (rh RequestImpl) ValidateBody(w http.ResponseWriter, i interface{}) *model.ErrorResBody {
	err := validator.New().Struct(i)
	if err != nil {
		log.Logger.Info("request is invalid")
		errModel := model.BadRequest("Failed to request validation.")
		model.WriteError(w, errModel.ToJson(), errModel.Code)
		return errModel
	}
	return nil
}

func (rh RequestImpl) validateHeader(r *http.Request) *model.ErrorResBody {
	if r.Header.Get("Content-Type") != "application/json" {
		log.Logger.Info("Not allowed content-type")
		return model.BadRequest("Need to content type is only json.")
	}
	return nil
}

func (rh RequestImpl) bindBody(r *http.Request) ([]byte, *model.ErrorResBody) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Logger.Info(err.Error())
		return nil, model.InternalServerError("WriteError request body bind")
	}

	return body, nil
}
