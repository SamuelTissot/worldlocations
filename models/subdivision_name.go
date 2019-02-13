package models

import (
	"encoding/json"
	"github.com/gobuffalo/pop/nulls"
	"time"
)

type SubdivisionName struct {
	SubdivisionCode    string       `json:"subdivision_code" db:"subdivision_code"`
	LanguageAlpha2Code string       `json:"language_alpha_2_code" db:"language_alpha_2_code"`
	Name               string       `json:"name" db:"name"`
	LocalVariation     nulls.String `json:"local_variation" db:"local_variation"`
	CreatedAt          time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt          time.Time    `json:"updated_at" db:"updated_at"`
}

func (sn SubdivisionName) String() string {
	jc, _ := json.Marshal(sn)
	return string(jc)
}

type SubdivisionNames []SubdivisionName

func (sns SubdivisionNames) String() string {
	jc, _ := json.Marshal(sns)
	return string(jc)
}
