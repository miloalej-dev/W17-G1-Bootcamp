package handler

import (
	"encoding/json"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/service"
	"net/http"
	"strconv"

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
	warehouses, err := h.sv.RetrieveAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = render.Render(w, r, response.NewResponse(warehouses, http.StatusOK))
}

func (h *WarehouseDefault) GetWarehouse(w http.ResponseWriter, r *http.Request) {
	idRequest := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idRequest)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse("invalid id", http.StatusBadRequest))
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
	warehouseJson := &request.WarehouseRequest{}
	if err := render.Bind(r, warehouseJson); err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}

	warehouse := models.NewWarehouse(
		0,									// placeholder, will be overwritten later
		*warehouseJson.Code,
		*warehouseJson.Address,
		*warehouseJson.Telephone,
		*warehouseJson.MinimumCapacity,
		*warehouseJson.MinimumTemperature,
	)

	warehouseResponse, err := h.sv.Register(*warehouse)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse("internal error", http.StatusInternalServerError))
		return
	}

	_ = render.Render(w, r, response.NewResponse(warehouseResponse, http.StatusCreated))
}

func (h *WarehouseDefault) PatchWarehouse(w http.ResponseWriter, r *http.Request) {
	idRequest := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idRequest)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse("invalid id", http.StatusBadRequest))
		return
	}

	var fields map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&fields)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}

	warehouseResponse, err := h.sv.PartialModify(id, fields)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusNotFound))
		return
	}

	_ = render.Render(w, r, response.NewResponse(warehouseResponse, http.StatusOK))
}

func (h *WarehouseDefault) DeleteWarehouse(w http.ResponseWriter, r *http.Request) {
	idRequest := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idRequest)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse("invalid id", http.StatusBadRequest))
		return
	}

	err = h.sv.Remove(id)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusNotFound))
		return
	}

	_ = render.Render(w, r, response.NewResponse(nil, http.StatusNoContent))
}
