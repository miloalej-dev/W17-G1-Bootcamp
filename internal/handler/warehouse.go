package handler

import (
	"net/http"
	"strconv"
	"encoding/json"

	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/service/warehouse"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/response"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/request"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// WarehouseDefault is a struct with methods that represent handlers for warehouses
type WarehouseDefault struct {
	// sv is the service that will be used by the handler
	sv service.WarehouseService
}

// NewWarehouseDefault is a function that returns a new instance of WarehouseDefault
func NewWarehouseDefault(sv service.WarehouseService) *WarehouseDefault {
	return &WarehouseDefault{sv: sv}
}

func (h *WarehouseDefault) FindAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		warehouses, err := h.sv.FindAll()
		if err != nil {
			http.Error(w, "Failed to retrieve sellers", http.StatusInternalServerError)
			return
		}

		warehousesJson := make([]models.WarehouseDoc, 0)
		for _, warehouse := range warehouses {
			warehousesJson = append(warehousesJson, models.WarehouseDoc{
				ID:                 warehouse.ID,
				Code:               warehouse.Code,
				Address:            warehouse.Address,
				Telephone:          warehouse.Telephone,
				MinimumCapacity:    warehouse.MinimumCapacity,
				MinimumTemperature: warehouse.MinimumTemperature,
			})
		}

		render.Render(w, r, response.NewResponse(warehousesJson, http.StatusOK))
	}
}

func (h *WarehouseDefault) FindById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idRequest := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idRequest)
		if err != nil {
			render.Render(w, r, response.NewErrorResponse("Invalid ID", http.StatusBadRequest))
			return
		}

		warehouse, err := h.sv.FindById(id)
		if err != nil {
			render.Render(w, r, response.NewErrorResponse("Internal error", http.StatusInternalServerError))
			return
		}

		warehouseJson := models.WarehouseDoc{
			ID:                 warehouse.ID,
			Code:               warehouse.Code,
			Address:            warehouse.Address,
			Telephone:          warehouse.Telephone,
			MinimumCapacity:    warehouse.MinimumCapacity,
			MinimumTemperature: warehouse.MinimumTemperature,
		}

		render.Render(w, r, response.NewResponse(warehouseJson, http.StatusOK))
	}
}

func (h *WarehouseDefault) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		warehouseJson := &request.WarehouseRequest{}
		if err := render.Bind(r, warehouseJson); err != nil {
			render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusBadRequest))
			return
		}

		warehouse := models.Warehouse {
			WarehouseAttributes: models.WarehouseAttributes {
				Code: *warehouseJson.Code,
				Address: *warehouseJson.Address,
				Telephone: *warehouseJson.Telephone,
				MinimumCapacity: *warehouseJson.MinimumCapacity,
				MinimumTemperature: *warehouseJson.MinimumTemperature,
			},
		}

		warehouseResponse, err := h.sv.Create(warehouse)
		if err != nil {
			render.Render(w, r, response.NewErrorResponse("Internal error", http.StatusInternalServerError))
			return
		}

		render.Render(w, r, response.NewResponse(warehouseResponse, http.StatusCreated))
	}
}

func (h *WarehouseDefault) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idRequest := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idRequest)
		if err != nil {
			render.Render(w, r, response.NewErrorResponse("Invalid ID", http.StatusBadRequest))
			return
		}

		warehouseJson := &request.WarehouseRequest{}
		if err := json.NewDecoder(r.Body).Decode(&warehouseJson); err != nil {
			render.Render(w, r, response.NewErrorResponse("Internal error", http.StatusInternalServerError))
			return
		}

		warehouse, err := h.sv.FindById(id)
		if err != nil {
			render.Render(w, r, response.NewErrorResponse("Internal error", http.StatusInternalServerError))
			return
		}

		if warehouseJson.Code != nil &&
			*warehouseJson.Code != warehouse.Code {
			warehouse.Code = *warehouseJson.Code
		}
		if warehouseJson.Address != nil &&
			*warehouseJson.Address != warehouse.Address {
			warehouse.Address = *warehouseJson.Address
		}
		if warehouseJson.Telephone != nil &&
			*warehouseJson.Telephone != warehouse.Telephone {
			warehouse.Telephone = *warehouseJson.Telephone
		}
		if warehouseJson.MinimumCapacity != nil &&
			*warehouseJson.MinimumCapacity != warehouse.MinimumCapacity {
			warehouse.MinimumCapacity = *warehouseJson.MinimumCapacity
		}
		if warehouseJson.MinimumTemperature != nil &&
			*warehouseJson.MinimumTemperature != warehouse.MinimumTemperature {
			warehouse.MinimumTemperature = *warehouseJson.MinimumTemperature
		}

		warehouseResponse,err := h.sv.Update(warehouse)
		if err != nil {
			render.Render(w, r, response.NewErrorResponse("Internal error", http.StatusInternalServerError))
			return
		}

		render.Render(w, r, response.NewResponse(warehouseResponse, http.StatusOK))
	}
}

func (h *WarehouseDefault) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idRequest := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idRequest)
		if err != nil {
			render.Render(w, r, response.NewErrorResponse("Invalid ID", http.StatusBadRequest))
			return
		}

		err = h.sv.Delete(id)
		if err != nil {
			render.Render(w, r, response.NewErrorResponse("Internal error", http.StatusInternalServerError))
			return
		}

		render.Render(w, r, response.NewResponse(nil, http.StatusNoContent))
	}
}
