package stub

import (
	"github.com/tomoyane/grant-n-z/domain/entity"
	"github.com/satori/go.uuid"
)

type UserRepositoryStub struct {
}

func (r UserRepositoryStub) FindByEmail(email string) *entity.User {
	users := entity.User{
		Username: "test",
		Email: "test@gmail.com",
		Uuid: uuid.NewV4(),
		Password: "test12345",
	}
	return &users
}

func (r UserRepositoryStub) Save(user entity.User) *entity.User {
	return &user
}

func (r UserRepositoryStub) Update(user entity.User) *entity.User {
	return &user
}

