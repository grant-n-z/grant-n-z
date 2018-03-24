package controllers

import "github.com/revel/revel"
import "revel-api/app/domain/model"

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.RenderJSON(model.UserModel{})
}

