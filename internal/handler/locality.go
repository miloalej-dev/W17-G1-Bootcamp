package handler

import (
	"github.com/go-chi/render"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/service"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/request"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/response"
	"net/http"
	"strconv"
)

type LocalityHandler struct {
	service service.LocalityService
}

func NewLocalityHandler(service service.LocalityService) *LocalityHandler {
	return &LocalityHandler{service: service}
}

func (h *LocalityHandler) GetLocalities(w http.ResponseWriter, r *http.Request) {
	localities, err := h.service.RetrieveAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = render.Render(w, r, response.NewResponse(localities, http.StatusOK))
}

func (h *LocalityHandler) GetLocality(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	if idParam == "" {
		localites, err := h.service.RetrieveAllLocalitiesBySeller()
		if err != nil {
			_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusNotFound))
			return
		}
		_ = render.Render(w, r, response.NewResponse(localites, http.StatusOK))
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(repository.ErrIDInvalid.Error(), http.StatusBadRequest))
		return
	}

	locality, err := h.service.RetrieveLocalityBySeller(id)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusNotFound))
		return
	}

	res := []models.LocalitySellerCount{locality}

	_ = render.Render(w, r, response.NewResponse(res, http.StatusOK))
}

func (h *LocalityHandler) PostLocality(w http.ResponseWriter, r *http.Request) {
	data := &request.LocalityRequest{}
	if err := render.Bind(r, data); err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}
	locality := models.LocalityDoc{
		Id:       data.Id,
		Locality: *data.Locality,
		Province: *data.Province,
		Country:  *data.Country,
	}

	localityCreated, err := h.service.RegisterWithNames(locality)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}
	_ = render.Render(w, r, response.NewResponse(localityCreated, http.StatusCreated))
}

func (h *LocalityHandler) GetCarrier(w http.ResponseWriter, r *http.Request) {
	idRequest := r.URL.Query().Get("id")
	var id int

	// If there is an Id, get all carriers by that locality
	// If there isn't an id, get all carriers from all localities
	if idRequest != "" {
		var err error
		id, err = strconv.Atoi(idRequest)
		if err != nil {
			_ = render.Render(w, r, response.NewErrorResponse(ErrInvalidId.Error(), http.StatusBadRequest))
			return
		}

		carriers, err := h.service.RetrieveCarriersByLocality(id)
		if err != nil {
			_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusNotFound))
			return
		}
		_ = render.Render(w, r, response.NewResponse(carriers, http.StatusOK))
		return
	}

	carriers, err := h.service.RetrieveCarriers()
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusNotFound))
		return
	}

	_ = render.Render(w, r, response.NewResponse(carriers, http.StatusOK))
}
