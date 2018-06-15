package service

import (
	"authentication-server/tests"
	"authentication-server/app/domains/service"
	"golang.org/x/crypto/bcrypt"

	"authentication-server/tests/infrastructures"
	"authentication-server/app/domains/entity"
	"github.com/satori/go.uuid"
)

type UserServiceTest struct {
	tests.AppTest
}

var userService = service.UserService{}
var userRepositoryStub = infrastructures.UserRepositoryStub{}.NewUserRepository()

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

func (t UserServiceTest) TestGetUserByEmailOk() {
	var userName = "test"
	var email = "test@gmail.com"
	var password = "testtest"

	users := userRepositoryStub.FindByEmail(email)

	t.AssertEqual(userName, users.Username)
	t.AssertEqual(email, users.Email)
	t.AssertEqual(password, users.Password)
	t.AssertNotEqual(users.Uuid, nil)
}

func (t UserServiceTest) TestInsertUserOk() {

	users := entity.Users{
		Id: 1,
		Uuid: uuid.Must(uuid.NewV4()).String(),
		Username: "test",
		Email: "test@gmail.com",
		Password: "testtest",
	}

	userResponse := userRepositoryStub.Save(users)

	t.AssertEqual(users.Id, userResponse.Id)
	t.AssertEqual(users.Uuid, userResponse.Uuid)
	t.AssertEqual(users.Username, userResponse.Username)
	t.AssertEqual(users.Email, userResponse.Email)
	t.AssertEqual(users.Password, userResponse.Password)
}
