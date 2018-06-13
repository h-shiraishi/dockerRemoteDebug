package controller

import (
	"github.com/labstack/echo"
	"net/http"
)

/**
トップ画面を表示する
@param c echo.Context
 */
func ShowIndexHtml(c echo.Context) error {
	return c.Render(http.StatusOK, "index", nil)
	
}
