package server

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/witchakornb/YouTube-Payment/config"
	"github.com/witchakornb/YouTube-Payment/database"
)

type echoServer struct {
	server *echo.Echo
	db database.Database
	config *config.Config
}

func NewEchoServer(config *config.Config, db database.Database) Server {
	echoServerApp := echo.New()
	echoServerApp.Logger.SetLevel(log.DEBUG)

	return &echoServer{
		server: echoServerApp,
		db: db,
		config: config,
	}
}

func (e *echoServer) Start() {
	e.server.Use(middleware.Logger())
	e.server.Use(middleware.Recover())

	e.server.GET("api/v1/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})


	e.server.Logger.Fatal(e.server.Start(fmt.Sprintf(":%d", e.config.Server.Port)))
}

func (e *echoServer) initializePaymentHttpHandler() {
	// Initialize your payment http handler here

}