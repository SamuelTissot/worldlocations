package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/pkg/errors"
	"worldlocations/models"
)

type Subdivisions struct {
	Count int                      `json:"count"`
	Data  *models.SubdivisionCodes `json:"data"`
}

func (scs Subdivisions) List(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	subCodes := &models.SubdivisionCodes{}
	if err := tx.All(subCodes); err != nil {
		return errors.WithStack(err)
	}
	subCodeRes := Subdivisions{
		Count: len(*subCodes),
		Data:  subCodes,
	}

	return c.Render(200, r.JSON(subCodeRes))
}

func (scs Subdivisions) CountrySubdivisions(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	scsq := &models.SubdivisionCodes{}
	if err := tx.Where("alpha_2_code = (?)", c.Param("alpha_2_code")).All(scsq); err != nil {
		return c.Error(404, err)
	}

	subdivisionsRes := Subdivisions{
		Count: len(*scsq),
		Data:  scsq,
	}

	return c.Render(200, r.Auto(c, subdivisionsRes))
}
