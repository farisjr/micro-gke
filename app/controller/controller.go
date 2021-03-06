package controller

import (
	"app/config"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func CreateGetNewsController(config config.Config) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"status":  "ok",
			"message": "success read news",
		})
	}
}

func CreatePostNewsController(config config.Config) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"status":  "ok",
			"message": "success post news",
		})
	}
}

func CreateGatewayGetNewsController(config config.Config) echo.HandlerFunc {
	return func(c echo.Context) error {
		resp, err := http.Get(fmt.Sprintf("%s/news", config.ReaderUrl))
		if err != nil {
			fmt.Println(err)
			return c.String(500, "reader error")
		}
		var data interface{}
		json.NewDecoder(resp.Body).Decode(&data)
		return c.JSON(200, data)
	}
}

func CreateGatewayPostNewsController(config config.Config) echo.HandlerFunc {
	return func(c echo.Context) error {
		resp, err := http.Post(fmt.Sprintf("%s/news", config.WriterUrl), "application/json", c.Request().Body)
		if err != nil {
			fmt.Println(err)
			return c.String(500, "reader error")
		}
		var data interface{}
		json.NewDecoder(resp.Body).Decode(&data)
		return c.JSON(200, data)
	}
}
