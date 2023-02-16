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
		posts := server.Group("posts")
		{
			posts.POST("", func(c echo.Context) error {
				return createPostHandler(db, c)
			})
			posts.DELETE("", func(c echo.Context) error {
				return deletePostHandler(db, c)
			})
			posts.GET("", func(c echo.Context) error {
				return findPosts(db, c)
			})
			posts.GET("/one/:postId", func(c echo.Context) error {
				return findPostByID(db, c)
			})
			posts.GET("/many/:posterId", func(c echo.Context) error {
				return findPostsByPosterID(db, c)
			})
		}
	})
}

func createPostHandler(db *ent.Client, c echo.Context) (err error) {
	createPostBody := new(dto.CreatePostDTO)
	posterId, err := uuid.Parse(c.FormValue("poster"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}
	createPostBody.Poster = posterId
	createPostBody.Caption = c.FormValue("caption")
	createPostBody.Location = c.FormValue("location")
	image, err := c.FormFile("image")
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}
	createPostBody.Image = image

	err, postId := services.CreatePost(db, createPostBody, context.Background())
	if err != nil {
		if err == services.ErrUserNotFound {
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"id": postId,
	})
}

func deletePostHandler(db *ent.Client, c echo.Context) (err error) {
	deletePostBody := new(dto.DeletePostDTO)
	err = c.Bind(deletePostBody)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": services.ErrBadRequest.Error(),
		})
	}

	postId, err := uuid.Parse(deletePostBody.ID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": services.ErrBadRequest.Error(),
		})
	}
	err = services.DeletePost(db, postId, context.Background())
	if err != nil {
		if err == services.ErrPostNotFound {
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return c.NoContent(http.StatusOK)
}

func findPostByID(db *ent.Client, c echo.Context) (err error) {
	postId := c.Param("postId")

	parsedPostId, err := uuid.Parse(postId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	err, post := services.FindPostByID(db, parsedPostId, context.Background())
	if err != nil {
		if err == services.ErrPostNotFound {
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	// jsonPost, err := json.Marshal(post)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"posts": []any{post},
	})
}

func findPostsByPosterID(db *ent.Client, c echo.Context) (err error) {
	posterId := c.Param("posterId")

	parsedPosterId, err := uuid.Parse(posterId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	err, posts := services.FindPostsByPosterID(db, parsedPosterId, context.Background())
	if err != nil {
		if err == services.ErrUserNotFound {
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	// jsonPosts, err := json.Marshal(posts)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"posts": posts,
	})
}

func findPosts(db *ent.Client, c echo.Context) (err error) {
	findPostsDTO := new(dto.FindPostsDTO)
	err = c.Bind(findPostsDTO)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	err, posts := services.FindPosts(db, findPostsDTO.Limit, context.Background())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"posts": posts,
	})
}
