package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/pkg/errors"
	"worldlocations/models"
)

type Inquiry struct{}

// fetch executes the query function (fn) and returns it's value
func (i *Inquiry) fetch(c buffalo.Context, fn func(tx *pop.Connection) error) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}
	return fn(tx)
}

// All returns a query function that fetches all rows
func (i *Inquiry) all(model interface{}) func(tx *pop.Connection) error {
	return func(tx *pop.Connection) error {
		if err := tx.All(model); err != nil {
			return errors.WithStack(err)
		}
		return nil
	}
}

// where returns a query function that does an "SELECT .... WHERE ...."
func (i *Inquiry) where(model interface{}, stmt string, args ...interface{}) func(tx *pop.Connection) error {
	return func(tx *pop.Connection) error {
		if err := tx.Where(stmt, args...).All(model); err != nil {
			return errors.WithStack(err)
		}
		return nil
	}
}

/*
 *
 * COUNTRIES
 *
 */

// countryList lists all countries
func (i *Inquiry) countryList(c buffalo.Context) (Countable, error) {
	model := &models.CountryCodes{}
	if err := i.fetch(c, i.all(model)); err != nil {
		return nil, err
	}
	return model, nil
}

// countryShow returns the country with alpha_2_code
func (i *Inquiry) countryShow(c buffalo.Context) (Countable, error) {
	model := &models.CountryCodes{}
	if err := i.fetch(c, i.where(model, "alpha_2_code = (?)", c.Param("alpha_2_code"))); err != nil {
		return nil, err
	}
	return model, nil
}

/*
 *
 * COUNTRIES NAMES
 *
 */
// countriesNamesList list all the countries names variations
func (i *Inquiry) countriesNamesList(c buffalo.Context) (Countable, error) {
	model := &models.CountryNames{}
	if err := i.fetch(c, i.all(model)); err != nil {
		return nil, err
	}
	return model, nil
}

// countryNames returns all the name for a country code (alpha_2_code)
func (i *Inquiry) countryNames(c buffalo.Context) (Countable, error) {
	model := &models.CountryNames{}
	if err := i.fetch(c, i.where(model, "alpha_2_code = (?)", c.Param("alpha_2_code"))); err != nil {
		return nil, err
	}
	return model, nil
}

/*
 *
 * Subdivision_names
 *
 */

// subdivisionNamesList returns all subdivision names
func (i *Inquiry) subdivisionNamesList(c buffalo.Context) (Countable, error) {
	model := &models.SubdivisionNames{}
	if err := i.fetch(c, i.all(model)); err != nil {
		return nil, err
	}
	return model, nil
}

// subdivisionNamesShow returns a single subdivision name based on subdivision_code
func (i *Inquiry) subdivisionNamesShow(c buffalo.Context) (Countable, error) {
	model := &models.SubdivisionNames{}
	if err := i.fetch(c, i.where(model, "subdivision_code = (?)", c.Param("subdivision_code"))); err != nil {
		return nil, err
	}
	return model, nil
}

/*
 *
 * Subdivision
 *
 */

// subdivisionsList returns all subdivision
func (i *Inquiry) subdivisionList(c buffalo.Context) (Countable, error) {
	model := &models.SubdivisionCodes{}
	if err := i.fetch(c, i.all(model)); err != nil {
		return nil, err
	}
	return model, nil
}

// subdivisionShow returns a single subdivision for subdivision_code
func (i *Inquiry) subdivisionShow(c buffalo.Context) (Countable, error) {
	model := &models.SubdivisionCodes{}
	if err := i.fetch(c, i.where(model, "subdivision_code = (?)", c.Param("subdivision_code"))); err != nil {
		return nil, err
	}
	return model, nil
}

// subdivisionShow returns a single subdivision for subdivision_code
func (i *Inquiry) countrySubdivisions(c buffalo.Context) (Countable, error) {
	model := &models.SubdivisionCodes{}
	if err := i.fetch(c, i.where(model, "alpha_2_code = (?)", c.Param("alpha_2_code"))); err != nil {
		return nil, err
	}
	return model, nil
}

/*
 *
 * Languages
 *
 */
// languagesList lists all the languages
func (i *Inquiry) languagesList(c buffalo.Context) (Countable, error) {
	model := &models.LanguageCodes{}
	if err := i.fetch(c, i.all(model)); err != nil {
		return nil, err
	}
	return model, nil
}

// languagesShow returns languages for language_alpha_2_code
func (i *Inquiry) languagesShow(c buffalo.Context) (Countable, error) {
	model := &models.LanguageCodes{}
	if err := i.fetch(c, i.where(model, "language_alpha_2_code = (?)", c.Param("language_alpha_2_code"))); err != nil {
		return nil, err
	}
	return model, nil
}

/*
 *
 * Cities
 *
 */

// citiesList lists all the cities
func (i *Inquiry) citiesList(c buffalo.Context) (Countable, error) {
	model := &models.Cities{}
	if err := i.fetch(c, i.all(model)); err != nil {
		return nil, err
	}
	return model, nil
}

// citiesShow returns the city with id = ?
func (i *Inquiry) citiesShow(c buffalo.Context) (Countable, error) {
	model := &models.Cities{}
	if err := i.fetch(c, i.where(model, "id = (?)", c.Param("id"))); err != nil {
		return nil, err
	}
	return model, nil
}

// countryCities returns all the cities of a country
func (i *Inquiry) countryCities(c buffalo.Context) (Countable, error) {
	model := &models.Cities{}
	if err := i.fetch(c, i.where(model, "alpha_2_code = (?)", c.Param("alpha_2_code"))); err != nil {
		return nil, err
	}
	return model, nil
}

// subdivisionCities returns all the cities of a subdivision
func (i *Inquiry) subdivisionCities(c buffalo.Context) (Countable, error) {
	model := &models.Cities{}
	if err := i.fetch(c, i.where(model, "subdivision_code = (?)", c.Param("subdivision_code"))); err != nil {
		return nil, err
	}
	return model, nil
}
