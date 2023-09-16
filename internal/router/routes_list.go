package router

import (
	"golang.org/x/exp/slices"
)

type Route struct {
	Name         string
	Path         string
	TemplateName string
	RouteInfo    interface{}
}

var Routes []Route = []Route{
	{
		Name:         "Home",
		Path:         "/",
		TemplateName: "home",
	},
	{
		Name:         "About",
		Path:         "/about",
		TemplateName: "about",
	},
}

func GetRouteIndexByPath(path string) int {
	return slices.IndexFunc(Routes, func(r Route) bool {
		return r.Path == path
	})
}
