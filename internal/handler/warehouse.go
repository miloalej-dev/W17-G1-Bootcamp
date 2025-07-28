package handler

import (
	"errors"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/service"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/request"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/response"

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

func (h *WarehouseDefault) GetWarehouses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	warehouses, err := h.sv.RetrieveAll()
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusInternalServerError))
		return
	}

	_ = render.Render(w, r, response.NewResponse(warehouses, http.StatusOK))
}

func (h *WarehouseDefault) GetWarehouse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idRequest := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idRequest)
	if err != nil || id < 1 {
		_ = render.Render(w, r, response.NewErrorResponse(ErrInvalidId.Error(), http.StatusBadRequest))
		return
	}

	warehouse, err := h.sv.Retrieve(id)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusNotFound))
		return
	}

	_ = render.Render(w, r, response.NewResponse(warehouse, http.StatusOK))
}

func (h *WarehouseDefault) PostWarehouse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	warehouseJson := &request.WarehouseRequest{}
	if err := render.Bind(r, warehouseJson); err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}

	warehouse := models.NewWarehouse(
		0, // placeholder, will be overwritten later
		*warehouseJson.WarehouseCode,
		*warehouseJson.Address,
		*warehouseJson.Telephone,
		*warehouseJson.MinimumCapacity,
		*warehouseJson.MinimumTemperature,
		*warehouseJson.LocalityId,
	)

	warehouseResponse, err := h.sv.Register(*warehouse)
	if err != nil {
		if errors.Is(err, repository.ErrLocalityNotFound) {
			_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusBadRequest))
			return
		}

		_ = render.Render(w, r, response.NewErrorResponse("internal error", http.StatusInternalServerError))
		return
	}

	_ = render.Render(w, r, response.NewResponse(warehouseResponse, http.StatusCreated))
}

// PutWarehouse handles PUT requests to update a warehouse
func (h *WarehouseDefault) PutWarehouse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id < 1 {
		_ = render.Render(w, r, response.NewErrorResponse(ErrInvalidId.Error(), http.StatusBadRequest))
		return
	}

	data := &request.WarehouseRequest{}

	err = render.Bind(r, data)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}

	warehouse := models.NewWarehouse(
		id, // placeholder, will be overwritten later
		*data.WarehouseCode,
		*data.Address,
		*data.Telephone,
		*data.MinimumCapacity,
		*data.MinimumTemperature,
		*data.LocalityId,
	)

	updatedWarehouse, err := h.sv.Modify(*warehouse)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusInternalServerError))
		return
	}

	_ = render.Render(w, r, response.NewResponse(updatedWarehouse, http.StatusOK))
}

func (h *WarehouseDefault) PatchWarehouse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idRequest := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idRequest)
	if err != nil || id < 1 {
		_ = render.Render(w, r, response.NewErrorResponse(ErrInvalidId.Error(), http.StatusBadRequest))
		return
	}

	var fields map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&fields)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(ErrUnexpectedJSON.Error(), http.StatusBadRequest))
		return
	}

	warehouseResponse, err := h.sv.PartialModify(id, fields)
	if err != nil {
		if errors.Is(err, repository.ErrEntityNotFound) {
			_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusBadRequest))
			return
		}

		_ = render.Render(w, r, response.NewErrorResponse("internal error", http.StatusInternalServerError))
		return
	}

	_ = render.Render(w, r, response.NewResponse(warehouseResponse, http.StatusOK))
}

func (h *WarehouseDefault) DeleteWarehouse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idRequest := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idRequest)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(ErrInvalidId.Error(), http.StatusBadRequest))
		return
	}

	err = h.sv.Remove(id)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusInternalServerError))
		return
	}

	_ = render.Render(w, r, response.NewResponse(nil, http.StatusNoContent))
}
