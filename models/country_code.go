package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
)

type CountryCode struct {
	Alpha2Code        string   `json:"alpha_2_code" db:"alpha_2_code"`
	Alpha3Code        string   `json:"alpha_3_code" db:"alpha_3_code"`
	NumericCode       int       `json:"numeric_code" db:"numeric_code"`
	InternationalName string `json:"international_name" db:"international_name"`
	IsIndependent     string `json:"is_independent" db:"is_independant"`
	IsoStatus         string  `json:"iso_status" db:"iso_status"`
	CreatedAt         time.Time `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (c CountryCode) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// CountryCodes is not required by pop and may be deleted
type CountryCodes []CountryCode

// String is not required by pop and may be deleted
func (c CountryCodes) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (c *CountryCode) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (c *CountryCode) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (c *CountryCode) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
