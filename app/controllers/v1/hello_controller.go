package v1

import (
	"github.com/revel/revel"
)

type HelloController struct {
	*revel.Controller
}

func (c HelloController) Index() revel.Result {
	hello := map[string]string{"key": "hello world"}
	return c.RenderJSON(hello)
}