package handler

import (
	"io/ioutil"
	"net/http"

	"gopkg.in/go-playground/validator.v9"

	"github.com/tomoyane/grant-n-z/gserver/common/property"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
	"github.com/tomoyane/grant-n-z/gserver/service"
)

var rhInstance RequestHandler

type RequestHandler interface {
	// Http interceptor
	// Verify http header
	// If it has request body, verify bind request body
	InterceptHttp(w http.ResponseWriter, r *http.Request) ([]byte, *model.ErrorResBody)

	// If need to verify authentication, verify authentication and authorization
	VerifyToken(w http.ResponseWriter, r *http.Request, authType string) (*model.AuthUser, *model.ErrorResBody)

	// Validate http request
	ValidateHttpRequest(w http.ResponseWriter, i interface{}) *model.ErrorResBody

	// Validate http header
	validateHttpHeader(r *http.Request) *model.ErrorResBody

	// Bind http request body
	bindRequestBody(r *http.Request) ([]byte, *model.ErrorResBody)
}

type RequestHandlerImpl struct {
	AuthService service.AuthService
}

func GetRequestHandlerInstance() RequestHandler {
	if rhInstance == nil {
		rhInstance = NewRequestHandler()
	}
	return rhInstance
}

func NewRequestHandler() RequestHandler {
	log.Logger.Info("New `RequestHandler` instance")
	log.Logger.Info("Inject `AuthService` to `RequestHandler`")
	return RequestHandlerImpl{
		AuthService: service.GetAuthServiceInstance(),
	}
}

func (rh RequestHandlerImpl) InterceptHttp(w http.ResponseWriter, r *http.Request) ([]byte, *model.ErrorResBody) {
	if err := rh.validateHttpHeader(r); err != nil {
		model.Error(w, err.ToJson(), err.Code)
		return nil, err
	}

	//apiKey := r.Header.Get("Api-Key")
	//log.Logger.Info("Api key =", apiKey)
	//if strings.EqualFold(apiKey, "") {
	//	property.ApiKey = apiKey
	//}

	bodyBytes, err := rh.bindRequestBody(r)
	if err != nil {
		model.Error(w, err.ToJson(), err.Code)
		return nil, err
	}

	return bodyBytes, nil
}

func (rh RequestHandlerImpl) VerifyToken(w http.ResponseWriter, r *http.Request, authType string) (*model.AuthUser, *model.ErrorResBody) {
	switch authType {
	case property.AuthOperator:
		authUser, err := rh.AuthService.VerifyOperatorMember(r.Header.Get("Operator-Authorization"))
		if err != nil {
			model.Error(w, err.ToJson(), err.Code)
			return nil, err
		}
		return authUser, err

	case property.AuthUser:
		authUser, err := rh.AuthService.VerifyServiceMember(r.Header.Get("Authorization"))
		if err != nil {
			model.Error(w, err.ToJson(), err.Code)
			return nil, err
		}
		return authUser, err
	}

	return nil, nil
}

func (rh RequestHandlerImpl) ValidateHttpRequest(w http.ResponseWriter, i interface{}) *model.ErrorResBody {
	err := validator.New().Struct(i)
	if err != nil {
		log.Logger.Info("Request is invalid")
		errModel := model.BadRequest("Failed to request validation.")
		model.Error(w, errModel.ToJson(), errModel.Code)
		return errModel
	}
	return nil
}

func (rh RequestHandlerImpl) validateHttpHeader(r *http.Request) *model.ErrorResBody {
	if r.Header.Get("Content-Type") != "application/json" {
		log.Logger.Info("Not allowed content-type")
		return model.BadRequest("Need to content type is only json.")
	}
	return nil
}

func (rh RequestHandlerImpl) bindRequestBody(r *http.Request) ([]byte, *model.ErrorResBody) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Logger.Info(err.Error())
		return nil, model.InternalServerError("Error request body bind")
	}

	return body, nil
}
