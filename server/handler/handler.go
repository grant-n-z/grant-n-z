package handler

import (
	"io/ioutil"
	"net/http"

	"gopkg.in/go-playground/validator.v9"

	"github.com/tomoyane/grant-n-z/server/domain/entity"
)

func Interceptor(r *http.Request) ([]byte, *entity.ErrorResponse) {
	if err := httpValidator(r, http.MethodPost); err != nil {
		return nil, err
	}

	bodyBytes, err := bind(r)
	if err != nil {
		return nil, err
	}

	return bodyBytes, nil
}

func httpValidator(r *http.Request, method string) *entity.ErrorResponse {
	if r.Method != method {
		return entity.MethodNotAllowed()
	}

	if r.Header.Get("Content-Type") != "application/json" {
		return entity.BadRequest()
	}

	return nil
}

func bind(r *http.Request) ([]byte, *entity.ErrorResponse) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return nil, entity.InternalServerError()
	}

	return body, nil
}

func BodyValidator(i interface{}) *entity.ErrorResponse {
	err := validator.New().Struct(i)
	if err != nil {
		return entity.BadRequest()
	}
	return nil
}
