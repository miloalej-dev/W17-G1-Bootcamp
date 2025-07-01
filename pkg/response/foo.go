package response

import (
	"github.com/go-chi/render"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
	"net/http"
)

type FooResponse struct {
	Id          int     `json:"id"`
	Name        *string `json:"name"`
	Description string  `json:"description"`
	Elapse      int     `json:"elapse"`
}

func (re *FooResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before marshalling to JSON

	// Always add a 100 elapse
	re.Elapse = 100
	return nil
}

func NewFooResponse(foo *models.Foo) *FooResponse {
	resp := &FooResponse{
		Id:          foo.ID,
		Name:        &foo.Name,
		Description: foo.Description,
		Elapse:      0,
	}

	return resp
}

func NewFooListResponse(foos []*models.Foo) []render.Renderer {
	var list []render.Renderer
	for _, foo := range foos {
		list = append(list, NewFooResponse(foo))
	}
	return list
}
