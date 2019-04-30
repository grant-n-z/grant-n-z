package handler

import (
	"io/ioutil"
	"net/http"

	"gopkg.in/go-playground/validator.v9"

	"github.com/tomoyane/grant-n-z/server/domain/entity"
	"github.com/tomoyane/grant-n-z/server/log"
)

func Interceptor(w http.ResponseWriter, r *http.Request, method string) ([]byte, *entity.ErrorResponse) {
	if err := httpValidator(r, method); err != nil {
		log.Logger.Warn("error http validation", err.Detail)
		http.Error(w, err.ToJson(), err.Code)
		return nil, err
	}

	bodyBytes, err := bind(r)
	if err != nil {
		log.Logger.Warn("error request bind", err.Detail)
		http.Error(w, err.ToJson(), err.Code)
		return nil, err
	}

	return bodyBytes, nil
}

func httpValidator(r *http.Request, method string) *entity.ErrorResponse {
	if r.Method != method {
		log.Logger.Warn("not allow http method")
		return entity.MethodNotAllowed()
	}

	if r.Header.Get("Content-Type") != "application/json" {
		log.Logger.Warn("not allow content-type")
		return entity.BadRequest()
	}

	return nil
}

func bind(r *http.Request) ([]byte, *entity.ErrorResponse) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return nil, entity.InternalServerError(err.Error())
	}

	return body, nil
}

func BodyValidator(w http.ResponseWriter, i interface{}) *entity.ErrorResponse {
	err := validator.New().Struct(i)
	if err != nil {
		err := entity.BadRequest(err.Error())
		log.Logger.Warn("error request validation", err.Detail)
		http.Error(w, err.ToJson(), err.Code)
		return err
	}
	return nil
}
