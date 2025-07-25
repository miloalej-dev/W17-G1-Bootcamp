package handler

import (
	"github.com/go-chi/render"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/service/default"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/request"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/response"
	"net/http"
)

// NewProductDefault is a function that returns a new instance of ProductDefault
func NewProductBatchDefault(sv *_default.ProductBatchDefault) *ProductBatchDefault {
	return &ProductBatchDefault{sv: sv}
}

// ProductDefault is a struct with methods that represent handlers for Products
type ProductBatchDefault struct {
	// sv is the service that will be used by the handler
	sv *_default.ProductBatchDefault
}

// PostProduct is a method that returns a handler for the route CREATE /product/{ID}
func (h *ProductBatchDefault) PostProductBatch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data := &request.ProductBatchRequest{}

	if err := render.Bind(r, data); err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusUnprocessableEntity))
	}

	product := models.NewProductBatch(
		0,
		*data.BatchNumber,
		*data.CurrentQuantity,
		*data.CurrentTemperature,
		*data.DueDate,
		*data.InitialQuantity,
		*data.ManufacturingDate,
		*data.ManufacturingHour,
		*data.MinimumTemperature,
		*data.SectionId,
		*data.ProductId,
	)
	createdProductBatch, errService := h.sv.Register(product)
	if errService != nil {
		_ = render.Render(w, r, response.NewErrorResponse(errService.Error(), http.StatusConflict))
		return
	}
	_ = render.Render(w, r, response.NewResponse(createdProductBatch, http.StatusCreated))
}
