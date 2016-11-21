package api

import (
	"github.com/labstack/echo"
	"net/http"
	"notice_500yen_saving_box/model/saving"
)

type NowBalance struct {
	Balance string `json:"balace"`
}

// PostSaving
// POST /api/v1/saving
func PostSaving(c echo.Context) (err error) {
	savingTime := c.FormValue("time")
	saving.SaveSavingCount(savingTime)

	return c.String(http.StatusOK, "saving")
}

// GetSaving
// GET /api/v1/saving
func GetSaving(c echo.Context) (err error) {
	start := c.QueryParam("start")
	end := c.QueryParam("end")

	savingCollection := saving.FetchSavingCountCollectionByDate(start, end)

	return c.JSON(http.StatusOK, savingCollection)
}

// GetBalance
// GET /api/v1/balace
func GetBalance(c echo.Context) (err error) {
	nowBlance := NowBalance{
		saving.GetNowBalance().Balance,
	}

	balanceByte, err := json.Marshal(nowBlance)
	if err != nil {
		fmt.Println(err)
		return
	}

	return c.JSON(http.StatusOK, balanceByte)
}

// PostResetBalance
// POST /api/v1/reset_balance
func PostResetBalance(c echo.Context) (err error) {
	resetTime := c.FormValue("time")
	saving.ResetBalance(resetTime)

	return c.String(http.StatusOK, "reset")
}
