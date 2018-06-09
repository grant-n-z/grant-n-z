package v1

import (
	"authentication-server/app/domains/entity"
	"authentication-server/app/controllers"
	"authentication-server/app/domains/service"
	"github.com/satori/go.uuid"
	"github.com/revel/revel"
)

type UserController struct {
	controllers.BaseApiController
}

func (c UserController) PostUser() revel.Result {

	users := entity.Users {
		Uuid: uuid.Must(uuid.NewV4()).String(),
	}

	if err := c.Params.BindJSON(&users); err != nil {
		return c.BadRequest("001")
	}

	// TODO: ユーザ存在確認

	var useService service.UserService
	if useService.InsertUser(users) == false {
		return c.InternalServer("002")
	}

	return c.RenderJSON(users)
}