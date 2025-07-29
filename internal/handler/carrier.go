package handler

import (
	"encoding/json"
	"errors"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository"
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

func (h *CarrierDefault) GetCarriers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	carriers, err := h.sv.RetrieveAll()
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse("something went wrong", http.StatusInternalServerError))
		return
	}

	_ = render.Render(w, r, response.NewResponse(carriers, http.StatusOK))
}

func (h *CarrierDefault) GetCarrier(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idRequest := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idRequest)
	if err != nil || id < 1 {
		_ = render.Render(w, r, response.NewErrorResponse(ErrInvalidId.Error(), http.StatusBadRequest))
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
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusUnprocessableEntity))
		return
	}

	carrier := models.NewCarrier(
		0, // placeholder, will be overwritten later
		*carrierJson.CId,
		*carrierJson.CompanyName,
		*carrierJson.Address,
		*carrierJson.Telephone,
		*carrierJson.LocalityId,
	)

	carrierResponse, err := h.sv.Register(*carrier)
	if err != nil {
		if errors.Is(err, repository.ErrEntityAlreadyExists) {
			_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusConflict))
			return
		}
		if errors.Is(err, repository.ErrLocalityNotFound) {
			_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusUnprocessableEntity))
			return
		}

		_ = render.Render(w, r, response.NewErrorResponse("internal error", http.StatusInternalServerError))
		return
	}

	_ = render.Render(w, r, response.NewResponse(carrierResponse, http.StatusCreated))
}

// PutWarehouse handles PUT requests to update a warehouse
func (h *CarrierDefault) PutCarrier(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id < 1 {
		_ = render.Render(w, r, response.NewErrorResponse(ErrInvalidId.Error(), http.StatusBadRequest))
		return
	}

	data := &request.CarrierRequest{}

	err = render.Bind(r, data)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}

	carrier := models.NewCarrier(
		id, // placeholder, will be overwritten later
		*data.CId,
		*data.CompanyName,
		*data.Address,
		*data.Telephone,
		*data.LocalityId,
	)

	updatedCarrier, err := h.sv.Modify(*carrier)
	if err != nil {
		if errors.Is(err, repository.ErrEntityAlreadyExists) {
			_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusConflict))
			return
		}
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusInternalServerError))
		return
	}

	_ = render.Render(w, r, response.NewResponse(updatedCarrier, http.StatusOK))
}

func (h *CarrierDefault) PatchCarrier(w http.ResponseWriter, r *http.Request) {
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

	carrierResponse, err := h.sv.PartialModify(id, fields)
	if err != nil {
		if errors.Is(err, repository.ErrEntityNotFound) {
			_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusNotFound))
			return
		}
		if errors.Is(err, repository.ErrEntityAlreadyExists) {
			_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusConflict))
			return
		}

		_ = render.Render(w, r, response.NewErrorResponse("internal error", http.StatusInternalServerError))
		return
	}

	_ = render.Render(w, r, response.NewResponse(carrierResponse, http.StatusOK))
}

func (h *CarrierDefault) DeleteCarrier(w http.ResponseWriter, r *http.Request) {
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
