package v1

import (
	"github.com/satori/go.uuid"
	"github.com/revel/revel"
	"gopkg.in/go-playground/validator.v9"
	"strings"
	"github.com/tomo0111/grant-n-z/app/controllers/base"
	"github.com/tomo0111/grant-n-z/app/domains/service"
	"github.com/tomo0111/grant-n-z/app/domains/entity"
)

type UserController struct {
	base.BaseApiController
	validate *validator.Validate
}

var validate = validator.New()
var useService service.UserService

func (c UserController) PostUser(users entity.Users) revel.Result {

	users.Uuid = uuid.Must(uuid.NewV4()).String()
	users.Password = useService.EncryptPw(users.Password)

	err := validate.Struct(users)
	if err != nil {
		return c.BadRequest("001")
	}

	userData := useService.GetUserByEmail(users.Email)
	if userData == nil {
		return c.InternalServer("002")
	}

	if !strings.EqualFold(userData.Email, "") {
		return c.UnprocessableEntity("003")
	}

	if useService.InsertUser(users) == nil {
		return c.InternalServer("004")
	}

	success := map[string]string {
		"message": "user creation succeeded.",
	}

	return c.RenderJSON(success)
}
