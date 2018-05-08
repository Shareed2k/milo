package internal

import (
	"github.com/labstack/echo"
	"net/http"
)

func bootdata(c *MiloContext) (err error) {
	return c.JSON(http.StatusOK, echo.Map{
		"user": c.user,
	})
}