package router

import (
	"test-grpc-project/pkg/handler"

	"github.com/gofiber/fiber/v2"
)

type Router struct {
	handler handler.IUserHandler
}

func NewRouter(handler handler.IUserHandler) *Router {
	return &Router{handler: handler}
}

func (r *Router) LoadRouter(app *fiber.App) {
	app.Post("/create", r.handler.Create)
	app.Delete("/delete", r.handler.Delete)
	app.Get("/:id", r.handler.Get)
}
