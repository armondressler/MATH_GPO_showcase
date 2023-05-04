package actions

import (
	"github.com/gobuffalo/buffalo"
)

func Logout(c buffalo.Context) error {
	AuthDestroy(c)
	return c.Redirect(302, "/login/show")
}
