package actions

import (
	"fmt"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/x/responder"

	"gpo/models"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Suppliercertificate)
// DB Table: Plural (suppliercertificates)
// Resource: Plural (Suppliercertificates)
// Path: Plural (/suppliercertificates)
// View Template Folder: Plural (/templates/suppliercertificates/)

// SuppliercertificatesResource is the resource for the Suppliercertificate model
type SuppliercertificatesResource struct {
	buffalo.Resource
}

// List gets all Suppliercertificates. This function is mapped to the path
// GET /suppliercertificates
func (v SuppliercertificatesResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	suppliercertificates := &models.Suppliercertificates{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Suppliercertificates from the DB
	if err := q.All(suppliercertificates); err != nil {
		return err
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// Add the paginator to the context so it can be used in the template.
		c.Set("pagination", q.Paginator)

		c.Set("suppliercertificates", suppliercertificates)
		return c.Render(http.StatusOK, r.HTML("suppliercertificates/index.plush.html"))
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(200, r.JSON(suppliercertificates))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(200, r.XML(suppliercertificates))
	}).Respond(c)
}

// Show gets the data for one Suppliercertificate. This function is mapped to
// the path GET /suppliercertificates/{suppliercertificate_id}
func (v SuppliercertificatesResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Suppliercertificate
	suppliercertificate := &models.Suppliercertificate{}

	// To find the Suppliercertificate the parameter suppliercertificate_id is used.
	if err := tx.Find(suppliercertificate, c.Param("suppliercertificate_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		c.Set("suppliercertificate", suppliercertificate)

		return c.Render(http.StatusOK, r.HTML("suppliercertificates/show.plush.html"))
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(200, r.JSON(suppliercertificate))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(200, r.XML(suppliercertificate))
	}).Respond(c)
}

// New renders the form for creating a new Suppliercertificate.
// This function is mapped to the path GET /suppliercertificates/new
func (v SuppliercertificatesResource) New(c buffalo.Context) error {
	c.Set("suppliercertificate", &models.Suppliercertificate{})

	return c.Render(http.StatusOK, r.HTML("suppliercertificates/new.plush.html"))
}

// Create adds a Suppliercertificate to the DB. This function is mapped to the
// path POST /suppliercertificates
func (v SuppliercertificatesResource) Create(c buffalo.Context) error {
	// Allocate an empty Suppliercertificate
	suppliercertificate := &models.Suppliercertificate{}

	// Bind suppliercertificate to the html form elements
	if err := c.Bind(suppliercertificate); err != nil {
		return err
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(suppliercertificate)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		return responder.Wants("html", func(c buffalo.Context) error {
			// Make the errors available inside the html template
			c.Set("errors", verrs)

			// Render again the new.html template that the user can
			// correct the input.
			c.Set("suppliercertificate", suppliercertificate)

			return c.Render(http.StatusUnprocessableEntity, r.HTML("suppliercertificates/new.plush.html"))
		}).Wants("json", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.JSON(verrs))
		}).Wants("xml", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.XML(verrs))
		}).Respond(c)
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// If there are no errors set a success message
		c.Flash().Add("success", T.Translate(c, "suppliercertificate.created.success"))

		// and redirect to the show page
		return c.Redirect(http.StatusSeeOther, "/suppliercertificates/%v", suppliercertificate.ID)
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusCreated, r.JSON(suppliercertificate))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusCreated, r.XML(suppliercertificate))
	}).Respond(c)
}

// Edit renders a edit form for a Suppliercertificate. This function is
// mapped to the path GET /suppliercertificates/{suppliercertificate_id}/edit
func (v SuppliercertificatesResource) Edit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Suppliercertificate
	suppliercertificate := &models.Suppliercertificate{}

	if err := tx.Find(suppliercertificate, c.Param("suppliercertificate_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	c.Set("suppliercertificate", suppliercertificate)
	return c.Render(http.StatusOK, r.HTML("suppliercertificates/edit.plush.html"))
}

// Update changes a Suppliercertificate in the DB. This function is mapped to
// the path PUT /suppliercertificates/{suppliercertificate_id}
func (v SuppliercertificatesResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Suppliercertificate
	suppliercertificate := &models.Suppliercertificate{}

	if err := tx.Find(suppliercertificate, c.Param("suppliercertificate_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	// Bind Suppliercertificate to the html form elements
	if err := c.Bind(suppliercertificate); err != nil {
		return err
	}

	verrs, err := tx.ValidateAndUpdate(suppliercertificate)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		return responder.Wants("html", func(c buffalo.Context) error {
			// Make the errors available inside the html template
			c.Set("errors", verrs)

			// Render again the edit.html template that the user can
			// correct the input.
			c.Set("suppliercertificate", suppliercertificate)

			return c.Render(http.StatusUnprocessableEntity, r.HTML("suppliercertificates/edit.plush.html"))
		}).Wants("json", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.JSON(verrs))
		}).Wants("xml", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.XML(verrs))
		}).Respond(c)
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// If there are no errors set a success message
		c.Flash().Add("success", T.Translate(c, "suppliercertificate.updated.success"))

		// and redirect to the show page
		return c.Redirect(http.StatusSeeOther, "/suppliercertificates/%v", suppliercertificate.ID)
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.JSON(suppliercertificate))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.XML(suppliercertificate))
	}).Respond(c)
}

// Destroy deletes a Suppliercertificate from the DB. This function is mapped
// to the path DELETE /suppliercertificates/{suppliercertificate_id}
func (v SuppliercertificatesResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Suppliercertificate
	suppliercertificate := &models.Suppliercertificate{}

	// To find the Suppliercertificate the parameter suppliercertificate_id is used.
	if err := tx.Find(suppliercertificate, c.Param("suppliercertificate_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	if err := tx.Destroy(suppliercertificate); err != nil {
		return err
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// If there are no errors set a flash message
		c.Flash().Add("success", T.Translate(c, "suppliercertificate.destroyed.success"))

		// Redirect to the index page
		return c.Redirect(http.StatusSeeOther, "/suppliercertificates")
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.JSON(suppliercertificate))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.XML(suppliercertificate))
	}).Respond(c)
}