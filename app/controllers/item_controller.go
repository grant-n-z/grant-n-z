package controllers

import (
	"github.com/revel/revel"
	"revel-api/app/domain/entity"
	"revel-api/app"
)

type JsonResponse struct {
	Response interface{} `json:"items"`
}

type ItemController struct {
	*revel.Controller
}

func (c ItemController) Get() revel.Result {
	items := []entity.Items{}

	app.Db.Order("id desc").Find(&items)

	response := JsonResponse{}
	response.Response = items

	return c.RenderJSON(response)
}