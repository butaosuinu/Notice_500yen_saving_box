package api

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"notice_500yen_saving_box/model/saving"
)

// PostDetectedVibration
// POST /api/v1/vibration
func PostDetectedVibration(c echo.Context) (err error) {
	time := c.FormValue("time")
	fmt.Println(time)

	return c.String(http.StatusOK, "Detected vibration")
}

// PostOpenBox
// POST /api/v1/open_box
func PostOpenBox(c echo.Context) (err error) {
	openTime := c.FormValue("time")
	saving.SaveOpenBox(openTime)

	return c.String(http.StatusOK, "open box")
}
