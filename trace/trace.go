package trace

import echo "github.com/labstack/echo/v4"

func ID(c echo.Context) string {
	return c.Request().Header.Get("X-Request-ID")
}
