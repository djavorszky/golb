package server

import (
	"github.com/labstack/echo"
	"net/http"
)

type Blog struct {
}

func (b Blog) Entries() func(c echo.Context) error {
	return func(c echo.Context) error {

		mockEntries := []struct {
			Title   string `json:"title"`
			Content string `json:"content"`
		}{
			{"hello title", "hello content"},
			{"title two", "content two"},
		}

		return c.JSON(http.StatusOK, mockEntries)
	}
}
