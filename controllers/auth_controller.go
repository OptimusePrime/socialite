package controllers

import (
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
	"socialite/dto"
	"socialite/ent"
	"socialite/services"
)

func init() {
	auth := server.Group("auth")
	{
		auth.GET("/login", func(c echo.Context) (err error) {
			return LoginUserHandler(c, database)
		})
		auth.GET("/refresh", func(c echo.Context) error {
			return RefreshUserAccessTokenHandler(c, database)
		})
	}
}

func LoginUserHandler(ctx echo.Context, db *ent.Client) (err error) {
	loginInfo := dto.LoginUserDTO{
		Email:    ctx.FormValue("email"),
		Password: ctx.FormValue("password"),
	}

	err, accessToken, refreshToken, isMatch := services.LoginUser(db, context.Background(), loginInfo)
	if !isMatch {
		return ctx.NoContent(http.StatusForbidden)
	}
	if err != nil {
		if ent.IsNotFound(err) {
			return ctx.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
}

func RefreshUserAccessTokenHandler(ctx echo.Context, db *ent.Client) (err error) {
	var refreshTokenBody dto.RefreshUserAccessTokenDTO
	err = ctx.Bind(&refreshTokenBody)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}
	refreshToken := refreshTokenBody.RefreshToken

	err, isValid, accessToken := services.RefreshUserAccessToken(db, context.Background(), refreshToken)
	if err != nil || !isValid {
		return ctx.JSON(http.StatusUnauthorized, echo.Map{
			"message": services.ErrInvalidRefreshToken.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"accessToken": accessToken,
	})
}
