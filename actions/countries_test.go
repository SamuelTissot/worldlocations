package actions

import (
	"worldlocations/models"
)

func (as *ActionSuite) Test_Countries_List() {
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

func (as *ActionSuite) Test_Countries_Show() {
	as.LoadFixture("country_codes")
	res := as.JSON("/v1/countries/ca").Get()

	as.Equal(200, res.Code)
	body := res.Body.String()
	var countryCode models.CountryCode
	as.NoError(as.DB.Where("alpha_2_code = (?)", "CA").First(&countryCode))
	as.Contains(body, countryCode.Alpha2Code)
}
