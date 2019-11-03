package templates

import (
	"github.com/labstack/echo"
	"html/template"
	"io"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func TemplateRegistry(e *echo.Echo) {
	t := &Template{
		templates: template.Must(template.ParseGlob("./templates/views/*.html")),
	}

	e.Renderer = t
}
