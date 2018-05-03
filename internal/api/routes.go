package api

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewRoutes(e *echo.Echo){
	// api
	api := e.Group("/api")
	api.Use(middleware.JWT([]byte("secret")))

	api.GET("/test", hello)

}
