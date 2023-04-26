package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

// LoginShow default implementation.
func LoginShow(c buffalo.Context) error {
	c.Set("render_without_navbar", "true")
	return c.Render(http.StatusOK, r.HTML("login/show.html"))
}
