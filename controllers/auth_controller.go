package controllers

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/meilisearch/meilisearch-go"
	"net/http"
	"socialite/dto"
	"socialite/ent"
	"socialite/services"
)

func init() {
	addController(func(server *echo.Echo, db *ent.Client, _ *meilisearch.Client) {
		auth := server.Group("auth")
		{
			auth.POST("/login", func(c echo.Context) (err error) {
				return loginUserHandler(c, db)
			})
			auth.POST("/refresh", func(c echo.Context) error {
				return refreshUserAccessTokenHandler(c, db)
			})
		}
	})
}

func loginUserHandler(ctx echo.Context, db *ent.Client) (err error) {
	var loginInfo dto.LoginUserDTO
	err = ctx.Bind(&loginInfo)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	err, accessToken, refreshToken, isMatch := services.LoginUser(db, context.Background(), loginInfo)
	fmt.Printf("Match: %v\n", isMatch)
	fmt.Printf("Error: %v\n", err)
	if err != nil {
		if ent.IsNotFound(err) {
			return ctx.JSON(http.StatusBadRequest, echo.Map{
				"message": services.ErrInvalidEmail.Error(),
			})
		}
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	if !isMatch {
		return ctx.JSON(http.StatusForbidden, echo.Map{
			"message": services.ErrInvalidPassword.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
}

func refreshUserAccessTokenHandler(ctx echo.Context, db *ent.Client) (err error) {
	var refreshTokenDto dto.RefreshUserAccessTokenDTO
	err = ctx.Bind(&refreshTokenDto)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": services.ErrInvalidRefreshToken.Error(),
		})
	}
	refreshToken := refreshTokenDto.RefreshToken

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
