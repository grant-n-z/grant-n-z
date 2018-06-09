package v1

import "github.com/revel/revel"

type TokenController struct {
	*revel.Controller
}

func (c TokenController) PostUser() revel.Result {
	hello := map[string]string{"key": "hello world"}
	return c.RenderJSON(hello)
}