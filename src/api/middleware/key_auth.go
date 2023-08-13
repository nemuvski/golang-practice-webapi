package api_middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func KeyAuth(validToken string) echo.MiddlewareFunc {
	return middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		KeyLookup: "header:" + echo.HeaderAuthorization,
		AuthScheme: "Bearer",
		Validator: func(key string, c echo.Context) (bool, error) {
			if key == validToken {
				return true, nil
			}
			return false, nil
		},
		ErrorHandler: func(err error, c echo.Context) error {
			return echo.ErrForbidden
		},
	})
}
