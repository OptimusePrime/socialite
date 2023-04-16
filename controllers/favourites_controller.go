package controllers

import (
	"context"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/meilisearch/meilisearch-go"
	"net/http"
	"socialite/dto"
	"socialite/ent"
	"socialite/services"
)

func init() {
	addController(func(server *echo.Echo, db *ent.Client, meili *meilisearch.Client) {
		favourites := server.Group("favourites")
		{
			favourites.POST("", func(c echo.Context) error {
				return createFavouriteHandler(c, db)
			})
			favourites.DELETE("", func(c echo.Context) error {
				return deleteFavouriteHandler(c, db)
			})
			favourites.GET("/:userId", func(c echo.Context) error {
				return findFavouriteByUserIdHandler(c, db)
			})
		}
	})
}

func createFavouriteHandler(ctx echo.Context, db *ent.Client) (err error) {
	var createFavouriteDto dto.CreateFavouriteDTO
	err = ctx.Bind(&createFavouriteDto)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	err = services.CreateFavourite(db, createFavouriteDto, context.Background())
	if err != nil {
		if err == services.ErrPostNotFound {
			return ctx.JSON(http.StatusNotFound, echo.Map{
				"message": err.Error(),
			})
		}
		if err == services.ErrFavouriteAlreadyExists {
			return ctx.JSON(http.StatusConflict, echo.Map{
				"message": err.Error(),
			})
		}

		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return ctx.NoContent(http.StatusCreated)
}

func deleteFavouriteHandler(ctx echo.Context, db *ent.Client) (err error) {
	var deleteFavouriteDto dto.DeleteFavouriteDTO
	err = ctx.Bind(&deleteFavouriteDto)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	err = services.DeleteFavourite(db, deleteFavouriteDto, context.Background())
	if err != nil {
		if err == services.ErrFavouriteNotFound {
			return ctx.JSON(http.StatusNotFound, echo.Map{
				"message": err.Error(),
			})
		}

		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return ctx.NoContent(http.StatusOK)
}

func findFavouriteByUserIdHandler(ctx echo.Context, db *ent.Client) (err error) {
	userId, err := uuid.Parse(ctx.Param("userId"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	err, favourites := services.FindFavouriteByUserId(db, userId, context.Background())
	if err != nil {
		if err == services.ErrUserNotFound {
			return ctx.JSON(http.StatusNotFound, echo.Map{
				"message": err.Error(),
			})
		}

		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, favourites)
}
