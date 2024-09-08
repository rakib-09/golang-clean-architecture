package middlewares

import (
	"github.com/rakib-09/golang-clean-architecture/config"
	"github.com/rakib-09/golang-clean-architecture/utils/msgutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CustomAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			headers := c.Request().Header
			appKey := headers.Get(config.App().AppKeyHeader)

			if appKey == config.App().AppKey {
				return next(c)
			}

			return c.JSON(http.StatusForbidden, msgutil.AccessForbiddenMsg())
		}
	}
}
