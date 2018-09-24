package service

import (
	"golang.org/x/crypto/bcrypt"
	"github.com/tomoyane/grant-n-z/domain/entity"
	"github.com/tomoyane/grant-n-z/domain/repository"
	"github.com/satori/go.uuid"
	"github.com/labstack/echo"
	"net/http"
	"github.com/tomoyane/grant-n-z/handler"
	"strings"
)

type UserService struct {
	UserRepository repository.UserRepository
}

func (u UserService) EncryptPw(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([] byte(password), bcrypt.DefaultCost)
	return string(hash)
}

func (u UserService) ComparePw(passwordHash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	if err != nil {
		return false
	}

	return true
}

func (u UserService) GetUserByEmail(email string) *entity.User {
	return u.UserRepository.FindByEmail(email)
}

func (u UserService) GetUserByUuid(username string, uuid string) *entity.User {
	return u.UserRepository.FindByUserNameAndUuid(username, uuid)
}

func (u UserService) insertUser(user entity.User) *entity.User {
	user.Uuid, _ = uuid.NewV4()
	user.Password = u.EncryptPw(user.Password)
	return u.UserRepository.Save(user)
}

func (u UserService) UpdateUser(user entity.User) *entity.User {
	user.Password = u.EncryptPw(user.Password)
	return u.UserRepository.Update(user)
}

func (u UserService) UpdateUserColumn(user entity.User, column string) *entity.User {
	user.Password = u.EncryptPw(user.Password)
	return u.UserRepository.UpdateUserColumn(user, column)
}

func (u UserService) PostUserData(c echo.Context, user *entity.User) (err error) {
	if err = c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, handler.BadRequest("001"))
	}

	if err = c.Validate(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, handler.BadRequest("002"))
	}

	userData := u.GetUserByEmail(user.Email)
	if userData == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, handler.InternalServerError("003"))
	}

	if len(userData.Email) > 0 {
		return echo.NewHTTPError(http.StatusConflict, handler.Conflict("004"))
	}

	userData = u.insertUser(*user)
	if userData == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, handler.InternalServerError("005"))
	}

	return nil
}

func (u UserService) PutUserColumnData(c echo.Context, tokeService TokenService, roleService RoleService,
	user *entity.User, token string, column string) (err error) {

	result := tokeService.VerifyToken(c, u, roleService, token)

	if result != nil {
		return result
	}
	if !strings.Contains(column, "username") && !strings.EqualFold(column, "email") && !strings.EqualFold(column, "password") {
		return echo.NewHTTPError(http.StatusBadRequest, handler.BadRequest(""))
	}

	if err = c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, handler.BadRequest(""))
	}

	if err = c.Validate(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, handler.BadRequest(""))
	}

	userData := u.GetUserByEmail(user.Email)
	if userData == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, handler.InternalServerError(""))
	}

	if len(userData.Email) == 0 {
		return echo.NewHTTPError(http.StatusNotFound, handler.NotFound(""))
	}

	userData = u.UpdateUserColumn(*user, column)
	if userData == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, handler.InternalServerError(""))
	}

	return nil
}