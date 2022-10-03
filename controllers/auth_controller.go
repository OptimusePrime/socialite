package controllers

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/meilisearch/meilisearch-go"
	"io"
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
	err = (&echo.DefaultBinder{}).BindBody(ctx, &loginInfo)
	b, _ := io.ReadAll(ctx.Request().Body)
	fmt.Printf("Info: %v", string(b))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	err, accessToken, refreshToken, isMatch := services.LoginUser(db, context.Background(), loginInfo)
	if !isMatch {
		return ctx.JSON(http.StatusForbidden, echo.Map{
			"message": err.Error(),
		})
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
		"accessToken":   accessToken,
		"refreshToken:": refreshToken,
	})
}

func refreshUserAccessTokenHandler(ctx echo.Context, db *ent.Client) (err error) {
	refreshTokenCookie, err := ctx.Cookie("socialite_refreshToken")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": services.ErrInvalidRefreshToken.Error(),
		})
	}
	refreshToken := refreshTokenCookie.Value

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
