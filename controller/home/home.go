package home

import (
	"github.com/flosch/pongo2"
	"github.com/labstack/echo"
	"net/http"
	"notice_500yen_saving_box/model/saving"
)

// GetHomeIndex
// GET /home
func GetHomeIndex(c echo.Context) (err error) {
	tpl, err := pongo2.FromFile("view/home.html")
	if err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	nowBalance := saving.GetNowBalance().Balance
	latestOpenBoxRecord := saving.FindLatestOpenBoxRecord()

	out, err := tpl.Execute(pongo2.Context{
		"balance":         nowBalance,
		"open_box_time":   latestOpenBoxRecord.Time,
		"open_box_taking": latestOpenBoxRecord.Taking,
	})
	if err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	return c.HTML(http.StatusOK, out)
}
