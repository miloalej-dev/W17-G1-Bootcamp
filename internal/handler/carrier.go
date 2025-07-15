package handler

import (
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/service"
	"net/http"
	"strconv"

	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/request"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/response"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// CarrierDefault is a struct with methods that represent handlers for carriers
type CarrierDefault struct {
	// sv is the service that will be used by the handler
	sv service.CarrierService
}

// NewCarrierDefault is a function that returns a new instance of CarrierDefault
func NewCarrierDefault(sv service.CarrierService) *CarrierDefault {
	return &CarrierDefault{sv: sv}
}

func (h *CarrierDefault) GetCarrier(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idRequest := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idRequest)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse("invalid id", http.StatusBadRequest))
		return
	}

	carrier, err := h.sv.Retrieve(id)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusNotFound))
		return
	}

	_ = render.Render(w, r, response.NewResponse(carrier, http.StatusOK))
}

func (h *CarrierDefault) PostCarrier(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	carrierJson := &request.CarrierRequest{}
	if err := render.Bind(r, carrierJson); err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}

	carrier := models.NewCarrier(
		0,							// placeholder, will be overwritten later
		*carrierJson.CId,
		*carrierJson.CompanyName,
		*carrierJson.Address,
		*carrierJson.Telephone,
		*carrierJson.LocalityId,
	)

	carrierResponse, err := h.sv.Register(*carrier)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse("internal error", http.StatusInternalServerError))
		return
	}

	_ = render.Render(w, r, response.NewResponse(carrierResponse, http.StatusCreated))
}
