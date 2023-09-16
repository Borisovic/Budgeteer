package router

import "github.com/gofiber/fiber/v2"

func HandleRoute(c *fiber.Ctx) error {
	htmxRequest := len(c.Request().Header.Peek("HX-Request")) > 0
	if htmxRequest {
		return c.Render("home", nil)
	}

	return c.Render("home-full", nil)
}
