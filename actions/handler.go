package actions

import (
	"github.com/gobuffalo/buffalo"
)

type Countable interface {
	Count() int
}

type query func(c buffalo.Context) (Countable, error)

type V1Handler struct {
	query query       `json:"-"`
	Count int         `json:"count"`
	Data  interface{} `json:"data"`
}

func v1Handler(q query) func(c buffalo.Context) error {
	h := &V1Handler{query: q}
	return h.do
}

func (h *V1Handler) do(c buffalo.Context) error {
	res, err := h.query(c)
	if err != nil {
		return err
	}
	h.Count = res.Count()
	h.Data = res
	return c.Render(200, r.JSON(h))
}
