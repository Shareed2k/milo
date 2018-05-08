package internal

import (
	"github.com/labstack/echo"
	"github.com/milo/db/models"
	"github.com/ulule/paging"
	"net/http"
)

func servers(c *MiloContext) (err error) {
	db := c.GetMaster().GetDatabase()

	servers := []models.Server{}

	store, err := paging.NewGORMStore(db.DB, &servers)
	if err != nil {
		return err
	}

	options := paging.NewOptions()

	paginator, err := paging.NewOffsetPaginator(store, c.Request(), options)

	if err := paginator.Page(); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"paginator": paginator,
		"items":     servers,
	})
}