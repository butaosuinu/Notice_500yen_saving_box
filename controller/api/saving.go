package api

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"notice_500yen_saving_box/model/saving"
)

// PostSaving
// POST /api/v1/saving
func PostSaving(c echo.Context) (err error) {
	savingTime := c.FormValue("time")
	fmt.Println(savingTime)
	saving.SaveSavingCount(savingTime)

	return c.String(http.StatusOK, "saving")
}
