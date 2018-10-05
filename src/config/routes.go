//-- Package Declaration -----------------------------------------------------------------------------------------------
package config

//-- Imports -----------------------------------------------------------------------------------------------------------
import (
	"github.com/JustonDavies/fhtagn/src/controllers/invites"
	"github.com/labstack/echo"
)

//-- Constants ---------------------------------------------------------------------------------------------------------
var (
	routes = []*route{

		//-- Reports ----------
		{`GET`, `/`, invites.Index, nil},
		{`POST`, `/attune`, invites.Read, nil},
		{`POST`, `/invoke`, invites.Update, nil},
	}
)

//-- Structs -----------------------------------------------------------------------------------------------------------
type route struct {
	Verb       string
	Path       string
	Handler    echo.HandlerFunc
	Middleware []echo.MiddlewareFunc
}

//-- Exported Functions ------------------------------------------------------------------------------------------------
func InjectRoutes(server *echo.Echo) {
	for _, route := range routes {
		server.Add(route.Verb, route.Path, route.Handler, route.Middleware...)
	}
}

//-- Internal Functions ------------------------------------------------------------------------------------------------
