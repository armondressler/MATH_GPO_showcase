package actions

import (
	"fmt"
	"gpo/models"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/x/responder"
	"github.com/pkg/errors"
)

type ContractorsResource struct {
	buffalo.Resource
}

// List default implementation.
func (v ContractorsResource) List(c buffalo.Context) error {
	//return c.Render(http.StatusOK, r.String("Contractor#List"))
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}
	contractors := &models.Contractors{}

	q := tx.PaginateFromParams(c.Params())

	mailcontains := c.Params().Get("mailcontains")
	if mailcontains != "" {
		c.Logger().Info("GOT A MAILCONTAINS:[", mailcontains, "]")
		q = q.RawQuery("select * from contractors where email ~ ?", mailcontains)
	}

	// Retrieve all contractors from the DB
	if err := q.Eager().All(contractors); err != nil {
		return err
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		c.Set("pagination", q.Paginator)
		c.Set("contractors", contractors)
		return c.Render(http.StatusNotFound, r.HTML("contractors/index.plush.html"))
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(200, r.JSON(contractors))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(200, r.XML(contractors))
	}).Respond(c)
}

// Show default implementation.
func (v ContractorsResource) Show(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.String("Contractor#Show"))
}

// Create default implementation.
func (v ContractorsResource) Create(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.String("Contractor#Create"))
}

// Update default implementation.
func (v ContractorsResource) Update(c buffalo.Context) error {
	err := setContractor(c)
	if err != nil {
		return errors.WithStack(err)
	}

	contractor, ok := c.Value("contractor").(*models.Contractor)
	if !ok {
		return fmt.Errorf("failed to fetch contractor from context")
	}

	if gpo_id := c.Params().Get("add-to-gpo"); gpo_id != "" {
		c.Logger().Info("CON: ", contractor.Gpoes)
		err = contractor.AddToGpo(gpo_id)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		return c.Render(http.StatusNotFound, r.String("use application/json"))
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(200, r.JSON(contractor))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(200, r.XML(contractor))
	}).Respond(c)
}

// Destroy default implementation.
func (v ContractorsResource) Destroy(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.String("Contractor#Destroy"))
}

// New default implementation.
func (v ContractorsResource) New(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.String("Contractor#New"))
}

// Edit default implementation.
func (v ContractorsResource) Edit(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.String("Contractor#Edit"))
}
