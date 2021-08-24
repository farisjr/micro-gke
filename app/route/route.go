package route

import (
	"app/config"
	"app/controller"

	"github.com/labstack/echo"
)

func SetRoute(config config.Config, e *echo.Echo) {
	if config.Mode == "" || config.Mode == "reader" {
		e.GET("news", controller.CreateGetNewsController(config))
	}
	if config.Mode == "" || config.Mode == "writer" {
		e.POST("news", controller.CreatePostNewsController(config))
	}
	if config.Mode == "gateway" {
		e.GET("news", controller.CreateGatewayGetNewsController(config))
		e.POST("news", controller.CreateGatewayPostNewsController(config))
	}
}
