package controllers

import (
	"net/http"
	"authentication-server/app/domains/model"
	"github.com/revel/revel"
)

type BaseApiController struct {
	*revel.Controller
}

// Bad Request Error
func (c BaseApiController) BadRequest(detail string) revel.Result {
	c.Response.Status = http.StatusBadRequest
	r := model.ErrorResponse {
		c.Response.Status,
		"Bad request.",
		detail,
	}
	return c.RenderJSON(r)
}

// Unauthorized Error
func (c BaseApiController) Unauthorized(detail string) revel.Result {
	c.Response.Status = http.StatusUnauthorized
	r := model.ErrorResponse {
		c.Response.Status,
		"Unauthorized.",
		detail,
	}
	return c.RenderJSON(r)
}

// Forbidden Error
func (c BaseApiController) Forbidden(detail string) revel.Result {
	c.Response.Status = http.StatusForbidden
	r := model.ErrorResponse {
		c.Response.Status,
		"Forbidden.",
		detail,
	}
	return c.RenderJSON(r)
}

// Not Found Error
func (c BaseApiController) NotFound(detail string) revel.Result {
	c.Response.Status = http.StatusNotFound
	r := model.ErrorResponse {
		c.Response.Status,
		"Not found.",
		detail,
	}
	return c.RenderJSON(r)
}

// Unprocessable Entity Error
func (c BaseApiController) UnprocessableEntity(detail string) revel.Result {
	c.Response.Status = http.StatusUnprocessableEntity
	r := model.ErrorResponse {
		c.Response.Status,
		"Unprocessable Entity.",
		detail,
	}
	return c.RenderJSON(r)
}

// Internal Server Error
func (c BaseApiController) InternalServer(detail string) revel.Result {
	c.Response.Status = http.StatusInternalServerError
	r := model.ErrorResponse {
		c.Response.Status,
		"Internal server error.",
		detail,
	}
	return c.RenderJSON(r)
}