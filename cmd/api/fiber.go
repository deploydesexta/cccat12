package main

import (
	"context"
	"github.com/deploydesexta/cccat12/src/infrastructure/http"
	"github.com/deploydesexta/cccat12/src/infrastructure/jsonutil"
	"github.com/gofiber/fiber/v2"
)

type FiberRouterAdapter struct {
	Fiber *fiber.App
}

type FiberRequestAdapter struct {
	c *fiber.Ctx
}

func NewFiberRouterAdapter() *FiberRouterAdapter {
	f := fiber.New()
	return &FiberRouterAdapter{Fiber: f}
}

func (r *FiberRouterAdapter) Router(ctrl http.Router) {
	ctrl.Bind(r)
}

func (r *FiberRouterAdapter) Get(path string, handler http.HandlerFunc) {
	r.Fiber.Get(path, func(c *fiber.Ctx) error {
		return handler(&FiberRequestAdapter{c})
	})
}

func (r *FiberRouterAdapter) Post(path string, handler http.HandlerFunc) {
	r.Fiber.Post(path, func(c *fiber.Ctx) error {
		return handler(&FiberRequestAdapter{c})
	})
}

func (r *FiberRouterAdapter) Start(port string) error {
	return r.Fiber.Listen(port)
}

func (req *FiberRequestAdapter) Bind(i interface{}) error {
	return jsonutil.FromJson(req.c.Body(), i)
}

func (req *FiberRequestAdapter) JSON(code int, i interface{}) error {
	req.c.Status(code)
	return req.c.JSON(i)
}

func (req *FiberRequestAdapter) Param(name string) string {
	return req.c.Params(name)
}

func (req *FiberRequestAdapter) QueryParam(name string) string {
	return req.c.Query(name)
}

func (req *FiberRequestAdapter) String(code int, s string) error {
	req.c.Status(code)
	return req.c.SendString(s)
}

func (req *FiberRequestAdapter) Context() context.Context {
	return req.c.Context()
}
