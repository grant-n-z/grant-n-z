package service

import (
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"
	"github.com/satori/go.uuid"

	"github.com/tomoyane/grant-n-z/server/config"
	"github.com/tomoyane/grant-n-z/server/entity"
	"github.com/tomoyane/grant-n-z/server/log"
	"github.com/tomoyane/grant-n-z/server/model"

	"github.com/tomoyane/grant-n-z/server/usecase/repository"
)

type userServiceImpl struct {
	userRepository repository.UserRepository
	appConfig      config.AppConfig
}

func NewUserService() UserService {
	log.Logger.Info("Inject `UserRepository` to `UserService`")
	return userServiceImpl{
		userRepository: repository.NewUserRepository(config.Db),
		appConfig:      config.App,
	}
}

func (us userServiceImpl) EncryptPw(password string) string {
	hash, err := bcrypt.GenerateFromPassword([] byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Logger.Error("Error password hash", err.Error())
		return ""
	}

	return string(hash)
}

func (us userServiceImpl) ComparePw(passwordHash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	if err != nil {
		log.Logger.Error("Error compare password", err.Error())
		return false
	}

	return true
}

func (us userServiceImpl) GetUserById(id int) (*entity.User, *model.ErrorResponse) {
	return us.userRepository.FindById(id)
}

func (us userServiceImpl) GetUserByEmail(email string) (*entity.User, *model.ErrorResponse) {
	return us.userRepository.FindByEmail(email)
}

func (us userServiceImpl) InsertUser(user *entity.User) (*entity.User, *model.ErrorResponse) {
	user.Uuid, _ = uuid.NewV4()
	user.Password = us.EncryptPw(user.Password)
	return us.userRepository.Save(*user)
}

func (us userServiceImpl) GenerateJwt(user *entity.User, role string) *string {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["user_uuid"] = user.Uuid
	claims["expires"] = time.Now().Add(time.Hour * 1).String()
	claims["role"] = role

	signedToken, err := token.SignedString([]byte(us.appConfig.PrivateKeyBase64))
	if err != nil {
		log.Logger.Error("Error signed token", err.Error())
		return nil
	}

	return &signedToken
}

func (us userServiceImpl) ParseJwt(token string) (map[string]string, bool) {
	resultMap := map[string]string{}

	parseToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		log.Logger.Error("Error parse token")
		return []byte(us.appConfig.PrivateKeyBase64), nil
	})

	if err != nil || !parseToken.Valid {
		return resultMap, false
	}

	claims := parseToken.Claims.(jwt.MapClaims)
	if _, ok := claims["username"].(string); !ok {
		return resultMap, false
	}

	if _, ok := claims["user_uuid"].(string); !ok {
		return resultMap, false
	}

	if _, ok := claims["expires"].(string); !ok {
		return resultMap, false
	}

	if _, ok := claims["role"].(string); !ok {
		return resultMap, false
	}

	resultMap["username"] = claims["username"].(string)
	resultMap["user_uuid"] = claims["user_uuid"].(string)
	resultMap["expires"] = claims["expires"].(string)
	resultMap["role"] = claims["role"].(string)

	return resultMap, true
}
