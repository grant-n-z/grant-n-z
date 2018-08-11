package domain

import "gopkg.in/go-playground/validator.v9"

type GrantValidator struct {
	Validator *validator.Validate
}

func (gv *GrantValidator) Validate(i interface{}) error {
	return gv.Validator.Struct(i)
}