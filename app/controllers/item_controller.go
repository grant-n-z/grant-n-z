package controllers

import (
	"github.com/revel/revel"
	"revel-api/app/domain/entity"
)

type HelloController struct {
	*revel.Controller
}

func (c HelloController) Index() revel.Result {
	hello := entity.HelloEntity {
		Key: "Hello world",
	}

	return c.RenderJSON(hello)
}