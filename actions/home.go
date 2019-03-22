package actions

import (
	"errors"
	"github.com/gobuffalo/buffalo"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	return c.Render(200, r.JSON(map[string]string{
		"message":       "Welcome to World Locations I/O",
		"github":        "https://github.com/SamuelTissot/worldlocations",
		"documentation": "https://documenter.getpostman.com/view/6999284/S17qTpUm",
	}))
}

func ErrorExample(c buffalo.Context) error {
	return errors.New("test error")
}
