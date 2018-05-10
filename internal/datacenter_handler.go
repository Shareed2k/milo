package internal

import (
	"github.com/labstack/echo"
	"github.com/milo/db/models"
	"github.com/ulule/paging"
	"net/http"
)

func indexDataCenter(c *MiloContext) (err error) {
	db := c.GetMaster().GetDatabase().DB

	db.Set("gorm:auto_preload", true)

	dc := new([]models.DataCenter)

	store, err := paging.NewGORMStore(db.Preload("Provider").Preload("Region"), dc)
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
		"items":     dc,
	})
}

func showDataCenter(c *MiloContext) (err error) {
	uuid := c.Param("uuid")

	if uuid != "" {
		db := c.GetMaster().GetDatabase().DB
		db.Preload("Provider").Preload("Region")

		dc := new(models.DataCenter)
		if err := db.First(dc, "uuid = ?", uuid).Error; err == nil {
			return c.JSON(http.StatusOK, dc)
		}
	}

	return c.JSON(http.StatusNotFound, echo.Map{
		"errors": "Entity not found",
	})
}

func storeDataCenter(c *MiloContext) (err error) {
	dc := new(models.DataCenter)
	if err = c.Bind(dc); err != nil {
		return
	}

	if err = c.Validate(dc); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	db := c.GetMaster().GetDatabase().DB

	if err := db.Create(dc).Error; err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"errors": err.Error(),
		})
	}

	db.Preload("Provider").Preload("Region").First(dc)

	return c.JSON(http.StatusCreated, dc)
}


func deleteDataCenter(c *MiloContext) (err error) {
	uuid := c.Param("uuid")

	if uuid != "" {
		if err := c.GetMaster().GetDatabase().Where("uuid = ?", uuid).Delete(models.DataCenter{}).Error; err != nil {
			return c.JSON(http.StatusUnprocessableEntity, echo.Map{
				"errors": err.Error(),
			})
		}
	}

	return c.JSON(http.StatusOK, echo.Map{
		"uuid": uuid,
	})
}

func updateDataCenter(c *MiloContext) (err error) {
	dc := new(models.DataCenter)
	if err = c.Bind(dc); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"errors": err.Error(),
		})
	}

	if err = c.Validate(dc); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"errors": err.Error(),
		})
	}

	db := c.GetMaster().GetDatabase().DB
	db.Preload("Provider").Preload("Region")

	if err := db.Save(dc).Error; err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"errors": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dc)
}