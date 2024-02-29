package api

import (
	"github.com/Brandon689/goReactTemplate/types"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Get(c echo.Context) error {
	var n types.User
	return c.JSON(http.StatusOK, n)
}
