package service

import (
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"
	"github.com/satori/go.uuid"

	"github.com/tomoyane/grant-n-z/gserver/cache"
	"github.com/tomoyane/grant-n-z/gserver/common/config"
	"github.com/tomoyane/grant-n-z/gserver/common/driver"
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
	"github.com/tomoyane/grant-n-z/gserver/repository"
)

var usInstance UserService

type UserService interface {
	EncryptPw(password string) string

	ComparePw(passwordHash string, password string) bool

	GetUserById(id int) (*entity.User, *model.ErrorResBody)

	GetUserByEmail(email string) (*entity.User, *model.ErrorResBody)

	GetUserWithRoleByEmail(email string) (*model.UserOperatorPolicy, *model.ErrorResBody)

	InsertUser(user *entity.User) (*entity.User, *model.ErrorResBody)

	InsertUserWithService(user *entity.User, userService *entity.UserService) (*entity.User, *model.ErrorResBody)

	UpdateUser(user *entity.User) (*entity.User, *model.ErrorResBody)

	GenerateJwt(user *entity.User, roleId int) *string

	ParseJwt(token string) (map[string]string, bool)
}

type userServiceImpl struct {
	userRepository repository.UserRepository
	appConfig      config.AppConfig
	redisClient    cache.RedisClient
}

func GetUserServiceInstance() UserService {
	if usInstance == nil {
		usInstance = NewUserService()
	}
	return usInstance
}

func NewUserService() UserService {
	log.Logger.Info("New `UserService` instance")
	log.Logger.Info("Inject `UserRepository`, `AppConfig`, `RedisClient` to `UserService`")
	return userServiceImpl{
		userRepository: repository.GetUserRepositoryInstance(driver.Db),
		appConfig:      config.App,
		redisClient:    cache.GetRedisClientInstance(),
	}
}

func (us userServiceImpl) EncryptPw(password string) string {
	hash, err := bcrypt.GenerateFromPassword([] byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Logger.Info("Error password hash", err.Error())
		return ""
	}

	return string(hash)
}

func (us userServiceImpl) ComparePw(passwordHash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	if err != nil {
		log.Logger.Info("Error compare password", err.Error())
		return false
	}

	return true
}

func (us userServiceImpl) GetUserById(id int) (*entity.User, *model.ErrorResBody) {
	return us.userRepository.FindById(id)
}

func (us userServiceImpl) GetUserByEmail(email string) (*entity.User, *model.ErrorResBody) {
	return us.userRepository.FindByEmail(email)
}

func (us userServiceImpl) GetUserWithRoleByEmail(email string) (*model.UserOperatorPolicy, *model.ErrorResBody) {
	return us.userRepository.FindUserWithRoleByEmail(email)
}

func (us userServiceImpl) InsertUser(user *entity.User) (*entity.User, *model.ErrorResBody) {
	user.Uuid, _ = uuid.NewV4()
	user.Password = us.EncryptPw(user.Password)
	return us.userRepository.Save(*user)
}

func (us userServiceImpl) InsertUserWithService(user *entity.User, userService *entity.UserService) (*entity.User, *model.ErrorResBody) {
	user.Uuid, _ = uuid.NewV4()
	user.Password = us.EncryptPw(user.Password)
	return us.userRepository.SaveUserWithUserService(*user, userService)
}

func (us userServiceImpl) UpdateUser(user *entity.User) (*entity.User, *model.ErrorResBody) {
	user.Password = us.EncryptPw(user.Password)
	return us.userRepository.Update(*user)
}

func (us userServiceImpl) GenerateJwt(user *entity.User, roleId int) *string {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["user_uuid"] = user.Uuid
	claims["user_id"] = strconv.Itoa(user.Id)
	//claims["service_id"] = handler.ApiKey
	claims["expires"] = time.Now().Add(time.Hour * 1).String()
	claims["role"] = strconv.Itoa(roleId)

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
		return []byte(us.appConfig.PrivateKeyBase64), nil
	})

	if err != nil || !parseToken.Valid {
		log.Logger.Error("Error parse token validation", err.Error())
		return resultMap, false
	}

	claims := parseToken.Claims.(jwt.MapClaims)
	if _, ok := claims["username"].(string); !ok {
		log.Logger.Info("Can not get username from token")
		return resultMap, false
	}

	if _, ok := claims["user_uuid"].(string); !ok {
		log.Logger.Info("Can not get user_uuid from token")
		return resultMap, false
	}

	if _, ok := claims["user_id"].(string); !ok {
		log.Logger.Info("Can not get user_id from token")
		return resultMap, false
	}

	if _, ok := claims["expires"].(string); !ok {
		log.Logger.Info("Can not get expires from token")
		return resultMap, false
	}

	if _, ok := claims["role"].(string); !ok {
		log.Logger.Info("Can not get role from token")
		return resultMap, false
	}

	resultMap["username"] = claims["username"].(string)
	resultMap["user_uuid"] = claims["user_uuid"].(string)
	resultMap["user_id"] = claims["user_id"].(string)
	resultMap["expires"] = claims["expires"].(string)
	resultMap["role"] = claims["role"].(string)

	return resultMap, true
}
