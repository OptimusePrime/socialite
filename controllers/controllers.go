package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/meilisearch/meilisearch-go"
	"socialite/ent"
)

type IController interface {
	Init(*echo.Echo, *ent.Client, *meilisearch.Client)
}

var controllers []IController

type templateController struct {
	init func(server *echo.Echo, db *ent.Client, meili *meilisearch.Client)
}

func (tc templateController) Init(server *echo.Echo, db *ent.Client, meili *meilisearch.Client) {
	tc.init(server, db, meili)
}

func addController(init func(server *echo.Echo, db *ent.Client, meili *meilisearch.Client)) {
	temp := templateController{}
	temp.init = init
	controllers = append(controllers, temp)
}

func initControllers(server *echo.Echo, db *ent.Client, meili *meilisearch.Client) {
	for _, controller := range controllers {
		controller.Init(server, db, meili)
	}
}
