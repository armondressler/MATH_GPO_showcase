package actions

import (
	"fmt"
	"gpo/models"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
)

func setCompany(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}
	company := &models.Company{}
	if err := tx.Find(company, c.Param("company_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}
	c.Set("company", company)
	return nil
}

func setGpo(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}
	gpo := &models.Gpo{}
	if err := tx.Find(gpo, c.Param("gpo_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}
	c.Set("gpo", gpo)
	return nil
}

func setContractor(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}
	contractor := &models.Contractor{}
	if err := tx.Eager().Find(contractor, c.Param("contractor_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}
	c.Set("contractor", contractor)
	return nil
}

func setSupplier(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}
	supplier := &models.Supplier{}
	if err := tx.Eager().Find(supplier, c.Param("supplier_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}
	c.Set("supplier", supplier)
	return nil
}
