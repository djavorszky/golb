package server

import (
	"github.com/labstack/echo"
	"net/http"
)

type Blog struct {
	logger  echo.Logger
	entries []Entry
}

type Entry struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func NewBlog(logger echo.Logger) Blog {
	return Blog{
		logger: logger,
		entries: []Entry{
			{"hello title", "hello content"},
			{"title two", "content two"},
		},
	}
}

func (b Blog) Entries(c echo.Context) error {
	return c.JSON(http.StatusOK, b.entries)
}
