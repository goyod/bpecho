package main

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"

	"github.com/goyod/pbecho/config"
	"github.com/goyod/pbecho/feature"
	"github.com/goyod/pbecho/trace"
	"github.com/goyod/pbecho/xservice"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", handler)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

var conf = config.New(viper.New())

// Handler
func handler(c echo.Context) error {
	service := xservice.New(conf.XClient(), conf.XRequest(), c.Request().Header.Get("X-Request-ID"))
	res, _ := feature.Feature(service, conf.XHandler(), trace.ID(c))
	return c.JSON(http.StatusOK, res)
}
