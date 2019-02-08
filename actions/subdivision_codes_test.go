package actions

import (
	"encoding/json"
	"fmt"
	"github.com/gobuffalo/pop/nulls"
	"time"
	"worldlocations/models"
)

func (as *ActionSuite) Test_SubdivisionCodes_List() {
	as.LoadFixture("subdivision_codes")
	res := as.JSON("/v1/subdivisions").Get()

	as.Equal(200, res.Code)

	body := res.Body.String()
	var subdivisionCodes models.SubdivisionCodes
	as.NoError(as.DB.All(&subdivisionCodes))
	for _, c := range subdivisionCodes {
		as.Contains(body, c.SubdivisionCode)
	}
}

func (as *ActionSuite) Test_SubdivisionCodes_CountrySubdivisions() {
	as.LoadFixture("subdivision_codes")
	testTime, err := time.Parse("2006-01-02 03:04:05", "2019-01-03 08:52:06")
	if err != nil {
		as.FailNow("could not generate test timestamp error: %s", err.Error())
	}

	tests := []struct {
		name       string
		alpha2Code string
		want       models.SubdivisionCodes
	}{
		{
			name:       "Canada",
			alpha2Code: "CA",
			want: []models.SubdivisionCode{
				{
					SubdivisionCode:   "CA-QC",
					Alpha2Code:        "CA",
					InternationalName: "Quebec",
					Category:          nulls.NewString("Province"),
					CreatedAt:         testTime,
					UpdatedAt:         testTime,
				},
				{
					SubdivisionCode:   "CA-ON",
					Alpha2Code:        "CA",
					InternationalName: "Ontario",
					Category:          nulls.NewString("Province"),
					CreatedAt:         testTime,
					UpdatedAt:         testTime,
				},
			},
		},
		{
			name:       "USA",
			alpha2Code: "US",
			want: []models.SubdivisionCode{
				{
					SubdivisionCode:   "US-NY",
					Alpha2Code:        "US",
					InternationalName: "New-York",
					Category:          nulls.NewString("State"),
					CreatedAt:         testTime,
					UpdatedAt:         testTime,
				},
			},
		},
	}
	for _, tt := range tests {

		res := as.JSON(fmt.Sprintf("/v1/countries/%s/subdivisions/", tt.alpha2Code)).Get()
		as.Equal(200, res.Code)

		scs := SubdivisionCodes{}
		err := json.Unmarshal(res.Body.Bytes(), &scs)
		if err != nil {
			as.FailNow(err.Error())
		}
		for _, s := range *scs.Data {
			s := s
			as.True(func(s models.SubdivisionCode) bool {
				for _, w := range tt.want {
					if w == s {
						return true
					}
				}
				return false
			}(s))
		}
	}
}
