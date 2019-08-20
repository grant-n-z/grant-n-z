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

type RequestHandlerImpl struct {
	AuthService service.AuthService
}

func NewRequestHandler() RequestHandler {
	log.Logger.Info("Inject `AuthService` to `RequestHandler`")
	return RequestHandlerImpl{
		AuthService: service.NewAuthService(),
	}
}

func (rh RequestHandlerImpl) InterceptHttp(w http.ResponseWriter, r *http.Request) ([]byte, *model.ErrorResponse) {
	if err := rh.validateHttpHeader(r); err != nil {
		http.Error(w, err.ToJson(), err.Code)
		return nil, err
	}

	bodyBytes, err := rh.bindRequestBody(r)
	if err != nil {
		http.Error(w, err.ToJson(), err.Code)
		return nil, err
	}

	return bodyBytes, nil
}

func (rh RequestHandlerImpl) VerifyToken(w http.ResponseWriter, r *http.Request, authType string) (*model.AuthUser, *model.ErrorResponse) {
	switch authType {
	case property.AuthOperator:
		authUser, err := rh.AuthService.VerifyOperatorMember(r.Header.Get("GrantNZ-Operator-Auth"))
		if err != nil {
			http.Error(w, err.ToJson(), err.Code)
			return nil, err
		}
		return authUser, err

	case property.AuthUser:
		authUser, err := rh.AuthService.VerifyServiceMember(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, err.ToJson(), err.Code)
			return nil, err
		}
		return authUser, err
	}

	return nil, nil
}

func (rh RequestHandlerImpl) ValidateHttpRequest(w http.ResponseWriter, i interface{}) *model.ErrorResponse {
	err := validator.New().Struct(i)
	if err != nil {
		log.Logger.Warn("Error request validation")
		errModel := model.BadRequest("Failed to request validation.")
		http.Error(w, errModel.ToJson(), errModel.Code)
		return errModel
	}
	return nil
}

func (rh RequestHandlerImpl) validateHttpHeader(r *http.Request) *model.ErrorResponse {
	if r.Header.Get("Content-Type") != "application/json" {
		log.Logger.Info("Not allowed content-type")
		return model.BadRequest("Need to content type is only json.")
	}

	return nil
}

func (rh RequestHandlerImpl) bindRequestBody(r *http.Request) ([]byte, *model.ErrorResponse) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Logger.Info(err.Error())
		return nil, model.InternalServerError("Error request body bind")
	}

	return body, nil
}
