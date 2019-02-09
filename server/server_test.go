package server

import (
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"golb/server/backend"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	mockEntriesJson = `[{"id":"one","author":"Joe","tags":["first tag","second tag"],"categories":["first category","second category"],"title":"first title","content":"first content","createDate":"2014-07-16T22:55:46+02:00","updateDate":"2014-07-16T22:55:46+02:00","publishDate":"2014-07-16T22:55:46+02:00"},{"id":"two","author":"Joe","tags":["first tag","third tag"],"categories":["first category","third category"],"title":"second title","content":"second content","createDate":"2014-07-17T22:55:46+02:00","updateDate":"2014-07-17T22:55:46+02:00","publishDate":"2014-07-17T22:55:46+02:00"}]
` // need to have a newline at the end because for some reason Echo adds one to the response.
)

func TestBlog_Entries(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/v0/blog/entries", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	blog := NewBlog(e.Logger, backend.NewMockStore())

	if assert.NoError(t, blog.Entries(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, mockEntriesJson, rec.Body.String())
	}
}
