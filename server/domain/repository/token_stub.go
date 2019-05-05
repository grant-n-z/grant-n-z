package repository

import (
	"github.com/tomoyane/grant-n-z/server/domain/entity"
	"github.com/satori/go.uuid"
)

type TokenRepositoryStub struct {
}

func (r TokenRepositoryStub) FindByUserUuid(userUuidStr string) *entity.Token {
	userUuid, _ := uuid.FromString(userUuidStr)
	token := entity.Token{
		Id: 1,
		TokenType: "Bearer",
		Token: "testToken",
		RefreshToken: "testRefreshToken",
		UserUuid: userUuid,
	}
	return &token
}

func (r TokenRepositoryStub) Save(token entity.Token) *entity.Token {
	return &token
}
