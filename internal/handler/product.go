package handler

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/service"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/service/default"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/request"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/response"
	"net/http"
	"strconv"
)

// NewProductDefault is a function that returns a new instance of ProductDefault
func NewProductDefault(sv *_default.ProductDefault) *ProductDefault {
	return &ProductDefault{sv: sv}
}

// ProductDefault is a struct with methods that represent handlers for Products
type ProductDefault struct {
	// sv is the service that will be used by the handler
	sv *_default.ProductDefault
}

// GetProducts GetAll is a method that returns a handler for the route GET /products
func (h *ProductDefault) GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// request
	// ...
	// process
	// - get all Products
	v, err := h.sv.RetrieveAll()
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusNotFound))
		return
	}
	_ = render.Render(w, r, response.NewResponse(v, http.StatusOK))
}

// PostProduct is a method that returns a handler for the route CREATE /product/{ID}
func (h *ProductDefault) PostProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data := &request.ProductRequest{}

	if err := render.Bind(r, data); err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusBadRequest))
	}

	product := models.NewProduct(
		0,
		*data.ProductCode, *data.Description,
		*data.Width,
		*data.Height,
		*data.Length,
		*data.NetWeight,
		*data.ExpirationRate,
		*data.RecommendedFreezingTemperature,
		*data.FreezingRate,
		*data.ProductTypeId,
		data.SellerId,
	)
	createdProduct, errService := h.sv.Register(*product)
	if errService != nil {
		_ = render.Render(w, r, response.NewErrorResponse(errService.Error(), http.StatusBadRequest))
		return
	}
	_ = render.Render(w, r, response.NewResponse(createdProduct, http.StatusCreated))
}

// GetProduct is a method that returns a handler for the route GET /product/{ID}
func (h *ProductDefault) GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, errConverter := strconv.Atoi(chi.URLParam(r, "id"))
	if errConverter != nil {
		_ = render.Render(w, r, response.NewErrorResponse(errConverter.Error(), http.StatusBadRequest))
		return
	}
	p, errServiceFindById := h.sv.Retrieve(id)
	if errServiceFindById != nil {
		_ = render.Render(w, r, response.NewErrorResponse(errServiceFindById.Error(), http.StatusNotFound))
		return
	}
	_ = render.Render(w, r, response.NewResponse(p, http.StatusOK))
}

// PatchProduct handles PATCH requests to partially update a product.
func (h *ProductDefault) PatchProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// 1. Get the product ID from the URL and handle conversion errors.
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse("Invalid ID format", http.StatusBadRequest))
		return
	}

	// 2. Decode the JSON body into a map, not a full struct.
	var fields map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&fields)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse("Invalid request body", http.StatusBadRequest))
		return
	}

	// 3. Call the service with the ID and the map of fields.
	updatedProduct, err := h.sv.PartialModify(id, fields)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusNotFound))
		return
	}

	// 4. Render the successful response with the updated product.
	_ = render.Render(w, r, response.NewResponse(updatedProduct, http.StatusOK))
}

func (h *ProductDefault) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, errConverter := strconv.Atoi(chi.URLParam(r, "id"))

	if errConverter != nil {
		_ = render.Render(w, r, response.NewErrorResponse(errConverter.Error(), http.StatusBadRequest))
		return
	}

	errServiceDelete := h.sv.Remove(id)

	if errServiceDelete != nil {
		_ = render.Render(w, r, response.NewErrorResponse(errServiceDelete.Error(), http.StatusNotFound))
		return
	}
	_ = render.Render(w, r, response.NewResponse("product Deleted", http.StatusNoContent))
}

func (h *ProductDefault) GetProductReport(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	idParam := r.URL.Query().Get("id")
	if idParam == "" {

		value, err := h.sv.RetrieveRecordsCount()
		if err != nil {
			_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusInternalServerError))
			return
		}
		_ = render.Render(w, r, response.NewResponse(value, http.StatusOK))
		return
	}
	id, errConverter := strconv.Atoi(idParam)
	if errConverter != nil {
		_ = render.Render(w, r, response.NewErrorResponse(errConverter.Error(), http.StatusBadRequest))
		return
	}

	value, err := h.sv.RetrieveRecordsCountByProductId(id)
	if err != nil {
		if errors.Is(err, service.ErrProductNotFound) {
			_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusNotFound))
			return
		}
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusInternalServerError))
		return
	}
	_ = render.Render(w, r, response.NewResponse(value, http.StatusOK))
}
