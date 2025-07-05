package handler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/service/default"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
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

// GetAllProducts GetAll is a method that returns a handler for the route GET /products
func (h *ProductDefault) GetAllProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...
		// process
		// - get all Products
		v, err := h.sv.FindAll()
		if err != nil {
			_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusNotFound))
			return
		}
		_ = render.Render(w, r, response.NewResponse(v, http.StatusOK))
		return
	}
}

// CreateProduct is a method that returns a handler for the route CREATE /product/{ID}
func (h *ProductDefault) CreateProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body models.Product
		_ = json.NewDecoder(r.Body).Decode(&body)

		product, errService := h.sv.Create(body)
		if errService != nil {
			_ = render.Render(w, r, response.NewErrorResponse(errService.Error(), http.StatusBadRequest))
			return
		}
		_ = render.Render(w, r, response.NewResponse(product, http.StatusCreated))
		return
	}
}

// FindByIDProduct is a method that returns a handler for the route GET /product/{ID}
func (h *ProductDefault) FindByIDProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, errConverter := strconv.Atoi(chi.URLParam(r, "ID"))
		if errConverter != nil {
			_ = render.Render(w, r, response.NewErrorResponse(errConverter.Error(), http.StatusBadRequest))
			return
		}
		p, errServiceFindById := h.sv.FindByID(id)
		if errServiceFindById != nil {
			_ = render.Render(w, r, response.NewErrorResponse(errServiceFindById.Error(), http.StatusNotFound))
			return
		}
		_ = render.Render(w, r, response.NewResponse(p, http.StatusOK))
		return
	}
}

func (h *ProductDefault) UpdateProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, errConverter := strconv.Atoi(chi.URLParam(r, "ID"))
		var body models.Product
		_ = json.NewDecoder(r.Body).Decode(&body)

		if errConverter != nil {
			_ = render.Render(w, r, response.NewErrorResponse(errConverter.Error(), http.StatusBadRequest))
			return
		}

		p, errServiceUpdateProduct := h.sv.UpdatePartiallyV2(id, body)

		if errServiceUpdateProduct != nil {
			_ = render.Render(w, r, response.NewErrorResponse(errServiceUpdateProduct.Error(), http.StatusNotFound))
			return
		}
		_ = render.Render(w, r, response.NewResponse(p, http.StatusOK))
		return
	}
}
func (h *ProductDefault) DeleteProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, errConverter := strconv.Atoi(chi.URLParam(r, "ID"))

		if errConverter != nil {
			_ = render.Render(w, r, response.NewErrorResponse(errConverter.Error(), http.StatusBadRequest))
			return
		}

		errServiceDelete := h.sv.Delete(id)

		if errServiceDelete != nil {
			_ = render.Render(w, r, response.NewErrorResponse(errServiceDelete.Error(), http.StatusNotFound))
			return
		}
		_ = render.Render(w, r, response.NewResponse("product Deleted", http.StatusOK))
		return
	}
}
