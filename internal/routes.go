package internal

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/milo/db/models"
	"net/http"
)

func NewRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", map[string]interface{}{
			"csrf": c.Get("csrf"),
		})
	})
	e.POST("/login", routeHandler(login))

	// api
	api := e.Group("/api")
	config := middleware.JWTConfig{
		Claims:     &jwtClaims{},
		SigningKey: []byte("secret"),
	}
	api.Use(middleware.JWTWithConfig(config))
	api.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.(*MiloContext)
			user := &models.User{}

			// Check if authorize header is passed, and validate if user exist
			if token := c.Get("user"); token != nil {
				claims := token.(*jwt.Token).Claims.(*jwtClaims)

				if err := ctx.GetMaster().GetDatabase().First(user, claims.UserId).Error; err != nil {
					return echo.ErrUnauthorized
				}
			}

			ctx.user = user

			return next(ctx)
		}
	})
	api.GET("/bootdata", routeHandler(bootdata))

	// regions
	region := api.Group("/regions")
	region.GET("", routeHandler(indexRegion))
	region.GET("/:uuid", routeHandler(showRegion))
	region.POST("", routeHandler(storeRegion))
	region.DELETE("/:uuid", routeHandler(deleteRegion))
	region.PUT("", routeHandler(updateRegion))

	// datacenter
	dc := api.Group("/datacenters")
	dc.GET("", routeHandler(indexDataCenter))
	dc.GET("/:uuid", routeHandler(showDataCenter))
	dc.POST("", routeHandler(storeDataCenter))

	// servers
	server := api.Group("/servers")
	server.GET("", routeHandler(indexServer))
	server.GET("/info", routeHandler(serverInfo))
}

func routeHandler(fn func(*MiloContext) error) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return fn(ctx.(*MiloContext))
	}
}
