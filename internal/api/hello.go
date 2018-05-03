package api

import (
	"github.com/labstack/echo"
	"net/http"
)

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}