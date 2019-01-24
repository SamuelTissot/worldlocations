package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/pkg/errors"
	"worldlocations/models"
)

type CountriesResource struct {
	buffalo.Resource
}

type Countries struct {
	Count int                  `json:"count"`
	Data  *models.CountryCodes `json:"data"`
}

func (cr CountriesResource) List(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	countryCodes := &models.CountryCodes{}

	// Retrieve all CountryCodes from the DB
	if err := tx.All(countryCodes); err != nil {
		return errors.WithStack(err)
	}

	countries := Countries{
		Count: len(*countryCodes),
		Data:  countryCodes,
	}

	return c.Render(200, r.JSON(countries))
}

func (cr CountriesResource) Show(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	country := &models.CountryCode{}

	// To find the Widget the parameter widget_id is used.
	if err := tx.Find(country, c.Param("alpha_2_code")); err != nil {
		return c.Error(404, err)
	}

	return c.Render(200, r.Auto(c, country))
}
