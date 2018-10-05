//-- Package Declaration -----------------------------------------------------------------------------------------------
package config

//-- Imports -----------------------------------------------------------------------------------------------------------
import (
	"net/http"

	"github.com/labstack/echo"
)

//-- Constants ---------------------------------------------------------------------------------------------------------

//-- Structs -----------------------------------------------------------------------------------------------------------

//-- Exported Functions ------------------------------------------------------------------------------------------------
func InjectErrorHandler(server *echo.Echo) {
	server.HTTPErrorHandler = customHTTPErrorHandler
}

//-- Internal Functions ------------------------------------------------------------------------------------------------
func customHTTPErrorHandler(err error, context echo.Context) {
	context.Logger().Error(err)
	var message string
	if he, ok := err.(*echo.HTTPError); ok {
		message = he.Message.(string)
	} else {
		message = err.Error()
	}

	context.Render(http.StatusInternalServerError, `exception`, message)
}
