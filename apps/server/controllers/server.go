package controllers

import (
	"github.com/labstack/echo/v4"
	"log"
)

func StartServer(port string) *echo.Echo {
	e := echo.New()

	log.Fatal(e.Start(":" + port))

	return e
}
