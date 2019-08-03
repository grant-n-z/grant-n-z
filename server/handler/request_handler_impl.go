package handler

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"gopkg.in/go-playground/validator.v9"

	"github.com/tomoyane/grant-n-z/server/log"
	"github.com/tomoyane/grant-n-z/server/model"
	"github.com/tomoyane/grant-n-z/server/usecase/service"
)

type RequestHandlerImpl struct {
	UserService service.UserService
}

func NewRequestHandler() RequestHandler {
	return RequestHandlerImpl{
		UserService: service.NewUserService(),
	}
}

func (rh RequestHandlerImpl) InterceptHttp(w http.ResponseWriter, r *http.Request) ([]byte, *model.ErrorResponse) {
	if err := rh.validateHttpHeader(r); err != nil {
		http.Error(w, err.ToJson(), err.Code)
		return nil, err
	}

	if !(strings.Contains(r.URL.String(), "users") && strings.EqualFold(r.Method, http.MethodPost)) {
		if !strings.Contains(r.URL.String(), "oauth") {
			if err := rh.verifyServiceAuth(r.Header.Get("Authorization")); err != nil {
				http.Error(w, err.ToJson(), err.Code)
				return nil, err
			}
		}
	}

	bodyBytes, err := rh.bindRequestBody(r)
	if err != nil {
		http.Error(w, err.ToJson(), err.Code)
		return nil, err
	}

	return bodyBytes, nil
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

func (rh RequestHandlerImpl) verifyServiceAuth(token string) *model.ErrorResponse {
	if !strings.Contains(token, "Bearer") {
		log.Logger.Info("Not contain `Bearer` authorization header")
		return model.Unauthorized("Not contain `Bearer` authorization header.")
	}

	jwt := strings.Replace(token, "Bearer ", "", 1)
	userData, result := rh.UserService.ParseJwt(jwt)
	if !result {
		log.Logger.Info("Failed to parse token")
		return model.Unauthorized("Failed to token.")
	}

	id, _ := strconv.Atoi(userData["user_id"])
	user, err := rh.UserService.GetUserById(id)
	if err != nil {
		return err
	}

	if user == nil {
		log.Logger.Info("User data is null")
		return model.Unauthorized("Failed to token.")
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
