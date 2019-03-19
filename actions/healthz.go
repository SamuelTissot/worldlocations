package actions

import "github.com/gobuffalo/buffalo"

func HealthzHandler(c buffalo.Context) error {
	return c.Render(200, r.JSON(map[string]string{"message": "Health OK"}))
}
