package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gofrs/uuid"
)

// Suppliercertificate is used by pop to map your suppliercertificates database table to your go code.
type Suppliercertificate struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Issuer    string    `json:"issuer" db:"issuer"`
	IssuedAt  time.Time `json:"issued_at" db:"issued_at"`
	ExpiresAt time.Time `json:"expires_at" db:"expires_at"`
	DataLink  string    `json:"data_link" db:"data_link"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (s Suppliercertificate) String() string {
	js, _ := json.Marshal(s)
	return string(js)
}

// Suppliercertificates is not required by pop and may be deleted
type Suppliercertificates []Suppliercertificate

// String is not required by pop and may be deleted
func (s Suppliercertificates) String() string {
	js, _ := json.Marshal(s)
	return string(js)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (s *Suppliercertificate) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (s *Suppliercertificate) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (s *Suppliercertificate) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
