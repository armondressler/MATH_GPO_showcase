package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gofrs/uuid"
)

// Contractor is used by pop to map your contractors database table to your go code.
type Contractor struct {
	ID         uuid.UUID `db:"id"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
	Email      string    `db:"email"`
	FirstName  string    `db:"first_name"`
	LastName   string    `db:"last_name"`
	Provider   string    `db:"provider"`
	ProviderID string    `db:"provider_id"`
	Gpoes      []Gpo     `json:"gpoes" many_to_many:"contractors_gpoes" db:"-"`
	Companies  []Company `json:"companies" many_to_many:"contractors_companies" db:"-"`
}

// String is not required by pop and may be deleted
func (c Contractor) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Contractors is not required by pop and may be deleted
type Contractors []Contractor

func (c Contractor) AddToGpo(gpo_id string) error {
	tx := DB
	some_id, _ := uuid.NewV4()
	now := time.Now()
	q := tx.RawQuery("insert into contractors_gpoes(id, contractor_id, gpo_id, created_at, updated_at) values(?, ?, ?, ?, ?) on conflict(contractor_id, gpo_id) do nothing;", some_id, c.ID, gpo_id, now, now)
	err := q.Exec()
	return err
}

// String is not required by pop and may be deleted
func (c Contractors) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (c *Contractor) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (c *Contractor) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (c *Contractor) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
