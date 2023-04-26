package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

// SuppliersShow default implementation.
func SuppliersShow(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("suppliers/show.html"))
}

// SuppliersList default implementation.
func SuppliersList(c buffalo.Context) error {
	c.Set("names", []string{"John", "Paul", "George", "Ringo"})
	return c.Render(http.StatusOK, r.HTML("suppliers/list.html"))
}

// SuppliersCreate default implementation.
func SuppliersCreate(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("suppliers/create.html"))
}
