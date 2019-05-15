package models

import (
	"encoding/json"
	"github.com/gobuffalo/nulls"
	"time"
)

type CountryName struct {
	ID                 int          `json:"id" db:"id"`
	Alpha2Code         string       `json:"alpha_2_code" db:"alpha_2_code"`
	LanguageAlpha2Code string       `json:"language_alpha_2_code" db:"language_alpha_2_code"`
	Name               string       `json:"name" db:"name"`
	FullName           nulls.String `json:"full_name" db:"full_name"`
	CreatedAt          time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt          time.Time    `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (cn CountryName) String() string {
	jc, _ := json.Marshal(cn)
	return string(jc)
}

// CountryCodes is not required by pop and may be deleted
type CountryNames []CountryName

// String is not required by pop and may be deleted
func (cns CountryNames) String() string {
	jc, _ := json.Marshal(cns)
	return string(jc)
}

func (cns CountryNames) Count() int {
	return len(cns)
}
