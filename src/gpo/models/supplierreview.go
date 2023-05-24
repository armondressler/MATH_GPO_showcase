package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gofrs/uuid"
)

// Supplierreview is used by pop to map your supplierreviews database table to your go code.
type Supplierreview struct {
	ID         uuid.UUID `json:"id" db:"id"`
	Metric     string    `json:"metric" db:"metric"`
	MinValue   int       `json:"min_value" db:"min_value"`
	MaxValue   int       `json:"max_value" db:"max_value"`
	Value      int       `json:"value" db:"value"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
	SupplierID uuid.UUID `db:"supplier_id"`
	Supplier   *Supplier `json:"Supplier,omitempty" belongs_to:"Supplier"`
	UserID     uuid.UUID `db:"user_id"`
	User       *User     `json:"User,omitempty" belongs_to:"User"`
}

// String is not required by pop and may be deleted
func (s Supplierreview) String() string {
	js, _ := json.Marshal(s)
	return string(js)
}

// Supplierreviews is not required by pop and may be deleted
type Supplierreviews []Supplierreview

// String is not required by pop and may be deleted
func (s Supplierreviews) String() string {
	js, _ := json.Marshal(s)
	return string(js)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (s *Supplierreview) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (s *Supplierreview) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (s *Supplierreview) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
