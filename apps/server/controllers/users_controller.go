package controllers

import (
	"context"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"socialite/dto"
	"socialite/ent"
	"socialite/services"
)

func init() {
	ctx := context.Background()
	usersEndpoint := server.Group("users")
	{
		usersEndpoint.POST("", func(c echo.Context) error {
			return CreateUserHandler(c, database, ctx)
		})
		usersEndpoint.GET("", func(c echo.Context) error {
			return FindAllUsersHandler(c, database, ctx)
		})
		usersEndpoint.GET("/:id", func(c echo.Context) error {
			return FindUserByUUIDHandler(c, database, ctx)
		})
		/*		usersEndpoint.PUT("/:id", func(c echo.Context) error {
				return UpdateOneUserHandler(c, )
			})*/
		usersEndpoint.DELETE("/:id", func(c echo.Context) error {
			return DeleteOneUserHandler(c, database, ctx)
		})
	}
}

func CreateUserHandler(ctx echo.Context, db *ent.Client, c context.Context) error {
	var user dto.CreateUserDTO
	err := ctx.Bind(&user)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	err = services.CreateUser(db, user, c)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	return ctx.NoContent(http.StatusCreated)
}

func FindAllUsersHandler(ctx echo.Context, db *ent.Client, c context.Context) error {
	users, err := services.FindAllUsers(db, c)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}
	return ctx.JSON(http.StatusOK, users)
}

func FindUserByUUIDHandler(ctx echo.Context, db *ent.Client, c context.Context) error {
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

func DeleteOneUserHandler(ctx echo.Context, db *ent.Client, c context.Context) error {
	id := ctx.Param("id")
	parsedId, err := uuid.Parse(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	err = services.DeleteOneUser(db, parsedId, c)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.JSON(http.StatusNotFound, echo.Map{"message": err.Error()})
		}
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}
	return ctx.NoContent(http.StatusNoContent)
}
