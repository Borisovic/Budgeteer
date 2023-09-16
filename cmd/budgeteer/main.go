package main

import (
	"io"

	"internal/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django/v3"
)

type Views interface {
	Load() error
	Render(io.Writer, string, interface{}, ...string) error
}

func main() {
	engine := django.New("../../web/templates", ".django")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.All("/api/*", func(c *fiber.Ctx) error {
		return c.SendString("This is an API endpoint")
	})

	app.Get("/*", router.HandleRoute)

	app.Static("/", "../../web/static")
	app.Listen("localhost:3000")
}
