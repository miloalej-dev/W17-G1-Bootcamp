package handler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/service/product"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/response"

	"net/http"
	"strconv"
)

// NewProductDefault is a function that returns a new instance of ProductDefault
func NewProductDefault(sv productService.ProductService) *ProductDefault {
	return &ProductDefault{sv: sv}
}

// ProductDefault is a struct with methods that represent handlers for Products
type ProductDefault struct {
	// sv is the service that will be used by the handler
	sv productService.ProductService
}

// GetAll is a method that returns a handler for the route GET /products
func (h *ProductDefault) GetAllProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...
		// process
		// - get all Products
		v, err := h.sv.FindAll()
		if err != nil {
			render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusNotFound))
			return
		}

		// response
		data := make(map[int]models.Product)
		for key, value := range v {
			data[key] = models.Product{
				ID:                             value.ID,
				ProductCode:                    value.ProductCode,
				Description:                    value.Description,
				Width:                          value.Width,
				Height:                         value.Height,
				Length:                         value.Length,
				NetWeight:                      value.NetWeight,
				ExpirationRate:                 value.ExpirationRate,
				RecommendedFreezingTemperature: value.RecommendedFreezingTemperature,
				FreezingRate:                   value.FreezingRate,
				ProductTypeID:                  value.ProductTypeID,
				SellerID:                       value.SellerID,
			}
		}
		render.Render(w, r, response.NewResponse(data, http.StatusOK))
	}
}

// Create is a method that returns a handler for the route CREATE /product/{ID}
func (h *ProductDefault) CreateProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var Body models.Product
		json.NewDecoder(r.Body).Decode(&Body)

		errService := h.sv.Create(Body)
		if errService != nil {
			render.Render(w, r, response.NewErrorResponse(errService.Error(), http.StatusBadRequest))

			return
		}
		render.Render(w, r, response.NewResponse(Body, http.StatusCreated))
		return
	}
}

// FindByID
func (h *ProductDefault) FindyByIDProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ID, errConverter := strconv.Atoi(chi.URLParam(r, "ID"))

		if errConverter != nil {
			render.Render(w, r, response.NewErrorResponse(errConverter.Error(), http.StatusBadRequest))

			return
		}

		p, errServiceFindByID := h.sv.FindByID(ID)
		if errServiceFindByID != nil {
			render.Render(w, r, response.NewErrorResponse(errServiceFindByID.Error(), http.StatusNotFound))
			return
		}
		render.Render(w, r, response.NewResponse(p, http.StatusOK))

		return
	}
}

func (h *ProductDefault) UpdateProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ID, errConverter := strconv.Atoi(chi.URLParam(r, "ID"))
		var Body models.Product
		json.NewDecoder(r.Body).Decode(&Body)

		if errConverter != nil {
			render.Render(w, r, response.NewErrorResponse(errConverter.Error(), http.StatusBadRequest))
			return
		}

		p, errServiceUpdateProduct := h.sv.UpdateProduct(ID, Body)

		if errServiceUpdateProduct != nil {
			render.Render(w, r, response.NewErrorResponse(errServiceUpdateProduct.Error(), http.StatusNotFound))
			return
		}

		render.Render(w, r, response.NewResponse(p, http.StatusOK))

		return
	}
}
func (h *ProductDefault) DeleteProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ID, errConverter := strconv.Atoi(chi.URLParam(r, "ID"))

		if errConverter != nil {
			render.Render(w, r, response.NewErrorResponse(errConverter.Error(), http.StatusBadRequest))
			return
		}

		errServiceDelete := h.sv.Delete(ID)

		if errServiceDelete != nil {
			render.Render(w, r, response.NewErrorResponse(errServiceDelete.Error(), http.StatusNotFound))
			return
		}

		render.Render(w, r, response.NewResponse("El producto ha sido borrado", http.StatusOK))

		return
	}
}
