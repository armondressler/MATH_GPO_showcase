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
