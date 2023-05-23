package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/pkg/errors"
)

// SuppliermanagementShow default implementation.
func SuppliermanagementShow(c buffalo.Context) error {
	err := setGpo(c)
	if err != nil {
		return errors.WithStack(err)
	}
	err = setCompany(c)
	if err != nil {
		return errors.WithStack(err)
	}
	addBreadcrumbs(c)

	return c.Render(http.StatusOK, r.HTML("suppliermanagement/show.html"))
}
