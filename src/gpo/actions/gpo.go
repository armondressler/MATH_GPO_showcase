package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

// GpoShow default implementation.
func GpoShow(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("gpo/show.html"))
}

// GpoList default implementation.
func GpoList(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("gpo/list.html"))
}

// GpoCreate default implementation.
func GpoCreate(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("gpo/create.html"))
}

// GpoDelete default implementation.
func GpoDelete(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("gpo/delete.html"))
}
