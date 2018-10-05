//-- Package Declaration -----------------------------------------------------------------------------------------------
package config

//-- Imports -----------------------------------------------------------------------------------------------------------
import (
	"html/template"
	"io"

	"github.com/labstack/echo"
)

//-- Constants ---------------------------------------------------------------------------------------------------------

//-- Structs -----------------------------------------------------------------------------------------------------------
type renderer struct {
	templates *template.Template
}

//-- Exported Functions ------------------------------------------------------------------------------------------------
func InjectRenderer(server *echo.Echo) {
	server.Renderer = &renderer{
		templates: template.Must(template.ParseGlob("src/controllers/*/*.gohtml")),
	}

	server.Static(`/assets`, `assets`)
}

func (engine *renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return engine.templates.ExecuteTemplate(w, name, data)
}

//-- Internal Functions ------------------------------------------------------------------------------------------------
