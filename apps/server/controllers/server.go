package controllers

import (
	"context"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"log"
	"strings"
	"testing"
)

var server = echo.New()

func StartServer(port string, db *gorm.DB) *echo.Echo {
	usersEndpoint := server.Group("users")
	{
		usersEndpoint.POST("", func(c echo.Context) error {
			return CreateUserHandler(c, db)
		})
		usersEndpoint.GET("", func(c echo.Context) error {
			return FindAllUsersHandler(c, db)
		})
		usersEndpoint.GET("/:id", func(c echo.Context) error {
			return FindUserByUUIDHandler(c, db)
		})
		usersEndpoint.PUT("/:id", func(c echo.Context) error {
			return UpdateOneUserHandler(c, db)
		})
		usersEndpoint.DELETE("/:id", func(c echo.Context) error {
			return DeleteOneUserHandler(c, db)
		})
	}

	log.Fatal(server.Start(":" + port))

	return server
}

func InitTestServer(t *testing.T, port string, db *gorm.DB) func(string) string {
	go func() {
		app := StartServer(port, db)

		t.Cleanup(func() {
			err := app.Shutdown(context.Background())
			if err != nil {
				t.Error(err)
			}
		})
	}()

	return func(endpoint string) string {
		port := port
		endpoint = strings.Trim(endpoint, " ")
		if endpoint[0] == '/' {
			strings.Replace(endpoint, "/", "", 1)
		}
		return "http://localhost:" + port + "/" + endpoint
	}
}
