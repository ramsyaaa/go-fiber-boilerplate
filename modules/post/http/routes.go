package http

import (
	"github.com/gofiber/fiber/v2"
)

func PostRoutes(app *fiber.App, handler *PostHandler) {

	app.Get("/posts/:id", handler.Show)
	app.Post("/posts", handler.Store)
	app.Put("/posts/:id", handler.Update)
	app.Delete("/posts/:id", handler.Delete)
}
