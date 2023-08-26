package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {

	t := &Template{
		templates: template.Must(template.ParseGlob("../templates/**/*.html")),
	}

	e := echo.New()
	e.Renderer = t

	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	e.GET("/", HandleRoute)
	e.GET("/about", HandleRoute)

	e.GET("/api/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Some sort of info from an API")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
