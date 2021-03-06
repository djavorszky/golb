package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"golb/server"
	"golb/server/backend"
	"net/http"
)

func main() {
	e := echo.New()

	addMiddlewares(e)

	blog := server.NewBlog(e.Logger, backend.NewMockStore())

	routes := getRoutes(blog)
	defineRoutes(e, routes)

	e.Logger.Fatal(e.Start(":1323"))
}

func addMiddlewares(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.RequestID())
	//	e.Use(middleware.Recover())
}

func defineRoutes(e *echo.Echo, routes []Route) {
	seen := make(map[string]bool)
	for _, r := range routes {
		if _, ok := seen[r.Path]; ok {
			panic("route already defined for path: " + r.Path)
		}

		e.Add(r.Method, r.Path, r.Handler, r.Mw...)

		seen[r.Path] = true
	}
}

func index(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
