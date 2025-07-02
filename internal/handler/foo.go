package handler

import (
	"github.com/go-chi/render"
	//"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/request"
	//"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/response"
	"net/http"
)

// Structure that represents a foo handler
type FooHandler struct {
}

//// NewFooHandler is a function that returns a new instance of the foo handler
//func NewFooHandler() *FooHandler {
	//return &FooHandler{}
//}

//// FooHandler methods goes here
//func (h *FooHandler) GetAllFoo(w http.ResponseWriter, r *http.Request) {

	//foos := []*models.Foo{
		//{
			//ID:          1,
			//Name:        "Foo1",
			//Description: "Kind of foo",
		//},
		//{
			//ID:          2,
			//Name:        "Foo2",
			//Description: "Kind of foo",
		//},
		//{
			//ID:          3,
			//Name:        "Foo3",
			//Description: "Kind of foo",
		//},
	//}

	//render.Status(r, http.StatusOK)
	//render.RenderList(w, r, response.NewFooListResponse(foos))

//}

func (h *FooHandler) PostFoo(w http.ResponseWriter, r *http.Request) {
	data := &request.FooRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Status(r, http.StatusBadRequest)
		return
	}

	//render.Status(r, http.StatusCreated)

	//render.Render(w, r, response.NewFooResponse(&models.Foo{
		//ID:          10,
		//Name:        *data.Name,
		//Description: data.Description,
	//},
	//))
}
