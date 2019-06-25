package main

import (
	// "net/http"

	"github.com/labstack/echo"
	"github.com/swaggo/echo-swagger"
	_ "github.com/swaggo/echo-swagger/example/docs"

	"github.com/xpertks/gomissy/controllers"
	// "github.com/xpertks/gomissy/models"
)

func main() {
	e := echo.New()

	controllers.HomeController{}.Init(e.Group("/"))

	e.GET("/doc/*", echoSwagger.WrapHandler)

	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!")
	// })
	e.Logger.Fatal(e.Start(":1323"))
	println("ready")
}
