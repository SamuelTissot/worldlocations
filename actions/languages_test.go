package actions

import (
	"encoding/json"
	"fmt"
	"time"
	"worldlocations/models"
)

func (as *ActionSuite) TestLanguages_List() {
	as.LoadFixture("language_codes")
	res := as.JSON("/v1/languages").Get()

	as.Equal(200, res.Code)

	body := res.Body.String()
	var languageCodes models.LanguageCodes
	as.NoError(as.DB.All(&languageCodes))
	for _, c := range languageCodes {
		as.Contains(body, c.LanguageAlpha2Code)
	}
}

func (as *ActionSuite) TestLanguages_Show() {
	as.LoadFixture("language_codes")
	testTime, err := time.Parse("2006-01-02 03:04:05", "2019-01-03 08:52:06")
	if err != nil {
		as.FailNow("could not generate test timestamp error: %s", err.Error())
	}

	tests := []struct {
		name               string
		LanguageAlpha2Code string
		want               models.LanguageCodes
	}{
		{
			name:               "Spanish",
			LanguageAlpha2Code: "es",
			want: models.LanguageCodes{
				{
					LanguageAlpha2Code: "es",
					LanguageAlpha3Code: "spa",
					CreatedAt:          testTime,
					UpdatedAt:          testTime,
				},
			},
		},
		{
			name:               "French",
			LanguageAlpha2Code: "fr",
			want: models.LanguageCodes{
				{
					LanguageAlpha2Code: "fr",
					LanguageAlpha3Code: "fra",
					CreatedAt:          testTime,
					UpdatedAt:          testTime,
				},
			},
		},
	}
	for _, tt := range tests {

		res := as.JSON(fmt.Sprintf("/v1/languages/%s/", tt.LanguageAlpha2Code)).Get()
		as.Equal(200, res.Code)

		ls := &Languages{}
		err := json.Unmarshal(res.Body.Bytes(), ls)
		if err != nil {
			as.FailNow(err.Error())
		}
		as.Equal(tt.want, *ls.Data)
	}
}
