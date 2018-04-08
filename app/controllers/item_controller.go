package controllers

import (
	"github.com/revel/revel"
	"revel-api/app/domain/entity"
	"revel-api/app"
)

type ItemController struct {
	*revel.Controller
	app.GormController
}

func (c ItemController) Get() revel.Result {
	var user [] entity.ItemEntity

	c.Txn.Find(user)
	return c.RenderJSON(user)
}