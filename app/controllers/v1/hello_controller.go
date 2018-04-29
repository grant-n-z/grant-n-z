package v1

import (
	"github.com/revel/revel"
	"revel-api/app/domains/entity"
)

type HelloController struct {
	*revel.Controller
}

func (c HelloController) Index() revel.Result {
	hello := entity.Hellos{
		Key: "Hello world",
	}

	return c.RenderJSON(hello)
}