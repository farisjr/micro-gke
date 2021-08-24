package main

import (
	"app/config"
	"app/route"
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	config := config.NewConfig()
	e := echo.New()
	e.Use(middleware.Logger())
	route.SetRoute(config, e)
	// serve
	port := config.GetPort()
	if err := e.Start(fmt.Sprintf(":%s", port)); err != nil {
		panic(err)
	}
}
