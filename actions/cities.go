package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/pkg/errors"
	"worldlocations/models"
)

type Cities struct {
	Count int `json:"count"`
	Data  *models.Cities
}

func (cs Cities) List(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	cm := &models.Cities{}
	if err := tx.All(cm); err != nil {
		return errors.WithStack(err)
	}

	res := Cities{
		Count: len(*cm),
		Data:  cm,
	}

	return c.Render(200, r.JSON(res))
}

func (cs Cities) Show(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	cm := &models.City{}
	if err := tx.Find(cm, c.Param("id")); err != nil {
		return c.Error(404, err)
	}

	res := Cities{
		Count: 1,
		Data:  &models.Cities{*cm},
	}

	return c.Render(200, r.JSON(res))
}

func (cs Cities) CountryCities(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	cm := &models.Cities{}
	if err := tx.Where("alpha_2_code = (?)", c.Param("alpha_2_code")).All(cm); err != nil {
		return c.Error(404, err)
	}

	res := Cities{
		Count: len(*cm),
		Data:  cm,
	}

	return c.Render(200, r.JSON(res))
}

func (cs Cities) SubdivisionCities(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	cm := &models.Cities{}
	if err := tx.Where("subdivision_code = (?)", c.Param("subdivision_code")).All(cm); err != nil {
		return c.Error(404, err)
	}

	res := Cities{
		Count: len(*cm),
		Data:  cm,
	}

	return c.Render(200, r.JSON(res))
}
