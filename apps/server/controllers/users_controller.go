package controllers

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"socialite/models"
	"socialite/services"
)

func init() {

}

func CreateUserHandler(ctx echo.Context, db *gorm.DB) error {
	var user models.User
	err := ctx.Bind(&user)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	err = services.CreateUser(db, user)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	return ctx.NoContent(http.StatusCreated)
}

func FindAllUsersHandler(ctx echo.Context, db *gorm.DB) error {
	users, err := services.FindAllUsers(db)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}
	return ctx.JSON(http.StatusOK, users)
}

func FindUserByUUIDHandler(ctx echo.Context, db *gorm.DB) error {
	id := ctx.Param("id")

	parsedId, err := uuid.Parse(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	user, err := services.FindUserByUUID(db, parsedId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.NoContent(http.StatusNotFound)
		}
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}
	return ctx.JSON(http.StatusOK, user)
}

func UpdateOneUserHandler(ctx echo.Context, db *gorm.DB) error {
	var user models.User

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
}

func DeleteOneUserHandler(ctx echo.Context, db *gorm.DB) error {
	id := ctx.Param("id")
	parsedId, err := uuid.Parse(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	err = services.DeleteOneUser(db, parsedId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.JSON(http.StatusNotFound, echo.Map{"message": err.Error()})
		}
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}
	return ctx.NoContent(http.StatusNoContent)
}
