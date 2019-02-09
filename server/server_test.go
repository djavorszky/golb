package server

import (
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"golb/server/backend"
	"golb/server/model"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	mockEntries = []model.BlogEntry{
		{"hello title", "hello content"},
		{"title two", "content two"},
	}
	mockEntriesJson = `[{"title":"hello title","content":"hello content"},{"title":"title two","content":"content two"}]
`  // need to have a newline at the end because for some reason Echo adds one to the response.
)

func TestBlog_Entries(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/v0/blog/entries", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	blog := NewBlog(e.Logger, backend.MockStore{Entries: mockEntries})

	if assert.NoError(t, blog.Entries(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, mockEntriesJson, rec.Body.String())
	}
}
