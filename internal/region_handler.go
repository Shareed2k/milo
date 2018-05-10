package internal

import (
	"github.com/labstack/echo"
	"github.com/milo/db/models"
	"github.com/ulule/paging"
	"net/http"
	"fmt"
)

func indexRegion(c *MiloContext) (err error) {
	db := c.GetMaster().GetDatabase().DB

	regions := []models.Region{}

	store, err := paging.NewGORMStore(db.Preload("DataCenters"), &regions)
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
		"items":     regions,
	})
}

func showRegion(c *MiloContext) (err error) {
	uuid := c.Param("uuid")

	r := new(models.Region)

	db := c.GetMaster().GetDatabase().DB

	if err := db.First(r, "uuid = ?", uuid).Error; err == nil {

		fmt.Println(r.DataCenters)

		return c.JSON(http.StatusOK, echo.Map{
			"region": r,
		})
	}

	return c.JSON(http.StatusNotFound, echo.Map{
		"errors": "Entity not found",
	})
}

func storeRegion(c *MiloContext) (err error) {
	r := new(models.Region)
	if err = c.Bind(r); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"errors": err.Error(),
		})
	}

	if err = c.Validate(r); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"errors": err.Error(),
		})
	}

	if err := c.GetMaster().GetDatabase().Create(r).Error; err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"errors": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, r)
}

func deleteRegion(c *MiloContext) (err error) {
	uuid := c.Param("uuid")

	if uuid != "" {
		if err := c.GetMaster().GetDatabase().Where("uuid = ?", uuid).Delete(models.Region{}).Error; err != nil {
			return c.JSON(http.StatusUnprocessableEntity, echo.Map{
				"errors": err.Error(),
			})
		}
	}

	return c.JSON(http.StatusOK, echo.Map{
		"uuid": uuid,
	})
}

func updateRegion(c *MiloContext) (err error) {
	r := new(models.Region)
	if err = c.Bind(r); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"errors": err.Error(),
		})
	}

	if err = c.Validate(r); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"errors": err.Error(),
		})
	}

	if err := c.GetMaster().GetDatabase().Save(r).Error; err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"errors": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, r)
}