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
		follows := server.Group("follows")
		{
			follows.POST("", func(c echo.Context) (err error) {
				return createFollowHandler(db, c)
			})
			follows.DELETE("", func(c echo.Context) (err error) {
				return deleteFollowHandler(db, c)
			})
			follows.GET("/follower/:followerId", func(c echo.Context) (err error) {
				return findFolloweesOfFollowerHandler(db, c)
			})
			follows.GET("/followee/:followeeId", func(c echo.Context) (err error) {
				return findFollowersOfFolloweeHandler(db, c)
			})
		}
	})
}

func createFollowHandler(db *ent.Client, c echo.Context) (err error) {
	createUserBody := new(dto.CreateFollowDTO)
	err = c.Bind(createUserBody)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	err = services.CreateFollow(db, *createUserBody, context.Background())
	if err != nil {
		if err == services.ErrFolloweeNotFound || err == services.ErrFollowerNotFound {
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": err.Error(),
			})
		}
		if err == services.ErrCannotFollowTwiceSameUser {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return c.NoContent(http.StatusCreated)
}

func deleteFollowHandler(db *ent.Client, c echo.Context) (err error) {
	deleteFollowBody := new(dto.DeleteFollowDTO)
	err = c.Bind(deleteFollowBody)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	err = services.DeleteFollow(db, *deleteFollowBody, context.Background())
	if err != nil {
		if err == services.ErrFollowNotFound {
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return c.NoContent(http.StatusNoContent)
}

func findFolloweesOfFollowerHandler(db *ent.Client, c echo.Context) (err error) {
	followerId := c.Param("followerId")

	id, err := uuid.Parse(followerId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	follows, err := services.FindFolloweesOfFollower(db, id, context.Background())
	if err != nil {
		if err == services.ErrFollowNotFound {
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": err.Error(),
			})
		}
	}

	return c.JSON(http.StatusOK, follows)
}

func findFollowersOfFolloweeHandler(db *ent.Client, c echo.Context) (err error) {
	followeeId := c.Param("followeeId")

	id, err := uuid.Parse(followeeId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	follows, err := services.FindFollowersOfFollowee(db, id, context.Background())
	if err != nil {
		if err == services.ErrFollowNotFound {
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": err.Error(),
			})
		}
	}

	return c.JSON(http.StatusOK, follows)
}
