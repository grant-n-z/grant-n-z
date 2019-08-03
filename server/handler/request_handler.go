package handler

import (
	"io/ioutil"
	"net/http"
	"strings"

	"gopkg.in/go-playground/validator.v9"

	"github.com/tomoyane/grant-n-z/server/log"
	"github.com/tomoyane/grant-n-z/server/model"
)

func Interceptor(w http.ResponseWriter, r *http.Request) ([]byte, *model.ErrorResponse) {
	if err := validateHttpHeader(r); err != nil {
		http.Error(w, err.ToJson(), err.Code)
		return nil, err
	}

	bodyBytes, err := bindBody(r)
	if err != nil {
		http.Error(w, err.ToJson(), err.Code)
		return nil, err
	}

	return bodyBytes, nil
}

func ValidateHttpRequest(w http.ResponseWriter, i interface{}) *model.ErrorResponse {
	err := validator.New().Struct(i)
	if err != nil {
		err := model.BadRequest(err.Error())
		log.Logger.Warn("Error request validation", err.Detail)
		errModel := model.BadRequest("Failed to request validation.")
		http.Error(w, errModel.ToJson(), err.Code)
		return err
	}
	return nil
}

func validateHttpHeader(r *http.Request) *model.ErrorResponse {
	if strings.Contains(r.Header.Get("Authorization"), "Bearer") {
		log.Logger.Info("Need to authorization header")
		return model.Unauthorized("Need to authorization header.")
	}

	if r.Header.Get("Content-Type") != "application/json" {
		log.Logger.Info("Not allowed content-type")
		return model.BadRequest("Need to content type is only json.")
	}

	return nil
}

func verifyAuthentication() {

}

func bindBody(r *http.Request) ([]byte, *model.ErrorResponse) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Logger.Info(err.Error())
		return nil, model.InternalServerError("Error request body bind")
	}

	return body, nil
}
