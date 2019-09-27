package actions

import (
	"encoding/json"
	"fmt"
	"github.com/gobuffalo/nulls"
	"strings"
	"time"
	"worldlocations/models"
)

/*
 *
 * Countries tests
 *
 */
func (as *ActionSuite) Test__CountryList() {
	as.LoadFixture("country_codes")
	res := as.JSON("/v1/countries").Get()

	as.Equal(200, res.Code)

	body := res.Body.String()
	var countryCodes models.CountryCodes
	as.NoError(as.DB.All(&countryCodes))
	for _, c := range countryCodes {
		as.Contains(body, c.Alpha2Code)
	}
}

func (as *ActionSuite) Test__CountryShow() {
	as.LoadFixture("country_codes")
	res := as.JSON("/v1/countries/ca").Get()

	as.Equal(200, res.Code)
	body := res.Body.String()
	var countryCode models.CountryCode
	as.NoError(as.DB.Where("alpha_2_code = (?)", "CA").First(&countryCode))
	as.Contains(body, countryCode.Alpha2Code)
}

/*
 *
 * Countries Names tests
 *
 */
func (as *ActionSuite) Test_countriesNamesList() {
	as.LoadFixture("country_names")
	res := as.JSON("/v1/countries-names/").Get()

	as.Equal(200, res.Code)

	body := res.Body.String()
	var countryNames models.CountryNames
	as.NoError(as.DB.All(&countryNames))
	for _, c := range countryNames {
		as.Contains(body, c.Name)
	}
}

