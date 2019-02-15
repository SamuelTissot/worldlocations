package actions

import (
	"encoding/json"
	"fmt"
	"time"
	"worldlocations/models"
)

func (as *ActionSuite) TestCities_List() {
	as.LoadFixture("cities")
	res := as.JSON("/v1/cities").Get()

	as.Equal(200, res.Code)

	body := res.Body.String()
	var subdivisionCodes models.SubdivisionCodes
	as.NoError(as.DB.All(&subdivisionCodes))
	for _, c := range subdivisionCodes {
		as.Contains(body, c.SubdivisionCode)
	}
}

func (as *ActionSuite) TestCities_CountryCities() {
	as.LoadFixture("cities")
	testTime, err := time.Parse("2006-01-02 03:04:05", "2019-01-03 08:52:06")
	if err != nil {
		as.FailNow("could not generate test timestamp error: %s", err.Error())
	}

	tests := []struct {
		name       string
		alpha2Code string
		want       models.Cities
	}{
		{
			name:       "Canada",
			alpha2Code: "CA",
			want: models.Cities{
				{
					ID:                1,
					Alpha2Code:        "CA",
					SubdivisionCode:   "CA-QC",
					Locode:            "MTR",
					Name:              "Montreal",
					InternationalName: "Montreal",
					IataCode:          "YMQ",
					CreatedAt:         testTime,
					UpdatedAt:         testTime,
				},
				{
					ID:                2,
					Alpha2Code:        "CA",
					SubdivisionCode:   "CA-QC",
					Locode:            "LDT",
					Name:              "Lac-Drolet",
					InternationalName: "Lac-Drolet",
					IataCode:          "LDT",
					LatLong:           "4543N 07051W",
					CreatedAt:         testTime,
					UpdatedAt:         testTime,
				},
				{
					ID:                5,
					Alpha2Code:        "CA",
					SubdivisionCode:   "CA-ON",
					Locode:            "WFI",
					Name:              "Wolfe Island",
					InternationalName: "Wolfe Island",
					IataCode:          "WFI",
					LatLong:           "4411N 07626W",
					CreatedAt:         testTime,
					UpdatedAt:         testTime,
				},
			},
		},
		{
			name:       "USA",
			alpha2Code: "US",
			want: models.Cities{
				{
					ID:                3,
					Alpha2Code:        "US",
					SubdivisionCode:   "US-NY",
					Locode:            "YWG",
					Name:              "Woodridge",
					InternationalName: "Woodridge",
					IataCode:          "",
					LatLong:           "4142N 07434W",
					CreatedAt:         testTime,
					UpdatedAt:         testTime,
				},
				{
					ID:                4,
					Alpha2Code:        "US",
					SubdivisionCode:   "US-NY",
					Locode:            "YTW",
					Name:              "Youngstown",
					InternationalName: "Youngstown",
					IataCode:          "",
					LatLong:           "4314N 07903W",
					CreatedAt:         testTime,
					UpdatedAt:         testTime,
				},
			},
		},
	}

	for _, tt := range tests {

		res := as.JSON(fmt.Sprintf("/v1/countries/%s/cities/", tt.alpha2Code)).Get()
		as.Equal(200, res.Code)

		ls := &Cities{}
		err := json.Unmarshal(res.Body.Bytes(), ls)
		if err != nil {
			as.FailNow(err.Error())
		}
		as.Equal(tt.want, *ls.Data)
	}
}

func (as *ActionSuite) TestCities_SubdivisionCities() {
	as.LoadFixture("cities")
	testTime, err := time.Parse("2006-01-02 03:04:05", "2019-01-03 08:52:06")
	if err != nil {
		as.FailNow("could not generate test timestamp error: %s", err.Error())
	}

	tests := []struct {
		name            string
		subdivisionCode string
		want            models.Cities
	}{
		{
			name:            "QC",
			subdivisionCode: "CA-QC",
			want: models.Cities{
				{
					ID:                1,
					Alpha2Code:        "CA",
					SubdivisionCode:   "CA-QC",
					Locode:            "MTR",
					Name:              "Montreal",
					InternationalName: "Montreal",
					IataCode:          "YMQ",
					CreatedAt:         testTime,
					UpdatedAt:         testTime,
				},
				{
					ID:                2,
					Alpha2Code:        "CA",
					SubdivisionCode:   "CA-QC",
					Locode:            "LDT",
					Name:              "Lac-Drolet",
					InternationalName: "Lac-Drolet",
					IataCode:          "LDT",
					LatLong:           "4543N 07051W",
					CreatedAt:         testTime,
					UpdatedAt:         testTime,
				},
			},
		},
		{
			name:            "NY",
			subdivisionCode: "US-NY",
			want: models.Cities{
				{
					ID:                3,
					Alpha2Code:        "US",
					SubdivisionCode:   "US-NY",
					Locode:            "YWG",
					Name:              "Woodridge",
					InternationalName: "Woodridge",
					IataCode:          "",
					LatLong:           "4142N 07434W",
					CreatedAt:         testTime,
					UpdatedAt:         testTime,
				},
				{
					ID:                4,
					Alpha2Code:        "US",
					SubdivisionCode:   "US-NY",
					Locode:            "YTW",
					Name:              "Youngstown",
					InternationalName: "Youngstown",
					IataCode:          "",
					LatLong:           "4314N 07903W",
					CreatedAt:         testTime,
					UpdatedAt:         testTime,
				},
			},
		},
	}

	for _, tt := range tests {

		res := as.JSON(fmt.Sprintf("/v1/subdivisions/%s/cities/", tt.subdivisionCode)).Get()
		as.Equal(200, res.Code)

		ls := &Cities{}
		err := json.Unmarshal(res.Body.Bytes(), ls)
		if err != nil {
			as.FailNow(err.Error())
		}
		as.Equal(tt.want, *ls.Data)
	}

}

func (as *ActionSuite) TestCities_Show() {
	as.LoadFixture("cities")
	testTime, err := time.Parse("2006-01-02 03:04:05", "2019-01-03 08:52:06")
	if err != nil {
		as.FailNow("could not generate test timestamp error: %s", err.Error())
	}

	tests := []struct {
		name string
		id   int
		want models.Cities
	}{
		{
			name: "QC",
			id:   1,
			want: models.Cities{
				{
					ID:                1,
					Alpha2Code:        "CA",
					SubdivisionCode:   "CA-QC",
					Locode:            "MTR",
					Name:              "Montreal",
					InternationalName: "Montreal",
					IataCode:          "YMQ",
					CreatedAt:         testTime,
					UpdatedAt:         testTime,
				},
			},
		},
		{
			name: "NY",
			id:   3,
			want: models.Cities{
				{
					ID:                3,
					Alpha2Code:        "US",
					SubdivisionCode:   "US-NY",
					Locode:            "YWG",
					Name:              "Woodridge",
					InternationalName: "Woodridge",
					IataCode:          "",
					LatLong:           "4142N 07434W",
					CreatedAt:         testTime,
					UpdatedAt:         testTime,
				},
			},
		},
	}

	for _, tt := range tests {

		res := as.JSON(fmt.Sprintf("/v1/cities/%d/", tt.id)).Get()
		as.Equal(200, res.Code)

		ls := &Cities{}
		err := json.Unmarshal(res.Body.Bytes(), ls)
		if err != nil {
			as.FailNow(err.Error())
		}
		as.Equal(tt.want, *ls.Data)
	}
}
