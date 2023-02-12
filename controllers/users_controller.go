package controllers

import (
	"context"
	"fmt"
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
		ctx := context.Background()
		usersEndpoint := server.Group("users")
		{
			usersEndpoint.POST("", func(c echo.Context) error {
				return createUserHandler(c, db, meili, ctx)
			})
			usersEndpoint.GET("", func(c echo.Context) error {
				return findAllUsersHandler(c, db, ctx)
			})
			usersEndpoint.GET("/:id", func(c echo.Context) error {
				return findUserByUUIDHandler(c, db, ctx)
			})
			usersEndpoint.DELETE("", func(c echo.Context) error {
				return deleteOneUserHandler(c, db, meili, ctx)
			})
			usersEndpoint.PATCH("", func(c echo.Context) error {
				return updateOneUserHandler(c, db, meili, ctx)
			})
		}
	})
}

func createUserHandler(ctx echo.Context, db *ent.Client, meili *meilisearch.Client, c context.Context) (err error) {
	var user dto.CreateUserDTO
	err = ctx.Bind(&user)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	fmt.Printf("User received: %s", user)

	err = services.CreateUser(db, meili, user, c)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}
	return ctx.NoContent(http.StatusCreated)
}

func findAllUsersHandler(ctx echo.Context, db *ent.Client, c context.Context) (err error) {
	users, err := services.FindAllUsers(db, c)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}
	return ctx.JSON(http.StatusOK, users)
}

func findUserByUUIDHandler(ctx echo.Context, db *ent.Client, c context.Context) (err error) {
	id := ctx.Param("id")

	parsedId, err := uuid.Parse(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	user, err := services.FindUserByUUID(db, parsedId, c)
	if err != nil {
		if ent.IsNotFound(err) {
			return ctx.NoContent(http.StatusNotFound)
		}
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}
	return ctx.JSON(http.StatusOK, user)
}

/*func UpdateOneUserHandler(ctx echo.Context, db *ent.Client) error {
	id := ctx.Param("id")
	parsedId, err := uuid.Parse(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	err = ctx.Bind(&user)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	err = services.UpdateOneUser(db, user, parsedId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.NoContent(http.StatusNotFound)
		}
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}
	return ctx.NoContent(http.StatusNoContent)
}*/

func deleteOneUserHandler(ctx echo.Context, db *ent.Client, meili *meilisearch.Client, c context.Context) (err error) {
	err, accessToken := services.GetBearerToken(ctx)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": services.ErrInvalidAccessToken.Error(),
		})
	}

	err = services.DeleteOneUser(db, meili, c, accessToken)
	if err != nil {
		if err == services.ErrInvalidAccessToken {
			return ctx.JSON(http.StatusUnauthorized, echo.Map{
				"message": err.Error(),
			})
		}
		if ent.IsNotFound(err) {
			return ctx.JSON(http.StatusNotFound, echo.Map{
				"message": err.Error(),
			})
		}
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return ctx.NoContent(http.StatusNoContent)
}

func updateOneUserHandler(ctx echo.Context, db *ent.Client, meili *meilisearch.Client, c context.Context) (err error) {
	updateUserBody := new(dto.UpdateUserDTO)
	err = ctx.Bind(updateUserBody)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	err = services.UpdateUser(db, meili, c, *updateUserBody)
	if err != nil {
		return err
	}

	return nil
}
