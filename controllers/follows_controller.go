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
			follows.GET("/userFollows/:userId", func(c echo.Context) (err error) {
				return findWhoUserFollows(db, c)
			})
			follows.GET("/followersOfUser/:userId", func(c echo.Context) (err error) {
				return findFollowersOfUser(db, c)
			})
			follows.PUT("", func(c echo.Context) error {
				return findFollow(db, c)
			})
		}
	})
}

func createFollowHandler(db *ent.Client, c echo.Context) (err error) {
	createFollowBody := new(dto.FollowDTO)
	err = c.Bind(createFollowBody)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": services.ErrBadRequest,
		})
	}

	err = services.CreateFollow(db, *createFollowBody, context.Background())
	if err != nil {
		if err == services.ErrFolloweeNotFound || err == services.ErrFollowerNotFound {
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": err.Error(),
			})
		}
		if err == services.ErrCannotFollowTwiceSameUser {
			return c.JSON(http.StatusConflict, echo.Map{
				"message": err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": services.ErrInternalServerError,
		})
	}

	return c.NoContent(http.StatusCreated)
}

func deleteFollowHandler(db *ent.Client, c echo.Context) (err error) {
	deleteFollowBody := new(dto.FollowDTO)
	err = c.Bind(deleteFollowBody)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": services.ErrBadRequest,
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
			"message": services.ErrInternalServerError.Error(),
		})
	}

	return c.NoContent(http.StatusNoContent)
}

func findFollow(db *ent.Client, c echo.Context) (err error) {
	followBody := new(dto.FollowDTO)
	err = c.Bind(followBody)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": services.ErrBadRequest.Error(),
		})
	}

	isFollowing, err := services.FindFollow(db, *followBody, context.Background())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": services.ErrInternalServerError.Error(),
		})
	}
	if !isFollowing {
		return c.JSON(http.StatusNotFound, echo.Map{
			"message": services.ErrFollowNotFound.Error(),
		})
	}

	return c.NoContent(http.StatusOK)
}

func findWhoUserFollows(db *ent.Client, c echo.Context) (err error) {
	followerId := c.Param("userId")

	id, err := uuid.Parse(followerId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	followees, err := services.FindWhoUserFollows(db, id, context.Background())
	if err != nil {
		if err == services.ErrFollowNotFound {
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": err.Error(),
			})
		}
	}

	return c.JSON(http.StatusOK, followees)
}

func findFollowersOfUser(db *ent.Client, c echo.Context) (err error) {
	followeeId := c.Param("userId")

	id, err := uuid.Parse(followeeId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": services.ErrInvalidId.Error(),
		})
	}

	followers, err := services.FindFollowersOfUser(db, id, context.Background())
	if err != nil {
		if err == services.ErrFollowNotFound {
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": err.Error(),
			})
		}
	}

	return c.JSON(http.StatusOK, followers)
}
