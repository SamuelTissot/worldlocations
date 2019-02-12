package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/pkg/errors"
	"worldlocations/models"
)

type Languages struct {
	Count int
	Data  *models.LanguageCodes
}

func (l Languages) List(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	lcs := &models.LanguageCodes{}

	if err := tx.All(lcs); err != nil {
		return errors.WithStack(err)
	}

	resp := Languages{
		Count: len(*lcs),
		Data:  lcs,
	}

	return c.Render(200, r.JSON(resp))
}

func (l Languages) Show(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	lc := &models.LanguageCode{}

	if err := tx.Where("language_alpha_2_code = (?)", c.Param("language_alpha_2_code")).First(lc); err != nil {
		return errors.WithStack(err)
	}

	lcResp := Languages{
		Count: 1,
		Data:  &models.LanguageCodes{*lc},
	}

	return c.Render(200, r.Auto(c, lcResp))
}
