package main

import (
	"Notice_500yen_saving_box/controller/api"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	e.Static("/public", "public")

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world")
	})

	apiv1 := e.Group("/api/v1")
	apiv1.POST("/saving", api.postSaving)

	e.Run(standard.New(":8080"))
}
