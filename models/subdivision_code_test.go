package models

import (
	"fmt"
	"testing"
	"time"

	"github.com/gobuffalo/nulls"
)

func TestSubdivisionCode_String(t *testing.T) {
	tnow := time.Now()
	type fields struct {
		SubdivisionCode   string
		Alpha2Code        string
		InternationalName string
		Category          nulls.String
		CreatedAt         time.Time
		UpdatedAt         time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "to string",
			fields: fields{
				SubdivisionCode:   "BF-KOP",
				Alpha2Code:        "BF",
				InternationalName: "Koulpélogo",
				Category:          nulls.NewString("Province"),
				CreatedAt:         tnow,
				UpdatedAt:         tnow,
			},
			want: fmt.Sprintf("{\"subdivision_code\":\"BF-KOP\",\"alpha_2_code\":\"BF\",\"international_name\":\"Koulpélogo\",\"category\":\"Province\",\"created_at\":\"%[1]s\",\"updated_at\":\"%[1]s\"}", tnow.Format("2006-01-02T15:04:05.999999Z07:00")),
		},
		{
			name: "with nulls",
			fields: fields{
				SubdivisionCode:   "BF-KOP",
				Alpha2Code:        "BF",
				InternationalName: "Koulpélogo",
				Category:          nulls.NewString(""),
				CreatedAt:         tnow,
				UpdatedAt:         tnow,
			},
			want: fmt.Sprintf("{\"subdivision_code\":\"BF-KOP\",\"alpha_2_code\":\"BF\",\"international_name\":\"Koulpélogo\",\"category\":\"\",\"created_at\":\"%[1]s\",\"updated_at\":\"%[1]s\"}", tnow.Format("2006-01-02T15:04:05.999999Z07:00")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sc := &SubdivisionCode{
				SubdivisionCode:   tt.fields.SubdivisionCode,
				Alpha2Code:        tt.fields.Alpha2Code,
				InternationalName: tt.fields.InternationalName,
				Category:          tt.fields.Category,
				CreatedAt:         tt.fields.CreatedAt,
				UpdatedAt:         tt.fields.UpdatedAt,
			}
			if got := sc.String(); got != tt.want {
				t.Errorf("SubdivisionCode.String() = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestSubdivisionCodes_String(t *testing.T) {
	tnow := time.Now()
	tests := []struct {
		name string
		scs  *SubdivisionCodes
		want string
	}{
		{
			name: "default",
			scs: &SubdivisionCodes{
				{
					SubdivisionCode:   "BF-KOP",
					Alpha2Code:        "BF",
					InternationalName: "Koulpélogo",
					Category:          nulls.NewString("Province"),
					CreatedAt:         tnow,
					UpdatedAt:         tnow,
				},
				{
					SubdivisionCode:   "CA-QC",
					Alpha2Code:        "CA",
					InternationalName: "Quebec",
					Category:          nulls.NewString("Province"),
					CreatedAt:         tnow,
					UpdatedAt:         tnow,
				},
			},
			want: fmt.Sprintf("[{\"subdivision_code\":\"BF-KOP\",\"alpha_2_code\":\"BF\",\"international_name\":\"Koulpélogo\",\"category\":\"Province\",\"created_at\":\"%[1]s\",\"updated_at\":\"%[1]s\"},{\"subdivision_code\":\"CA-QC\",\"alpha_2_code\":\"CA\",\"international_name\":\"Quebec\",\"category\":\"Province\",\"created_at\":\"%[1]s\",\"updated_at\":\"%[1]s\"}]", tnow.Format("2006-01-02T15:04:05.999999Z07:00")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.scs.String(); got != tt.want {
				t.Errorf("SubdivisionCodes.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
