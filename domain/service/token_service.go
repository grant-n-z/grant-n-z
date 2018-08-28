package service

import (
	"github.com/tomoyane/grant-n-z/domain/entity"
	"github.com/tomoyane/grant-n-z/domain/repository"
	"time"
	"github.com/satori/go.uuid"
	"github.com/dgrijalva/jwt-go"
	"github.com/tomoyane/grant-n-z/domain"
	"net/http"
)

type TokenService struct {
	TokenRepository repository.TokenRepository
}

func (t TokenService) GenerateJwt(username string, userUuid uuid.UUID) string {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = username
	claims["user_uuid"] = userUuid.String()
	claims["exp"] = time.Now().Add(time.Hour * 365).Unix()

	signedToken, err := token.SignedString([]byte("secret"))
	if err != nil {
		domain.ErrorResponse{}.Print(http.StatusInternalServerError, "failed generate jwt", "")
		return ""
	}

	return signedToken
}

func (t TokenService) GetTokenByUserId(userId string) *entity.Token {
	return t.TokenRepository.FindByUserId(userId)
}

func (t TokenService) InsertToken(userUuid uuid.UUID, token string, refreshToken string) *entity.Token {
	data := entity.Token{
		TokenType: "Bearer",
		Token: token,
		RefreshToken: refreshToken,
		UserUuid: userUuid,
	}
	return t.TokenRepository.Save(data)
}