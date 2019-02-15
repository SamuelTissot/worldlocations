package models

import (
	"encoding/json"
	"time"
)

type City struct {
	ID                int       `db:"id"`
	Alpha2Code        string    `json:"alpha_2_code" db:"alpha_2_code"`
	SubdivisionCode   string    `json:"subdivision_code" db:"subdivision_code"`
	Locode            string    `json:"locode" db:"locode"`
	Name              string    `json:"name" db:"name"`
	InternationalName string    `json:"international_name" db:"international_name"`
	IataCode          string    `json:"iata_code" db:"iata_code"`
	LatLong           string    `json:"latitude_longitude" db:"latitude_longitude"`
	CreatedAt         time.Time `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (c *City) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// CountryCodes is not required by pop and may be deleted
type Cities []City

// String is not required by pop and may be deleted
func (cs *Cities) String() string {
	jc, _ := json.Marshal(cs)
	return string(jc)
}
