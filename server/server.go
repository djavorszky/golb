package server

import (
	"github.com/labstack/echo"
	"log"
	"net/http"
)

type Blog struct {
	BlogEntries []Entry
}

type Entry struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func NewBlog() Blog {
	return Blog{[]Entry{
		{"hello title", "hello content"},
		{"title two", "content two"},
	}}
}

func (b Blog) Entries() func(c echo.Context) error {
	return func(c echo.Context) error {
		log.Printf("blogentries: %v", b.BlogEntries)

		return c.JSON(http.StatusOK, b.BlogEntries)
	}
}
