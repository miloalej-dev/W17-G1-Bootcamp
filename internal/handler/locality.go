package handler

import (
	"github.com/go-chi/render"
	_default "github.com/miloalej-dev/W17-G1-Bootcamp/internal/service/default"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/request"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/response"
	"net/http"
	"strconv"
)

type LocalityHandler struct {
	service _default.LocalityService
}

func NewLocalityHandler(service *_default.LocalityService) *LocalityHandler {
	return &LocalityHandler{service: *service}
}

func (h *LocalityHandler) GetLocalities(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	localities, err := h.service.RetrieveAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = render.Render(w, r, response.NewResponse(localities, http.StatusOK))
}

func (h *LocalityHandler) GetLocality(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}

	locality, err := h.service.RetrieveBySellerId(id)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusNotFound))
		return
	}

	res := []models.LocalitySellerCount{locality}

	_ = render.Render(w, r, response.NewResponse(res, http.StatusOK))
}

func (h *LocalityHandler) CreateLocality(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data := &request.LocalityRequest{}
	if err := render.Bind(r, data); err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusUnprocessableEntity))
		return
	}
	locality := models.Locality{
		Id:       data.Id,
		Locality: data.Locality,
		Province: data.Province,
		Country:  data.Country,
	}

	localityCreated, err := h.service.Register(locality)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusInternalServerError))
		return
	}
	_ = render.Render(w, r, response.NewResponse(localityCreated, http.StatusCreated))
}

func (h *LocalityHandler) GetCarrier(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idRequest := r.URL.Query().Get("id")
	var id int
	if idRequest != "" {
		var err error
		id, err = strconv.Atoi(idRequest)
		if err != nil {
			_ = render.Render(w, r, response.NewErrorResponse("invalid id", http.StatusBadRequest))
			return
		}
	}

	carriers, err := h.service.RetrieveCarriers(id)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusNotFound))
		return
	}

	_ = render.Render(w, r, response.NewResponse(carriers, http.StatusOK))
}
