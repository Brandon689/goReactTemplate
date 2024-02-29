package api

import (
	"github.com/Brandon689/goReactTemplate/types"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetByID(c echo.Context) error {
	//id := c.Param("id")
	var n types.Example
	//if err := db.First(&todo, id).Error; err != nil {
	//	return c.NoContent(http.StatusNotFound)
	//}
	return c.JSON(http.StatusOK, n)
}
