package internal

import (
	"github.com/labstack/echo"
	"github.com/milo/db/models"
	"github.com/ulule/paging"
	"net/http"
)

func indexRegion(c *MiloContext) (err error) {
	db := c.GetMaster().GetDatabase()

	regions := []models.Region{}

	store, err := paging.NewGORMStore(db.DB, &regions)
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

	if uuid != "" {
		user := new(models.User)
		if err := c.GetMaster().GetDatabase().First(user, "uuid = ?", uuid).Error; err == nil {
			return c.JSON(http.StatusOK, user)
		}
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