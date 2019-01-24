package actions

import (
	"worldlocations/models"
)

func (as *ActionSuite) Test_Countries_List() {
	as.LoadFixture("country codes exist")
	res := as.JSON("/v1/countries").Get()

	as.Equal(200, res.Code)

	body := res.Body.String()
	var countryCodes models.CountryCodes
	as.NoError(as.DB.All(&countryCodes))
	for _, c := range countryCodes {
		as.Contains(body, c.Alpha2Code)
	}
}
