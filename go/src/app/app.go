package main

import (
	"io"
	"github.com/labstack/echo"
	"html/template"
	"github.com/labstack/echo/middleware"
	"./controller"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	app := echo.New()
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	renderer := &Template{
		templates: template.Must(template.ParseGlob("static/template/*.html")),
	}

	app.Renderer = renderer

	appGroup := app.Group("/app")

	// トップ
	appGroup.GET("/index", controller.ShowIndexHtml)
	
	app.Logger.Fatal(app.Start(":9000"))
}

