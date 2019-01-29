package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/pkg/errors"
	"worldlocations/models"
)

type Countries struct {
	Count int                  `json:"count"`
	Data  *models.CountryCodes `json:"data"`
}

func (countries Countries) List(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	countryCodes := &models.CountryCodes{}

	// Retrieve all CountryCodes from the DB
	if err := tx.All(countryCodes); err != nil {
		return errors.WithStack(err)
	}
	cts := Countries{
		Count: len(*countryCodes),
		Data:  countryCodes,
	}

	return c.Render(200, r.JSON(cts))
}

func (countries Countries) Show(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	cts := &models.CountryCode{}

	if err := tx.Where("alpha_2_code = (?)", c.Param("alpha_2_code")).First(cts); err != nil {
		return c.Error(404, err)
	}

	return c.Render(200, r.Auto(c, cts))
}
