package v1

import (
	"authentication-server/app/domains/entity"
	"authentication-server/app/controllers"
	"github.com/satori/go.uuid"
	"github.com/revel/revel"
	"gopkg.in/go-playground/validator.v9"
	"authentication-server/app/domains/service"
)

type UserController struct {
	controllers.BaseApiController
	validate *validator.Validate
}

var validate = validator.New()

func (c UserController) PostUser(users entity.Users) revel.Result {

	users.Uuid = uuid.Must(uuid.NewV4()).String()

	err := validate.Struct(users)
	if err != nil {
		return c.BadRequest("001")
	}

	var useService service.UserService
	if useService.GetUserByEmail(users.Email).Response == nil {
		return c.NotFound("002")
	}
	
	if useService.InsertUser(users) == false {
		return c.InternalServer("003")
	}

	return c.RenderJSON(users)
}
