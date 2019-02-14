package actions

import (
	"encoding/json"
	"fmt"
	"time"
	"worldlocations/models"
)

func (as *ActionSuite) TestSubdivisionNames_List() {
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

func (as *ActionSuite) TestSubdivisionNames_Show() {
	as.LoadFixture("subdivision_names")
	testTime, err := time.Parse("2006-01-02 03:04:05", "2019-01-03 08:52:06")
	if err != nil {
		as.FailNow("could not generate test timestamp error: %s", err.Error())
	}

	tests := []struct {
		name               string
		LanguageAlpha2Code string
		want               models.SubdivisionNames
	}{
		{
			name:               "Quebec",
			LanguageAlpha2Code: "CA-QC",
			want: models.SubdivisionNames{
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
	}

	for _, tt := range tests {

		res := as.JSON(fmt.Sprintf("/v1/subdivisions/%s/names/", tt.LanguageAlpha2Code)).Get()
		as.Equal(200, res.Code)

		ls := &SubdivisionNames{}
		err := json.Unmarshal(res.Body.Bytes(), ls)
		if err != nil {
			as.FailNow(err.Error())
		}
		as.Equal(tt.want, *ls.Data)
	}
}
