package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/pkg/errors"
	"worldlocations/models"
)

type SubdivisionNames struct {
	Count int                      `json:"count"`
	Data  *models.SubdivisionNames `json:"data"`
}

func (sns SubdivisionNames) List(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	subdivisionNames := &models.SubdivisionNames{}
	if err := tx.All(subdivisionNames); err != nil {
		return errors.WithStack(err)
	}
	names := SubdivisionNames{
		Count: len(*subdivisionNames),
		Data:  subdivisionNames,
	}

	return c.Render(200, r.JSON(names))
}

func (sns SubdivisionNames) Show(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	subdivisionNames := &models.SubdivisionNames{}
	if err := tx.Where("subdivision_code = (?)", c.Param("subdivision_code")).All(subdivisionNames); err != nil {
		return c.Error(404, err)
	}

	subdivisionsRes := SubdivisionNames{
		Count: len(*subdivisionNames),
		Data:  subdivisionNames,
	}

	return c.Render(200, r.Auto(c, subdivisionsRes))
}
