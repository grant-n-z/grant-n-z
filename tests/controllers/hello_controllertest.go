package controllers

import (
	"authentication-server/tests"
)

type HelloControllerTest struct {
	tests.AppTest
}

func (t HelloControllerTest) TestIndexOk() {
	t.Get("/hello")
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")
}

func (t HelloControllerTest) TestIndexBadResponse() {
	inCorrectHello := map[string]string{"key": "hello word"}
	correctHello := map[string]string{"key": "hello world"}

	t.Get("/hello")
	t.AssertNotEqual(inCorrectHello, correctHello)
	t.AssertContentType("application/json; charset=utf-8")
}
