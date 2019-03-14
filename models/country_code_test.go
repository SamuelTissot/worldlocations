package models

import (
	"fmt"
	"github.com/gobuffalo/nulls"
	"strings"
	"testing"
	"time"
)

func Test_CountryCode(t *testing.T) {
	cc := CountryCode{
		Alpha2Code:        "CA",
		Alpha3Code:        "CAN",
		NumericCode:       nulls.NewInt(124),
		InternationalName: "Canada",
		IsIndependent:     nulls.NewInt64(1),
		IsoStatus:         "officially-assigned",
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}

	str := cc.String()
	if !strings.Contains(str, "\"alpha_2_code\":\"CA\"") {
		t.Fatal(fmt.Sprintf("country_code  string miss match looing for: %s, in: %s", "\"alpha_2_code\":\"CA\"", str))
	}
}
