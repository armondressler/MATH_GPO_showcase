package actions

import (
	"fmt"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/x/responder"
	"github.com/pkg/errors"

	"gpo/models"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Company)
// DB Table: Plural (companies)
// Resource: Plural (Companies)
// Path: Plural (/companies)
// View Template Folder: Plural (/templates/companies/)

// CompaniesResource is the resource for the Company model
type CompaniesResource struct {
	buffalo.Resource
}

// List gets all Companies. This function is mapped to the path
// GET /companies
func (v CompaniesResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	companies := &models.Companies{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Companies from the DB
	if err := q.All(companies); err != nil {
		return err
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// Add the paginator to the context so it can be used in the template.
		c.Set("pagination", q.Paginator)

		c.Set("companies", companies)
		return c.Render(http.StatusOK, r.HTML("companies/index.plush.html"))
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(200, r.JSON(companies))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(200, r.XML(companies))
	}).Respond(c)
}

// Show gets the data for one Company. This function is mapped to
// the path GET /companies/{company_id}
func (v CompaniesResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Company
	company := &models.Company{}

	// To find the Company the parameter company_id is used.
	if err := tx.Find(company, c.Param("company_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}
	addBreadcrumbs(c)

	return responder.Wants("html", func(c buffalo.Context) error {
		c.Set("company", company)

		employees, err := GetEmployees(company.ID)
		if err != nil {
			return errors.WithStack(err)
		}
		c.Set("employees", employees)

		gpoes, err := GetGPOs(company.ID)
		if err != nil {
			return errors.WithStack(err)
		}
		c.Set("gpoes", gpoes)

		return c.Render(http.StatusOK, r.HTML("companies/show.plush.html"))
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(200, r.JSON(company))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(200, r.XML(company))
	}).Respond(c)
}

// New renders the form for creating a new Company.
// This function is mapped to the path GET /companies/new
func (v CompaniesResource) New(c buffalo.Context) error {
	c.Set("company", &models.Company{})

	return c.Render(http.StatusOK, r.HTML("companies/new.plush.html"))
}

// Create adds a Company to the DB. This function is mapped to the
// path POST /companies
func (v CompaniesResource) Create(c buffalo.Context) error {
	// Allocate an empty Company
	company := &models.Company{}

	// Bind company to the html form elements
	if err := c.Bind(company); err != nil {
		return err
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(company)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		return responder.Wants("html", func(c buffalo.Context) error {
			// Make the errors available inside the html template
			c.Set("errors", verrs)

			// Render again the new.html template that the user can
			// correct the input.
			c.Set("company", company)

			return c.Render(http.StatusUnprocessableEntity, r.HTML("companies/new.plush.html"))
		}).Wants("json", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.JSON(verrs))
		}).Wants("xml", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.XML(verrs))
		}).Respond(c)
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// If there are no errors set a success message
		c.Flash().Add("success", T.Translate(c, "company.created.success"))

		// and redirect to the show page
		return c.Redirect(http.StatusSeeOther, "/companies/%v", company.ID)
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusCreated, r.JSON(company))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusCreated, r.XML(company))
	}).Respond(c)
}

// Edit renders a edit form for a Company. This function is
// mapped to the path GET /companies/{company_id}/edit
func (v CompaniesResource) Edit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Company
	company := &models.Company{}

	if err := tx.Find(company, c.Param("company_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	c.Set("company", company)
	return c.Render(http.StatusOK, r.HTML("companies/edit.plush.html"))
}

// Update changes a Company in the DB. This function is mapped to
// the path PUT /companies/{company_id}
func (v CompaniesResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Company
	company := &models.Company{}

	if err := tx.Find(company, c.Param("company_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	// Bind Company to the html form elements
	if err := c.Bind(company); err != nil {
		return err
	}

	verrs, err := tx.ValidateAndUpdate(company)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		return responder.Wants("html", func(c buffalo.Context) error {
			// Make the errors available inside the html template
			c.Set("errors", verrs)

			// Render again the edit.html template that the user can
			// correct the input.
			c.Set("company", company)

			return c.Render(http.StatusUnprocessableEntity, r.HTML("companies/edit.plush.html"))
		}).Wants("json", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.JSON(verrs))
		}).Wants("xml", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.XML(verrs))
		}).Respond(c)
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// If there are no errors set a success message
		c.Flash().Add("success", T.Translate(c, "company.updated.success"))

		// and redirect to the show page
		return c.Redirect(http.StatusSeeOther, "/companies/%v", company.ID)
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.JSON(company))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.XML(company))
	}).Respond(c)
}

// Destroy deletes a Company from the DB. This function is mapped
// to the path DELETE /companies/{company_id}
func (v CompaniesResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Company
	company := &models.Company{}

	// To find the Company the parameter company_id is used.
	if err := tx.Find(company, c.Param("company_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	if err := tx.Destroy(company); err != nil {
		return err
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// If there are no errors set a flash message
		c.Flash().Add("success", T.Translate(c, "company.destroyed.success"))

		// Redirect to the index page
		return c.Redirect(http.StatusSeeOther, "/companies")
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.JSON(company))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.XML(company))
	}).Respond(c)
}
