package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/meilisearch/meilisearch-go"
	"log"
	"socialite/ent"
)

func StartServer(port string, db *ent.Client, meili *meilisearch.Client) *echo.Echo {
	server := echo.New()
	server.Use(middleware.Logger())
	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
	}))
	initControllers(server, db, meili)

	log.Fatal(server.Start(":" + port))
	return server
}

/*func InitTestServer(t *testing.T, port string, db *ent.Client) func(string) string {
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
}*/
