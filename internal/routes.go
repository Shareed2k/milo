package internal

import (
	"github.com/labstack/echo"
	"github.com/milo/db/models"
)

func NewRoutes(e *echo.Echo) {
	// api
	api := e.Group("/api")
	//api.Use(middleware.JWT([]byte("secret")))

	//api.POST("/login", restrictedHandler(login))
	api.POST("/login", simpleHandler(login))
}

func simpleHandler(fn func(*MiloContext) error) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return fn(ctx.(*MiloContext))
	}
}

func restrictedHandler(fn func(*MiloContext, *models.User) error) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		user := loadUser(ctx)

		if user == nil {
			return echo.ErrUnauthorized
		}

		return fn(ctx.(*MiloContext), user)
	}
}

func loadUser(c echo.Context) *models.User {
	user := new(models.User)

	/*t := c.Get("user").(*jwt.Token)
	claims := t.Claims.(jwt.MapClaims)

	id, err := internal.ParseID(fmt.Sprint(claims["user_id"].(float64)))

	if err != nil {
		return nil
	}

	if err := co.Users.Get(id, user); err != nil {
		return nil
	}*/

	return user
}
