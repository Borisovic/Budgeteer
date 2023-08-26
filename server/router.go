package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/exp/slices"
)

type Route struct {
	PageTitle string
	PageRoute string
}

type RouteData struct {
	Routes      []Route
	ActiveRoute Route
}

func HandleRoute(c echo.Context) error {

	routes := []Route{
		{PageTitle: "Home", PageRoute: "/"},
		{PageTitle: "About", PageRoute: "/about"},
	}

	routeIndex := slices.IndexFunc(routes, func(r Route) bool { return r.PageRoute == c.Request().RequestURI })
	if routeIndex == -1 {
		return c.Render(http.StatusNotFound, "not-found", nil)
	}

	return c.Render(http.StatusOK, "index.html", RouteData{Routes: routes, ActiveRoute: routes[routeIndex]})
}
