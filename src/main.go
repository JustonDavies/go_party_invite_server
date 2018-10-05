//-- Package Declaration -----------------------------------------------------------------------------------------------
package main

//-- Imports -----------------------------------------------------------------------------------------------------------
import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/JustonDavies/fhtagn/src/config"
	"github.com/labstack/echo"
)

//-- Constants ---------------------------------------------------------------------------------------------------------

//-- Structs -----------------------------------------------------------------------------------------------------------

//-- Exported Functions ------------------------------------------------------------------------------------------------
func main() {
	//-- Initialize server ----------
	var server = echo.New()

	//-- Inject Global Middleware ----------
	//server.Pre(middleware.Logger())

	//-- Inject Routes ----------
	config.InjectRoutes(server)
	config.InjectRenderer(server)
	config.InjectErrorHandler(server)

	//-- Start server ----------
	start(server)

	//-- Listen of interrupts ----------
	stop(server)
}

//-- Internal Functions ------------------------------------------------------------------------------------------------
func start(server *echo.Echo) {
	server.Logger.Info(`Initializing: Starting...`)

	go func() {
		if err := server.Start(`:8080`); err != nil {
			server.Logger.Fatalf(`Initializing: Unrecoverable error: %s`, err)
		} else {
			server.Logger.Info(` Initializing: Success`)
		}
	}()
}

func stop(server *echo.Echo) {
	//-- Setup blocking channel and wait for interrupt ----------
	var interruptChannel = make(chan os.Signal)
	signal.Notify(interruptChannel, os.Interrupt)
	<-interruptChannel

	server.Logger.Info(`Shutdown: Interrupt detected, shutting down...`)

	//-- Setup shutdown context ----------
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//-- Attempt shutdown ----------
	if err := server.Shutdown(ctx); err != nil {
		server.Logger.Fatalf(`Shutdown: Unrecoverable error: %s`, err)
	}

	server.Logger.Info(`Shutdown: Success`)
}
