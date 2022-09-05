package controllers

import (
	"github.com/labstack/echo/v4"
	"log"
)

func StartServer(port string) *echo.Echo {
	e := echo.New()

	e.GET("test", func(c echo.Context) error {
		return c.String(200, "Hello world!")
	})

	log.Fatal(e.Start(":" + port))

	return e
}
