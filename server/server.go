package server

import (
	"github.com/labstack/echo"
	"golb/server/backend"
	"net/http"
)

type Blog struct {
	logger echo.Logger
	be     backend.Store
}

func NewBlog(logger echo.Logger, be backend.Store) Blog {
	return Blog{
		logger: logger,
		be:     be,
	}
}

func (b Blog) Entries(c echo.Context) error {
	return c.JSON(http.StatusOK, b.be.GetAll())
}
