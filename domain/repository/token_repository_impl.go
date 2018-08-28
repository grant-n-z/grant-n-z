package repository

import (
	"github.com/tomoyane/grant-n-z/domain/entity"
	"github.com/tomoyane/grant-n-z/infra"
)

type TokenRepositoryImpl struct {
}

// Find users by token.user_id
func (t TokenRepositoryImpl) FindByUserId(userId string) *entity.Token {
	token := entity.Token{}

	if err := infra.Db.Where("user_id = ?", userId).First(&token).Error; err != nil {
		if err.Error() == "record not found" {
			return &entity.Token{}
		}
		return nil
	}

	return &token
}

// Save to token
func (t TokenRepositoryImpl) Save(token entity.Token) *entity.Token {
	if err := infra.Db.Where("user_uuid = ?", token.UserUuid).Assign(token).FirstOrCreate(&token).Error; err != nil {
		return nil
	}

	return &token
}