package internal

import (
	"github.com/milo/db/models"
	"github.com/ulule/paging"
	"net/http"
)

type UsersResponse struct {
	*paging.OffsetPaginator
	Items []models.Server `json:"items"`
}

func servers(c *MiloContext, user *models.User) (err error) {
	db := c.GetMaster().GetDatabase()

	servers := []models.Server{}

	store, err := paging.NewGORMStore(db.DB, &servers)
	if err != nil {
		return err
	}

	options := paging.NewOptions()

	// Step 3: create a paginator instance and pass your store, your current HTTP
	// request and your options as arguments.
	paginator, err := paging.NewOffsetPaginator(store, c.Request(), options)

	if err := paginator.Page(); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &UsersResponse{
		OffsetPaginator: paginator,
		Items:           servers,
	})
}
