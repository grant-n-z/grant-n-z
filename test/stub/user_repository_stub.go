package stub

import (
	"github.com/tomoyane/grant-n-z/domain/entity"
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
			Password: "test12345",
		}
	}

	return &entity.User{
		Username: "test",
		Email: "",
		Uuid: userUuid,
		Password: "test12345",
	}
}

func (u UserRepositoryStub) FindByUserNameAndUuid(username string, uuidStr string) *entity.User  {
	userUuid, _ := uuid.FromString(uuidStr)

	return &entity.User{
		Username: username,
		Email: "test2@gmail.com",
		Uuid: userUuid,
		Password: "test12345",
	}
}

func (u UserRepositoryStub) Save(user entity.User) *entity.User {
	return &user
}

func (u UserRepositoryStub) Update(user entity.User) *entity.User {
	return &user
}

