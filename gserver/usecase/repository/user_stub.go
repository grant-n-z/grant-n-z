package repository

import (
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/satori/go.uuid"
)

type UserRepositoryStub struct {
}

func (u UserRepositoryStub) FindByEmail(email string) *entity.User {
	userUuid, _ :=  uuid.NewV4()

	if email == "test2@gmail.com"{
		return &entity.User{
			Username: "test",
			Email: "test2@gmail.com",
			Uuid: userUuid,
			Password: "$2a$10$yHVbM8iJBbqCHW4z3lq9KuV1m6s2TY2Z2214XPp4fpP/7JxjQDu72",
		}
	}

	return &entity.User{
		Username: "test",
		Email: "",
		Uuid: userUuid,
		Password: "$2a$10$yHVbM8iJBbqCHW4z3lq9KuV1m6s2TY2Z2214XPp4fpP/7JxjQDu72",
	}
}

func (u UserRepositoryStub) FindByUserNameAndUuid(username string, uuidStr string) *entity.User  {
	userUuid, _ := uuid.FromString(uuidStr)

	return &entity.User{
		Username: username,
		Email: "test2@gmail.com",
		Uuid: userUuid,
		Password: "$2a$10$yHVbM8iJBbqCHW4z3lq9KuV1m6s2TY2Z2214XPp4fpP/7JxjQDu72",
	}
}

func (u UserRepositoryStub) Save(user entity.User) *entity.User {
	return &user
}

func (u UserRepositoryStub) Update(user entity.User) *entity.User {
	return &user
}

