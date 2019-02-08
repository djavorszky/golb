package main

import (
	"github.com/labstack/echo"
	"golb/server"
	"net/http"
)

// Route represents an http endpoint to be served by a router
type Route struct {
	Method  string
	Path    string
	Handler echo.HandlerFunc
	Mw      []echo.MiddlewareFunc
}

func getRoutes(blog server.Blog) []Route {
	return []Route{
		{Method: http.MethodGet, Path: "/", Handler: index},
		{Method: http.MethodGet, Path: "/v0/blog/entries", Handler: blog.Entries},
	}

}
