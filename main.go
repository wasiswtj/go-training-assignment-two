package main

import (
	"assignment-two/controller"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.POST("/order", controller.Create)
	e.PUT("/order", controller.Update)
	e.DELETE("/order/:id", controller.Delete)
	e.GET("/order/:id", controller.Get)

	e.Logger.Fatal(e.Start(":1323"))
}
