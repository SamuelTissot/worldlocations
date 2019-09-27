package models

import (
	"encoding/json"
	"github.com/gobuffalo/nulls"
	"time"
)

type SubdivisionCode struct {
	SubdivisionCode   string       `json:"subdivision_code" db:"subdivision_code"`
	Alpha2Code        string       `json:"alpha_2_code" db:"alpha_2_code"`
	InternationalName string       `json:"international_name" db:"international_name"`
	Category          nulls.String `json:"category" db:"category"`
	CreatedAt         time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time    `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (sc SubdivisionCode) String() string {
	jc, _ := json.Marshal(sc)
	return string(jc)
}

// CountryCodes is not required by pop and may be deleted
type SubdivisionCodes []SubdivisionCode

// String is not required by pop and may be deleted
func (scs SubdivisionCodes) String() string {
	jc, _ := json.Marshal(scs)
	return string(jc)
}

func (scs SubdivisionCodes) Count() int {
	return len(scs)
}

func (scs SubdivisionCodes) Paginate(p, count int) (Model, bool) {
	if scs.Count() < (p-1)*count {
		return SubdivisionCodes{}, false
	}
	if scs.Count() > count {
		return scs[(p-1)*count : p*count], true
	}
	return scs[(p-1)*count:], false
}
