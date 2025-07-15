package handler

import (
	"github.com/go-chi/render"
	_default "github.com/miloalej-dev/W17-G1-Bootcamp/internal/service/default"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
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

	locality, err := h.service.Retrieve(id)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusNotFound))
		return
	}

	res := []models.Locality{locality}

	_ = render.Render(w, r, response.NewResponse(res, http.StatusOK))
}
