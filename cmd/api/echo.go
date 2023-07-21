package main

import (
	"context"
	"github.com/deploydesexta/cccat12/src/infrastructure/http"
	echo "github.com/labstack/echo/v4"
)

type EchoRouterAdapter struct {
	*echo.Echo
}

type EchoRequestAdapter struct {
	c echo.Context
}

func NewEchoRouterAdapter() *EchoRouterAdapter {
	return &EchoRouterAdapter{Echo: echo.New()}
}

func (r *EchoRouterAdapter) Router(ctrl http.Router) {
	ctrl.Bind(r)
}

func (r *EchoRouterAdapter) Get(path string, handler http.HandlerFunc) {
	r.Echo.GET(path, func(c echo.Context) error {
		return handler(&EchoRequestAdapter{c})
	})
}

func (r *EchoRouterAdapter) Post(path string, handler http.HandlerFunc) {
	r.Echo.POST(path, func(c echo.Context) error {
		return handler(&EchoRequestAdapter{c})
	})
}

func (r *EchoRouterAdapter) Start(port string) error {
	return r.Echo.Start(port)
}

func (req *EchoRequestAdapter) Bind(i interface{}) error {
	return req.c.Bind(i)
}

func (req *EchoRequestAdapter) JSON(code int, i interface{}) error {
	return req.c.JSON(code, i)
}

func (req *EchoRequestAdapter) Param(name string) string {
	return req.c.Param(name)
}

func (req *EchoRequestAdapter) QueryParam(name string) string {
	return req.c.QueryParam(name)
}

func (req *EchoRequestAdapter) String(code int, s string) error {
	return req.c.String(code, s)
}

func (req *EchoRequestAdapter) Context() context.Context {
	return req.c.Request().Context()
}
