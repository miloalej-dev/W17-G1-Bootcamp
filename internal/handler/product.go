package handler

import (
	"encoding/json"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/service/default"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
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
			response.JSON(w, http.StatusNotFound, nil)
			return
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    v,
		})
	}
}

// CreateProduct is a method that returns a handler for the route CREATE /product/{ID}
func (h *ProductDefault) CreateProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body models.Product
		json.NewDecoder(r.Body).Decode(&body)

		product, errService := h.sv.Create(body)
		if errService != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": errService.Error(),
			})
			return
		}

		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "Producto Creado",
			"data":    product,
		})
		return
	}
}

// FindByIDProduct is a method that returns a handler for the route GET /product/{ID}
func (h *ProductDefault) FindByIDProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id, errConverter := strconv.Atoi(chi.URLParam(r, "ID"))
		if errConverter != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": errConverter.Error(),
			})
			return
		}

		p, errServiceFindById := h.sv.FindByID(id)
		if errServiceFindById != nil {
			response.JSON(w, http.StatusNotFound, map[string]any{
				"message": errServiceFindById.Error(),
			})
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "Producto Creado",
			"data":    p,
		})
		return
	}
}

func (h *ProductDefault) UpdateProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, errConverter := strconv.Atoi(chi.URLParam(r, "ID"))
		var body models.Product
		json.NewDecoder(r.Body).Decode(&body)

		if errConverter != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": errConverter.Error(),
			})
			return
		}

		p, errServiceUpdateProduct := h.sv.UpdatePartiallyV2(id, body)

		if errServiceUpdateProduct != nil {
			response.JSON(w, http.StatusNotFound, map[string]any{
				"message": errServiceUpdateProduct.Error(),
			})
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "Producto Actualizado",
			"data":    p,
		})
		return
	}
}
func (h *ProductDefault) DeleteProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, errConverter := strconv.Atoi(chi.URLParam(r, "ID"))

		if errConverter != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": errConverter.Error(),
			})
			return
		}

		errServiceDelete := h.sv.Delete(id)

		if errServiceDelete != nil {
			response.JSON(w, http.StatusNotFound, map[string]any{
				"message": errServiceDelete.Error(),
			})
			return
		}

		response.JSON(w, http.StatusNoContent, map[string]any{
			"message": "Producto Borrado",
		})
		return
	}
}
