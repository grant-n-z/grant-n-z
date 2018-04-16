package controllers

import (
	"github.com/revel/revel"
	"net/http"
)

type ApiController struct {
	*revel.Controller
}

// Response error JSON structure
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Response ok JSON structure
type Response struct {
	Results interface{} `json:"results"`
}

// Bad error structure
func (c *ApiController) HandleBadRequestError(s string) revel.Result {
	c.Response.Status = http.StatusBadRequest
	r := ErrorResponse{c.Response.Status, s}
	return c.RenderJSON(r)
}

// Not found structure
func (c *ApiController) HandleNotFoundError(s string) revel.Result {
	c.Response.Status = http.StatusNotFound
	r := ErrorResponse{c.Response.Status, s}
	return c.RenderJSON(r)
}

// Internal server error structure
func (c *ApiController) HandleInternalServerError(s string) revel.Result {
	c.Response.Status = http.StatusInternalServerError
	r := ErrorResponse{c.Response.Status, s}
	return c.RenderJSON(r)
}