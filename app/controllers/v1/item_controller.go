package v1

import (
	"github.com/revel/revel"
	"revel-api/app/domain/entity"
	"revel-api/app"
	"revel-api/app/controllers"
)

type ItemController struct {
	*revel.Controller
}

func (c ItemController) Get() revel.Result {
	var items []entity.Items

	app.Db.Order("id desc").Find(&items)

	response := controllers.Response{items}

	return c.RenderJSON(response)
}