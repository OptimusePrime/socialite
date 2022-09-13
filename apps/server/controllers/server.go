package controllers

import (
	"context"
	"github.com/labstack/echo/v4"
	"log"
	"socialite/ent"
	"strings"
	"testing"
)

var server = echo.New()
var database *ent.Client

func StartServer(port string, db *ent.Client) *echo.Echo {
	database = db

	log.Fatal(server.Start(":" + port))
	return server
}

func InitTestServer(t *testing.T, port string, db *ent.Client) func(string) string {
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
