package repository

import (
	"github.com/satori/go.uuid"
	"github.com/tomoyane/grant-n-z/domain/entity"
	"github.com/tomoyane/grant-n-z/infra"
)

type MemberRepositoryImpl struct {
}

func (m MemberRepositoryImpl) FindByUserUuidAndServiceUuid(userUuid uuid.UUID, serviceUuid uuid.UUID) *entity.Member {
	member := entity.Member{}

	if err := infra.Db.Where("user_uuid = ? AND service_uuid", userUuid, serviceUuid).First(&member).Error; err != nil {
		if err.Error() == "record not found" {
			return &entity.Member{}
		}
		return nil
	}

	return &member
}

func (m MemberRepositoryImpl) Save(member entity.Member) *entity.Member {
	if err := infra.Db.Create(&member).Error; err != nil {
		return nil
	}

	return &member
}