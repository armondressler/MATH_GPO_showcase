package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	c.Set("render_without_navbar", "true")
	return c.Render(http.StatusOK, r.HTML("home/index.plush.html"))
}
