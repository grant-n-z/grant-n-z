package handler

import (
	"io/ioutil"
	"net/http"

	"gopkg.in/go-playground/validator.v9"

	"github.com/tomoyane/grant-n-z/server/log"
	"github.com/tomoyane/grant-n-z/server/model"
)

func Interceptor(w http.ResponseWriter, r *http.Request) ([]byte, *model.ErrorResponse) {
	if err := validateHttpHeader(r); err != nil {
		log.Logger.Warn("error http validation", err.Detail)
		http.Error(w, err.ToJson(), err.Code)
		return nil, err
	}

	bodyBytes, err := bindBody(r)
	if err != nil {
		log.Logger.Warn("error request bind", err.Detail)
		http.Error(w, err.ToJson(), err.Code)
		return nil, err
	}

	return bodyBytes, nil
}

func ValidateHttpRequest(w http.ResponseWriter, i interface{}) *model.ErrorResponse {
	err := validator.New().Struct(i)
	if err != nil {
		err := model.BadRequest(err.Error())
		log.Logger.Warn("error request validation", err.Detail)
		http.Error(w, err.ToJson(), err.Code)
		return err
	}
	return nil
}

func validateHttpHeader(r *http.Request) *model.ErrorResponse {
	if r.Header.Get("Content-Type") != "application/json" {
		log.Logger.Warn("not allow content-type")
		return model.BadRequest()
	}

	return nil
}

func verifyAuthentication() {

}

func bindBody(r *http.Request) ([]byte, *model.ErrorResponse) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return nil, model.InternalServerError(err.Error())
	}

	return body, nil
}
