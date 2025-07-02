package productHandler

import (
	"encoding/json"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/service/product"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
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
func (h *ProductDefault) GetAll() http.HandlerFunc {
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
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

// Create is a method that returns a handler for the route CREATE /product/{ID}
func (h *ProductDefault) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var Body models.Product
		json.NewDecoder(r.Body).Decode(&Body)

		errService := h.sv.Create(Body)
		if errService != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": errService.Error(),
			})
			return
		}

		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "Producto Creado",
			"data":    Body,
		})
		return
	}
}

// FindByID
func (h *ProductDefault) FindyByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ID, errConverter := strconv.Atoi(chi.URLParam(r, "ID"))

		if errConverter != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": errConverter.Error(),
			})
			return
		}

		p, errServiceFindByID := h.sv.FindByID(ID)
		if errServiceFindByID != nil {
			response.JSON(w, http.StatusNotFound, map[string]any{
				"message": errServiceFindByID.Error(),
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
		ID, errConverter := strconv.Atoi(chi.URLParam(r, "ID"))
		var Body models.Product
		json.NewDecoder(r.Body).Decode(&Body)

		if errConverter != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": errConverter.Error(),
			})
			return
		}

		p, errServiceUpdateProduct := h.sv.UpdateProduct(ID, Body)

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
func (h *ProductDefault) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ID, errConverter := strconv.Atoi(chi.URLParam(r, "ID"))

		if errConverter != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": errConverter.Error(),
			})
			return
		}

		errServiceDelete := h.sv.Delete(ID)

		if errServiceDelete != nil {
			response.JSON(w, http.StatusNotFound, map[string]any{
				"message": errServiceDelete.Error(),
			})
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "Producto Borrado",
		})
		return
	}
}
