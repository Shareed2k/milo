package internal

import (
	"github.com/labstack/echo"
	"github.com/milo/db/models"
	"github.com/ulule/paging"
	"net/http"
)

func indexServer(c *MiloContext) (err error) {
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

func serverInfo(c *MiloContext) (err error) {
	minionClient := NewMinionGrpcClient(c.GetSettings())
	defer minionClient.Close()

	minionClient.ConnectToServer("127.0.0.1", "8552")
	response, err := minionClient.GetStats(&StatsRequest{})

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"errors": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"info": response.Message,
		"procs": response.Processes,
	})
}