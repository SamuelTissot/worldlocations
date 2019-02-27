package models

import (
	"encoding/json"
	"time"
)

type LanguageCode struct {
	LanguageAlpha2Code string    `json:"language_alpha_2_code" db:"language_alpha_2_code"`
	LanguageAlpha3Code string    `json:"language_alpha_3_code" db:"language_alpha_3_code"`
	CreatedAt          time.Time `json:"created_at" db:"created_at"`
	UpdatedAt          time.Time `json:"updated_at" db:"updated_at"`
}

func (lc LanguageCode) String() string {
	jc, _ := json.Marshal(lc)
	return string(jc)
}

type LanguageCodes []LanguageCode

func (lcs LanguageCodes) String() string {
	jc, _ := json.Marshal(lcs)
	return string(jc)
}

func (lcs LanguageCodes) Count() int {
	return len(lcs)
}
