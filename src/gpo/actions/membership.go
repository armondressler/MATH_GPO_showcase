package actions

import (
	"fmt"
	"gpo/models"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/pkg/errors"
)

// MembershipShow default implementation.
func MembershipShow(c buffalo.Context) error {
	gpoCompaniesEmployees := map[*models.Company][]*models.User{}

	err := setGpo(c)
	if err != nil {
		return errors.WithStack(err)
	}
	err = setCompany(c)
	if err != nil {
		return errors.WithStack(err)
	}
	gpo := c.Value("gpo")

	gpoCompanies, err := GetGpoCompanies(gpo.(*models.Gpo).ID)
	if err != nil {
		return errors.WithStack(err)
	}

	c.Set("gpoCompanies", gpoCompanies)
	for _, gpoCompany := range gpoCompanies {
		employees, err := GetEmployees(gpoCompany.ID)
		if err != nil {
			return fmt.Errorf("getting employees of all companies in gpo %s: %s", gpo.(*models.Gpo).Name, err)
		}
		gpoCompaniesEmployees[gpoCompany] = employees
	}
	c.Set("gpoCompaniesEmployees", gpoCompaniesEmployees)

	gpoContractors, err := GetGPOContractors(gpo.(*models.Gpo).ID)
	if err != nil {
		return errors.WithStack(err)
	}
	c.Set("gpoContractors", gpoContractors)

	addBreadcrumbs(c)

	return c.Render(http.StatusOK, r.HTML("membership/show.html"))
}
