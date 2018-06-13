package service

import (
	"authentication-server/tests"
	"authentication-server/app/domains/service"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceTest struct {
	tests.AppTest
}

var userService = service.UserService{}

func (t UserServiceTest) Before() {
}

func (t UserServiceTest) After() {
}

func (t UserServiceTest) TestEncryptPwOk() {
	var password = "test"
	var encryptPassword = userService.EncryptPw(password)

	var result = bcrypt.CompareHashAndPassword(
		[]byte(encryptPassword),
		[]byte(password),
	)

	t.AssertEqual(result, nil)
}

func (t UserServiceTest) TestEncryptPwNotMatching() {
	var password = "test"
	var encryptPassword = userService.EncryptPw(password)

	var result = bcrypt.CompareHashAndPassword(
		[]byte(encryptPassword),
		[]byte("bad_pass"),
	)

	t.AssertNotEqual(result, nil)
}
