package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

type HomeController struct {
}

func (c HomeController) Init(g *echo.Group) {
	println("Home Controller - Call Get")

	g.GET("", c.Get)
}

func (HomeController) Get(c echo.Context) error {
	fmt.Println("Home Controller")

	data := {
		Name:  "Jon",
		Email: "jon@labstack.com"
	}
	println(data)
	return c.JSON(http.StatusOK, data)
}

// func (HomeController) Get(c echo.Context) error {
// 	return c.Render(http.StatusOK, "index", nil)
// }
