package handler

import (
	"github.com/go-chi/render"
	_default "github.com/miloalej-dev/W17-G1-Bootcamp/internal/service/default"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/response"
	"net/http"
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
