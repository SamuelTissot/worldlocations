package actions

import (
	"fmt"
	"github.com/gobuffalo/buffalo"
	"regexp"
	"strconv"
	"strings"
	"worldlocations/models"
)

const maxPageCount int = 30000

type query func(c buffalo.Context) (models.Model, error)

type Page struct {
	Current int    `json:"current,omitempty"`
	Next    string `json:"next,omitempty"`
}

type V1Handler struct {
	query query       `json:"-"`
	Count int         `json:"count"`
	Data  interface{} `json:"data"`
	Page  *Page       `json:"page,omitempty"`
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

	//paginate if response is bigger than maxPageCount
	if h.Count > maxPageCount {
		p := c.Param("p")
		if p == "" {
			p = "1"
		}
		pn, err := strconv.Atoi(p)
		if err != nil {
			return err
		}
		var isMore bool
		res, isMore = res.Paginate(pn, maxPageCount)
		h.Page = &Page{Current: pn}
		if isMore {
			scheme := "http"
			if strings.Contains(c.Request().Proto, "https") {
				scheme = "https"
			}
			//remove previous pagination
			var re = regexp.MustCompile(`\?p=\d+`)
			url := re.ReplaceAll([]byte(c.Request().URL.String()), []byte(""))
			fmt.Println(url)
			h.Page.Next = fmt.Sprintf("%s://%s%s?p=%d", scheme, c.Request().Host, string(url), pn+1)
		}
	}

	h.Data = res

	respCode := 200
	if h.Count == 0 {
		respCode = 204
	}
	return c.Render(respCode, r.JSON(h))
}
