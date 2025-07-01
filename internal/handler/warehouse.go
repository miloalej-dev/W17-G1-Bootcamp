package handler

import (
	"net/http"

	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/service/warehouse"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"

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
