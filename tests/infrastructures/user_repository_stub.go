package infrastructures

import (
	"github.com/satori/go.uuid"
	"github.com/tomo0111/grant-n-z/app/infrastructures"
	"github.com/tomo0111/grant-n-z/app/domains/entity"
)

type UserRepositoryStub struct {}

func (r UserRepositoryStub) NewUserRepository() infrastructures.UserRepository {
	return &UserRepositoryStub{}
}

func (r UserRepositoryStub) FindByEmail(email string) *entity.Users {
	users := entity.Users{
		Id: 1,
		Uuid: uuid.NewV4().String(),
		Username: "test",
		Email: email,
		Password: "testtest",
	}

	return &users
}

func (r UserRepositoryStub) Save(users entity.Users) *entity.Users {
	return &users
}

func (r UserRepositoryStub) Update(users entity.Users) *entity.Users {
	return &users
}
