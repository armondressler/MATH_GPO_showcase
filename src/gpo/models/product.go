package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/nulls"
	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gofrs/uuid"
)

// Product is used by pop to map your products database table to your go code.
type Product struct {
	ID              uuid.UUID    `json:"id" db:"id"`
	Name            string       `json:"name" db:"name"`
	IsGdsn          bool         `json:"is_gdsn" db:"is_gdsn"`
	GdsnCountryCode nulls.String `json:"gdsn_country_code" db:"gdsn_country_code"`
	GdsnGtin        nulls.String `json:"gdsn_gtin" db:"gdsn_gtin"`
	GdsnGln         nulls.String `json:"gdsn_gln" db:"gdsn_gln"`
	CreatedAt       time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time    `json:"updated_at" db:"updated_at"`
	SupplierID      uuid.UUID    `db:"supplier_id"`
	Supplier        *Supplier    `json:"Supplier,omitempty" belongs_to:"Supplier"`
}

// String is not required by pop and may be deleted
func (p Product) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// Products is not required by pop and may be deleted
type Products []Product

// String is not required by pop and may be deleted
func (p Products) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (p *Product) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (p *Product) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (p *Product) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
