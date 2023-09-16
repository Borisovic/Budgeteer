package router

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func HandleRoute(c *fiber.Ctx) error {
	isHtmx := len(c.Request().Header.Peek("HX-Request")) > 0
	endpoint := string(c.Context().Request.RequestURI()[:])

	routeIndex := GetRouteIndexByPath(endpoint)

	if routeIndex == -1 {
		return handleNotFound(c, isHtmx)
	}
	return handleRoute(routeIndex, c, isHtmx)
}

func handleRoute(routeIndex int, c *fiber.Ctx, isHtmx bool) error {
	route := Routes[routeIndex]
	if !isHtmx {
		return c.Render("full", fiber.Map{
			"template_name": fmt.Sprintf("%v.django", route.TemplateName),
			"routes":        Routes,
		})
	}

	return c.Render(route.TemplateName, fiber.Map{
		"routeInfo": route.RouteInfo,
	})
}

func handleNotFound(c *fiber.Ctx, isHtmx bool) error {
	if !isHtmx {
		return c.Render("full", fiber.Map{
			"template_name": "404.django",
			"routes":        Routes,
		})
	}

	return c.Render("404", fiber.Map{
		"routes": Routes,
	})
}
