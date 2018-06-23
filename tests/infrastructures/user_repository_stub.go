package infrastructures

import (
	"grant-n-z/app/infrastructures"
	"grant-n-z/app/domains/entity"
	"github.com/satori/go.uuid"
)

type UserRepositoryStub struct {}

func (r UserRepositoryStub) NewUserRepository() infrastructures.UserRepository {
	return &UserRepositoryStub{}
}

func (r UserRepositoryStub) FindByEmail(email string) *entity.Users {
	users := entity.Users{
		Id: 1,
		Uuid: uuid.Must(uuid.NewV4()).String(),
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
