package stub

import (
	"github.com/tomoyane/grant-n-z/domain/entity"
	"github.com/satori/go.uuid"
)

type UserRepositoryStub struct {
}

func (r UserRepositoryStub) FindByEmail(email string) *entity.User {
	if email == "test2@gmail.com"{
		return &entity.User{
			Username: "test",
			Email: "test2@gmail.com",
			Uuid: uuid.NewV4(),
			Password: "test12345",
		}
	}

	return &entity.User{
		Username: "test",
		Email: "",
		Uuid: uuid.NewV4(),
		Password: "test12345",
	}
}

func (r UserRepositoryStub) Save(user entity.User) *entity.User {
	return &user
}

func (r UserRepositoryStub) Update(user entity.User) *entity.User {
	return &user
}
