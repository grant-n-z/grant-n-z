package base

import (
	"net/http"
	"github.com/revel/revel"
	"github.com/tomo0111/grant-n-z/app/domains/model"
)

type BaseApiController struct {
	*revel.Controller
}

// Bad Request Error
func (c BaseApiController) BadRequest(detail string) revel.Result {
	c.Response.Status = http.StatusBadRequest
	r := model.ErrorResponse {
		Code:    c.Response.Status,
		Message: "Bad request.",
		Detail:  detail,
	}
	return c.RenderJSON(r)
}

// Unauthorized Error
func (c BaseApiController) Unauthorized(detail string) revel.Result {
	c.Response.Status = http.StatusUnauthorized
	r := model.ErrorResponse {
		Code:    c.Response.Status,
		Message: "Unauthorized.",
		Detail:  detail,
	}
	return c.RenderJSON(r)
}

// Forbidden Error
func (c BaseApiController) Forbidden(detail string) revel.Result {
	c.Response.Status = http.StatusForbidden
	r := model.ErrorResponse {
		Code:    c.Response.Status,
		Message: "Forbidden.",
		Detail:  detail,
	}
	return c.RenderJSON(r)
}

// Not Found Error
func (c BaseApiController) NotFound(detail string) revel.Result {
	c.Response.Status = http.StatusNotFound
	r := model.ErrorResponse {
		Code:    c.Response.Status,
		Message: "Not found.",
		Detail:  detail,
	}
	return c.RenderJSON(r)
}

// Unprocessable Entity Error
func (c BaseApiController) UnprocessableEntity(detail string) revel.Result {
	c.Response.Status = http.StatusUnprocessableEntity
	r := model.ErrorResponse {
		Code:    c.Response.Status,
		Message: "Unprocessable Entity.",
		Detail:  detail,
	}
	return c.RenderJSON(r)
}

// Internal Server Error
func (c BaseApiController) InternalServer(detail string) revel.Result {
	c.Response.Status = http.StatusInternalServerError
	r := model.ErrorResponse {
		Code:    c.Response.Status,
		Message: "Internal server error.",
		Detail:  detail,
	}
	return c.RenderJSON(r)
}