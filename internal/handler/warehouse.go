package handler

import (
	"net/http"
	"strconv"
	"encoding/json"

	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/service/warehouse"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"

	"github.com/go-chi/render"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"

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

func (h *WarehouseDefault) GetAll() http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		warehouses,err := h.sv.GetAll()
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, nil)
			return
		}

		if len(warehouses) == 0 {
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, nil)
			return
		}

		warehousesJson := make(map[int]models.WarehouseDoc)
		for id,warehouse := range warehouses {
			warehousesJson[id] = models.WarehouseDoc {
				ID:	warehouse.ID,
				Code: warehouse.Code,
				Address: warehouse.Address,
				Telephone: warehouse.Telephone,
				MinimunCapacity: warehouse.MinimunCapacity,
				MinimumTemperature: warehouse.MinimumTemperature,
			}
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, warehousesJson)
	}
}

func (h *WarehouseDefault) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idRequest := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idRequest)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, "Invalid id")
			return
		}

		warehouse, err := h.sv.GetById(id)
		if err != nil {
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, err.Error())
			return
		}

		warehouseJson := models.WarehouseDoc {
			ID:	warehouse.ID,
			Code: warehouse.Code,
			Address: warehouse.Address,
			Telephone: warehouse.Telephone,
			MinimunCapacity: warehouse.MinimunCapacity,
			MinimumTemperature: warehouse.MinimumTemperature,
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, warehouseJson)
	}
}

func (h *WarehouseDefault) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var warehouseJson models.WarehouseDoc
		if err := json.NewDecoder(r.Body).Decode(&warehouseJson); err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, "Internal error")
			return
		}

		warehouse, err := h.sv.Create(warehouseJson)
		if err != nil {
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, err.Error())
			return
		}

		render.Status(r, http.StatusCreated)
		render.JSON(w, r, warehouse)
	}
}

func (h *WarehouseDefault) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var warehouseJson models.WarehouseDoc
		if err := json.NewDecoder(r.Body).Decode(&warehouseJson); err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, "Internal error")
			return
		}

		idRequest := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idRequest)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, "Invalid id")
			return
		}

		warehouse, err := h.sv.GetById(id)
		if err != nil {
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, err.Error())
			return
		}

		if warehouseJson.Code != "" && warehouseJson.Code != warehouse.Code {
			warehouse.Code = warehouseJson.Code
		}
		if warehouseJson.Address != "" && warehouseJson.Address != warehouse.Address {
			warehouse.Address = warehouseJson.Address
		}
		if warehouseJson.Telephone != "" && warehouseJson.Telephone != warehouse.Telephone {
			warehouse.Telephone = warehouseJson.Telephone
		}

		err = h.sv.Update(warehouse)
		if err != nil {
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, err.Error())
			return
		}

		render.Status(r, http.StatusCreated)
		render.JSON(w, r, warehouse)
	}
}

func (h *WarehouseDefault) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idRequest := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idRequest)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, "Invalid id")
			return
		}

		err = h.sv.Delete(id)
		if err != nil {
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, err.Error())
			return
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, nil)
	}
}
