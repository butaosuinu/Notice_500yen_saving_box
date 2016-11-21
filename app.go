package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	"net/http"
	"notice_500yen_saving_box/controller/api"
	"notice_500yen_saving_box/controller/home"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	e.Static("/public", "public")

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world")
	})
	e.GET("/home", home.GetHomeIndex)

	apiv1 := e.Group("/api/v1")
	apiv1.GET("/saving", api.GetSaving)
	apiv1.POST("/saving", api.PostSaving)
	apiv1.GET("/balance", api.GetBalance)
	apiv1.POST("/balance/reset", api.PostResetBalance)

	apiv1.POST("/vibration", api.PostDetectedVibration)
	apiv1.POST("/open_box", api.PostOpenBox)

	e.Run(standard.New(":8080"))
}
