package main

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/xpertks/gomissy/controllers"
	"github.com/xpertks/gomissy/models"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
