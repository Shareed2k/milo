package internal

import (
	"github.com/labstack/echo"
	"github.com/milo/db/models"
	"github.com/ulule/paging"
	"net/http"
)

func indexProvider(c *MiloContext) (err error) {
	db := c.GetMaster().GetDatabase()

	p := []models.Provider{}

	store, err := paging.NewGORMStore(db.DB, &p)
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
		"items":     p,
	})
}

func showProvider(c *MiloContext) (err error) {
	uuid := c.Param("uuid")

	if uuid != "" {
		p := new(models.Provider)
		if err := c.GetMaster().GetDatabase().First(p, "uuid = ?", uuid).Error; err == nil {
			return c.JSON(http.StatusOK, p)
		}
	}

	return c.JSON(http.StatusNotFound, echo.Map{
		"errors": "Entity not found",
	})
}

func storeProvider(c *MiloContext) (err error) {
	p := new(models.Provider)
	if err = c.Bind(p); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"errors": err.Error(),
		})
	}

	if err = c.Validate(p); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"errors": err.Error(),
		})
	}

	if err := c.GetMaster().GetDatabase().Create(p).Error; err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"errors": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, p)
}

func deleteProvider(c *MiloContext) (err error) {
	uuid := c.Param("uuid")

	if uuid != "" {
		if err := c.GetMaster().GetDatabase().Where("uuid = ?", uuid).Delete(models.Provider{}).Error; err != nil {
			return c.JSON(http.StatusUnprocessableEntity, echo.Map{
				"errors": err.Error(),
			})
		}
	}

	return c.JSON(http.StatusOK, echo.Map{
		"uuid": uuid,
	})
}

func updateProvider(c *MiloContext) (err error) {
	p := new(models.Provider)
	if err = c.Bind(p); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"errors": err.Error(),
		})
	}

	if err = c.Validate(p); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"errors": err.Error(),
		})
	}

	if err := c.GetMaster().GetDatabase().Save(p).Error; err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"errors": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, p)
}