func (as *ActionSuite) Test_countryNames() {
	as.LoadFixture("country_names")
	testTime, err := time.Parse("2006-01-02 03:04:05", "2019-01-03 08:52:06")
	if err != nil {
		as.FailNow("could not generate test timestamp error: %s", err.Error())
	}
	tests := []struct {
		name        string
		Alpha2Codee string
		want        V1Handler
	}{
		{
			name:        "Canada",
			Alpha2Codee: "CA",
			want: V1Handler{
				Count: 2,
				Data: models.CountryNames{
					{
						ID:                 1,
						Alpha2Code:         "CA",
						LanguageAlpha2Code: "en",
						Name:               "Canada",
						CreatedAt:          testTime,
						UpdatedAt:          testTime,
					},
					{
						ID:                 2,
						Alpha2Code:         "CA",
						LanguageAlpha2Code: "fr",
						Name:               "Canada (le)",
						CreatedAt:          testTime,
						UpdatedAt:          testTime,
					},
				},
			},
		},
		{
			name:        "USA",
			Alpha2Codee: "US",
			want: V1Handler{
				Count: 2,
				Data: models.CountryNames{
					{
						ID:                 3,
						Alpha2Code:         "US",
						LanguageAlpha2Code: "en",
						Name:               "United States of America (the)",
						FullName:           nulls.NewString("the United States of America"),
						CreatedAt:          testTime,
						UpdatedAt:          testTime,
					},
					{
						ID:                 4,
						Alpha2Code:         "US",
						LanguageAlpha2Code: "fr",
						Name:               "États-Unis d'Amérique (les)",
						FullName:           nulls.NewString("les États-Unis d'Amérique"),
						CreatedAt:          testTime,
						UpdatedAt:          testTime,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		res := as.JSON(fmt.Sprintf("/v1/countries/%s/names", tt.Alpha2Codee)).Get()
		as.Equal(200, res.Code)
		wStr, err := json.Marshal(tt.want)
		if err != nil {
			as.FailNow(err.Error())
		}
		as.Equal(string(wStr), strings.TrimSpace(string(res.Body.Bytes())))
	}
}

/*
 *
 * SubdivisionNames tests
 *
 */
func (as *ActionSuite) Test_SubdivisionNames_List() {
	as.LoadFixture("subdivision_names")
	res := as.JSON("/v1/subdivisions/names/").Get()

	as.Equal(200, res.Code)

	body := res.Body.String()
	var subdivisionNames models.SubdivisionNames
	as.NoError(as.DB.All(&subdivisionNames))
	for _, c := range subdivisionNames {
		as.Contains(body, c.SubdivisionCode)
	}
}

func (as *ActionSuite) Test_SubdivisionNames_Show() {
	as.LoadFixture("subdivision_names")
	testTime, err := time.Parse("2006-01-02 03:04:05", "2019-01-03 08:52:06")
	if err != nil {
		as.FailNow("could not generate test timestamp error: %s", err.Error())
	}
	tests := []struct {
		name            string
		subdivisionCode string
		want            V1Handler
	}{
		{
			name:            "Quebec",
			subdivisionCode: "CA-QC",
			want: V1Handler{
				Count: 2,
				Data: models.SubdivisionNames{
					{
						SubdivisionCode:    "CA-QC",
						LanguageAlpha2Code: "en",
						Name:               "Quebec",
						CreatedAt:          testTime,
						UpdatedAt:          testTime,
					},
					{
						SubdivisionCode:    "CA-QC",
						LanguageAlpha2Code: "fr",
						Name:               "Québec",
						CreatedAt:          testTime,
						UpdatedAt:          testTime,
					},
				},
			},
		},
		{
			name:            "Bouches-du-Rhône",
			subdivisionCode: "FR-13",
			want: V1Handler{
				Count: 1,
				Data: models.SubdivisionNames{
					models.SubdivisionName{
						SubdivisionCode:    "FR-13",
						LanguageAlpha2Code: "fr",
						Name:               "Bouches-du-Rhône",
						LocalVariation:     nulls.NewString("local variant"),
						CreatedAt:          testTime,
						UpdatedAt:          testTime,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		res := as.JSON(fmt.Sprintf("/v1/subdivisions/%s/names/", tt.subdivisionCode)).Get()
		as.Equal(200, res.Code)
		wStr, err := json.Marshal(tt.want)
		if err != nil {
			as.FailNow(err.Error())
		}
		as.Equal(string(wStr), strings.TrimSpace(string(res.Body.Bytes())))
	}
}

/*
 *
 * Subdivisions tests
 *
 */
func (as *ActionSuite) Test_Subdivisions_List() {
	as.LoadFixture("subdivision_codes")
	res := as.JSON("/v1/subdivisions/").Get()

	as.Equal(200, res.Code)

	body := res.Body.String()
	var subdivisionCodes models.SubdivisionCodes
	as.NoError(as.DB.All(&subdivisionCodes))
	for _, c := range subdivisionCodes {
		as.Contains(body, c.SubdivisionCode)
	}
}

func (as *ActionSuite) Test_Subdivisions_Show() {
	as.LoadFixture("subdivision_codes")
	testTime, err := time.Parse("2006-01-02 03:04:05", "2019-01-03 08:52:06")
	if err != nil {
		as.FailNow("could not generate test timestamp error: %s", err.Error())
	}
	tests := []struct {
		name            string
		subdivisionCode string
		want            V1Handler
	}{
		{
			name:            "Quebec",
			subdivisionCode: "CA-QC",
			want: V1Handler{
				Count: 1,
				Data: models.SubdivisionCodes{
					models.SubdivisionCode{
						SubdivisionCode:   "CA-QC",
						Alpha2Code:        "CA",
						InternationalName: "Quebec",
						Category:          nulls.NewString("Province"),
						CreatedAt:         testTime,
						UpdatedAt:         testTime,
					},
				},
			},
		},
		{
			name:            "New-York",
			subdivisionCode: "US-NY",
			want: V1Handler{
				Count: 1,
				Data: models.SubdivisionCodes{
					models.SubdivisionCode{
						SubdivisionCode:   "US-NY",
						Alpha2Code:        "US",
						InternationalName: "New-York",
						Category:          nulls.NewString("State"),
						CreatedAt:         testTime,
						UpdatedAt:         testTime,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		res := as.JSON(fmt.Sprintf("/v1/subdivisions/%s/", tt.subdivisionCode)).Get()
		as.Equal(200, res.Code)
		wStr, err := json.Marshal(tt.want)
		if err != nil {
			as.FailNow(err.Error())
		}
		as.Equal(string(wStr), strings.TrimSpace(string(res.Body.Bytes())))
	}
}

func (as *ActionSuite) Test_CountrySubdivisions() {
	as.LoadFixture("subdivision_codes")
	testTime, err := time.Parse("2006-01-02 03:04:05", "2019-01-03 08:52:06")
	if err != nil {
		as.FailNow("could not generate test timestamp error: %s", err.Error())
	}
	tests := []struct {
		name       string
		alpha2Code string
		want       V1Handler
	}{
		{
			name:       "Canada",
			alpha2Code: "CA",
			want: V1Handler{
				Count: 2,
				Data: models.SubdivisionCodes{
					models.SubdivisionCode{
						SubdivisionCode:   "CA-QC",
						Alpha2Code:        "CA",
						InternationalName: "Quebec",
						Category:          nulls.NewString("Province"),
						CreatedAt:         testTime,
						UpdatedAt:         testTime,
					},
					models.SubdivisionCode{
						SubdivisionCode:   "CA-ON",
						Alpha2Code:        "CA",
						InternationalName: "Ontario",
						Category:          nulls.NewString("Province"),
						CreatedAt:         testTime,
						UpdatedAt:         testTime,
					},
				},
			},
		},
		{
			name:       "USA",
			alpha2Code: "US",
			want: V1Handler{
				Count: 1,
				Data: models.SubdivisionCodes{
					models.SubdivisionCode{
						SubdivisionCode:   "US-NY",
						Alpha2Code:        "US",
						InternationalName: "New-York",
						Category:          nulls.NewString("State"),
						CreatedAt:         testTime,
						UpdatedAt:         testTime,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		res := as.JSON(fmt.Sprintf("/v1/countries/%s/subdivisions/", tt.alpha2Code)).Get()
		as.Equal(200, res.Code)
		wStr, err := json.Marshal(tt.want)
		if err != nil {
			as.FailNow(err.Error())
		}
		as.Equal(string(wStr), strings.TrimSpace(string(res.Body.Bytes())))
	}
}

/*
 *
 * languages tests
 *
 */
func (as *ActionSuite) Test_Languages_List() {
	as.LoadFixture("language_codes")
	res := as.JSON("/v1/languages/").Get()

	as.Equal(200, res.Code)

	body := res.Body.String()
	var languageCodes models.LanguageCodes
	as.NoError(as.DB.All(&languageCodes))
	for _, c := range languageCodes {
		as.Contains(body, c.LanguageAlpha3Code)
	}
}

func (as *ActionSuite) Test_Languages_Show() {
	as.LoadFixture("language_codes")
	testTime, err := time.Parse("2006-01-02 03:04:05", "2019-01-03 08:52:06")
	if err != nil {
		as.FailNow("could not generate test timestamp error: %s", err.Error())
	}
	tests := []struct {
		name               string
		languageAlpha2Code string
		want               V1Handler
	}{
		{
			name:               "English",
			languageAlpha2Code: "en",
			want: V1Handler{
				Count: 1,
				Data: models.LanguageCodes{
					models.LanguageCode{
						LanguageAlpha2Code: "en",
						LanguageAlpha3Code: "eng",
						CreatedAt:          testTime,
						UpdatedAt:          testTime,
					},
				},
			},
		},
		{
			name:               "French",
			languageAlpha2Code: "fr",
			want: V1Handler{
				Count: 1,
				Data: models.LanguageCodes{
					models.LanguageCode{
						LanguageAlpha2Code: "fr",
						LanguageAlpha3Code: "fra",
						CreatedAt:          testTime,
						UpdatedAt:          testTime,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		res := as.JSON(fmt.Sprintf("/v1/languages/%s/", tt.languageAlpha2Code)).Get()
		as.Equal(200, res.Code)
		wStr, err := json.Marshal(tt.want)
		if err != nil {
			as.FailNow(err.Error())
		}
		as.Equal(string(wStr), strings.TrimSpace(string(res.Body.Bytes())))
	}
}

/*
 *
 * cities test
 *
 */
func (as *ActionSuite) Test_cities_List() {
	as.LoadFixture("cities")
	res := as.JSON("/v1/cities/").Get()

	as.Equal(200, res.Code)

	body := res.Body.String()
	var cities models.Cities
	as.NoError(as.DB.All(&cities))
	for _, c := range cities {
		as.Contains(body, c.Name)
	}
}

func (as *ActionSuite) Test_cities_Show() {
	as.LoadFixture("cities")
	testTime, err := time.Parse("2006-01-02 03:04:05", "2019-01-03 08:52:06")
	if err != nil {
		as.FailNow("could not generate test timestamp error: %s", err.Error())
	}
	tests := []struct {
		name string
		id   int
		want V1Handler
	}{
		{
			name: "Canada",
			id:   1,
			want: V1Handler{
				Count: 1,
				Data: models.Cities{
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
		},
		{
			name: "USA",
			id:   3,
			want: V1Handler{
				Count: 1,
				Data: models.Cities{
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
		},
	}
	for _, tt := range tests {
		res := as.JSON(fmt.Sprintf("/v1/cities/%d/", tt.id)).Get()
		as.Equal(200, res.Code)
		wStr, err := json.Marshal(tt.want)
		if err != nil {
			as.FailNow(err.Error())
		}
		as.Equal(string(wStr), strings.TrimSpace(string(res.Body.Bytes())))
	}
}

func (as *ActionSuite) Test_countryCities() {
	as.LoadFixture("cities")
	testTime, err := time.Parse("2006-01-02 03:04:05", "2019-01-03 08:52:06")
	if err != nil {
		as.FailNow("could not generate test timestamp error: %s", err.Error())
	}
	tests := []struct {
		name       string
		alpha2Code string
		want       V1Handler
	}{
		{
			name:       "Canada",
			alpha2Code: "CA",
			want: V1Handler{
				Count: 3,
				Data: models.Cities{
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
		},
		{
			name:       "USA",
			alpha2Code: "US",
			want: V1Handler{
				Count: 2,
				Data: models.Cities{
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
		},
	}
	for _, tt := range tests {
		res := as.JSON(fmt.Sprintf("/v1/countries/%s/cities/", tt.alpha2Code)).Get()
		as.Equal(200, res.Code)
		wStr, err := json.Marshal(tt.want)
		if err != nil {
			as.FailNow(err.Error())
		}
		as.Equal(string(wStr), strings.TrimSpace(string(res.Body.Bytes())))
	}
}

func (as *ActionSuite) Test_subdivisionCities() {
	as.LoadFixture("cities")
	testTime, err := time.Parse("2006-01-02 03:04:05", "2019-01-03 08:52:06")
	if err != nil {
		as.FailNow("could not generate test timestamp error: %s", err.Error())
	}
	tests := []struct {
		name            string
		subdivisionCode string
		want            V1Handler
	}{
		{
			name:            "Canada",
			subdivisionCode: "CA-QC",
			want: V1Handler{
				Count: 2,
				Data: models.Cities{
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
		},
		{
			name:            "USA",
			subdivisionCode: "US-NY",
			want: V1Handler{
				Count: 2,
				Data: models.Cities{
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
		},
	}
	for _, tt := range tests {
		res := as.JSON(fmt.Sprintf("/v1/subdivisions/%s/cities/", tt.subdivisionCode)).Get()
		as.Equal(200, res.Code)
		wStr, err := json.Marshal(tt.want)
		if err != nil {
			as.FailNow(err.Error())
		}
		as.Equal(string(wStr), strings.TrimSpace(string(res.Body.Bytes())))
	}
}

func (as *ActionSuite) Test_500_error() {
	res := as.JSON("/generate-error/").Get()
	as.Equal(500, res.Code)
	as.Contains(res.Body.String(), "(╯°□°）╯︵ ┻━┻  unknown error")
}

func (as *ActionSuite) Test_404_notfound() {
	res := as.JSON("/badrequest/").Get()
	as.Equal(404, res.Code)
	as.Contains(res.Body.String(), "「(°ヘ°) resource not found")
}
