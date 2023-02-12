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
		likes := server.Group("likes")
		{
			likes.POST("", func(c echo.Context) error {
				return createLikeHandler(db, c)
			})
			likes.DELETE("", func(c echo.Context) error {
				return deleteLikeHandler(db, c)
			})
			likes.GET("/:postId", func(c echo.Context) error {
				return findLikesByPostIDHandler(db, c)
			})
		}
	})
}

func createLikeHandler(db *ent.Client, c echo.Context) (err error) {
	createLikeBody := new(dto.LikeDTO)
	err = c.Bind(createLikeBody)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": services.ErrBadRequest,
		})
	}

	err = services.CreateLike(db, *createLikeBody, context.Background())
	if err != nil {
		if err == services.ErrUserNotFound {
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": err.Error(),
			})
		}
		if err == services.ErrPostNotFound {
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": err.Error(),
			})
		}
		if err == services.ErrCannotLikeTwiceSamePost {
			return c.JSON(http.StatusConflict, echo.Map{
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return c.NoContent(http.StatusCreated)
}

func deleteLikeHandler(db *ent.Client, c echo.Context) (err error) {
	deleteLikeBody := new(dto.LikeDTO)
	err = c.Bind(deleteLikeBody)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": services.ErrBadRequest,
		})
	}

	err = services.DeleteLike(db, *deleteLikeBody, context.Background())
	if err != nil {
		if err == services.ErrLikeNotFound {
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": services.ErrInternalServerError.Error(),
		})
	}

	return c.NoContent(http.StatusNoContent)
}

func findLikesByPostIDHandler(db *ent.Client, c echo.Context) (err error) {
	postId := c.Param("postId")

	parsedPostId, err := uuid.Parse(postId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": services.ErrBadRequest,
		})
	}

	err, likes := services.FindLikesByPostId(db, parsedPostId, context.Background())
	if err != nil {
		if err == services.ErrLikeNotFound {
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"likes": likes,
	})
}
