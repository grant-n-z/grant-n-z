package api

import (
	"io/ioutil"
	"net/http"

	"gopkg.in/go-playground/validator.v9"

	"github.com/tomoyane/grant-n-z/gnz/cache"
	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
	"github.com/tomoyane/grant-n-z/gnzserver/service"
)

const (
	Authorization = "Authorization"
	Key           = "Api-Key"
	ContentType   = "Content-Type"
)

var rhInstance Request

type Request interface {
	// Http request interceptor
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
	return RequestImpl{
		tokenService: service.GetTokenServiceInstance(),
		userService:  service.GetUserServiceInstance(),
		redisClient:  cache.GetRedisClientInstance(),
	}
}

func (rh RequestImpl) Intercept(w http.ResponseWriter, r *http.Request, authType string) ([]byte, *model.ErrorResBody) {
	return nil, nil
}

func (rh RequestImpl) ValidateBody(w http.ResponseWriter, i interface{}) *model.ErrorResBody {
	err := validator.New().Struct(i)
	if err != nil {
		log.Logger.Info("request is invalid", err.Error())
		errModel := model.BadRequest("Failed to request validation.")
		model.WriteError(w, errModel.ToJson(), errModel.Code)
		return errModel
	}
	return nil
}

func (rh RequestImpl) validateHeader(r *http.Request) *model.ErrorResBody {
	if r.Method != http.MethodGet {
		if r.Header.Get(ContentType) != "application/json" {
			log.Logger.Info("Not allowed content-type")
			return model.BadRequest("Need to content type is only json.")
		}
	}
	return nil
}

func (rh RequestImpl) bindBody(r *http.Request) ([]byte, *model.ErrorResBody) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Logger.Info(err.Error())
		return nil, model.InternalServerError("Failed to request body bind")
	}
	return body, nil
}
