package handler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/service"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/request"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/response"
	"net/http"
	"strconv"
)

func NewSectionDefault(sv service.SectionService) *SectionHandler {
	return &SectionHandler{sv: sv}
}

type SectionHandler struct {
	sv service.SectionService
}

func (s *SectionHandler) GetSections(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	sections, err := s.sv.RetrieveAll()
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusInternalServerError))
		return
	}
	if len(sections) == 0 {
		_ = render.Render(w, r, response.NewErrorResponse("not found", http.StatusNotFound))
		return
	}
	_ = render.Render(w, r, response.NewResponse(sections, http.StatusOK))

}

func (s *SectionHandler) GetSection(w http.ResponseWriter, r *http.Request) {
	idRequest := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idRequest)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}
	section, err := s.sv.Retrieve(id)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusNotFound))
		return
	}
	_ = render.Render(w, r, response.NewResponse(section, http.StatusOK))

}

func (s *SectionHandler) PostSection(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data := &request.SectionRequest{}
	if err := render.Bind(r, data); err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}

	section := models.Section{
		SectionNumber:      *data.SectionNumber,
		CurrentTemperature: *data.CurrentTemperature,
		MinimumTemperature: *data.MinimumTemperature,
		CurrentCapacity:    *data.CurrentCapacity,
		MinimumCapacity:    *data.MinimumCapacity,
		MaximumCapacity:    *data.MaximumCapacity,
		WarehouseId:        *data.WarehouseId,
		ProductTypeId:      *data.ProductTypeId,
		ProductsBatch:      *data.ProductsBatch,
	}

	createdSection, err := s.sv.Register(section)

	if err != nil {
		if err.Error() == "Section already exists" {
			_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusConflict))
			return
		}
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}
	_ = render.Render(w, r, response.NewResponse(createdSection, http.StatusOK))

}

func (s *SectionHandler) PatchSection(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idRequest := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idRequest)
	if err != nil {
		_ = render.Render(w, r, response.NewResponse(err.Error(), http.StatusBadRequest))
		return
	}
	var fields map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&fields)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}
	updatedSecton, err := s.sv.PartialModify(id, fields)

	if err != nil {
		if err.Error() == "section not found" {
			_ = render.Render(w, r, response.NewResponse(err.Error(), http.StatusNotFound))
			return
		}
		_ = render.Render(w, r, response.NewResponse(err.Error(), http.StatusInternalServerError))
		return
	}
	_ = render.Render(w, r, response.NewResponse(updatedSecton, http.StatusOK))

}

func (s *SectionHandler) DeleteSection(w http.ResponseWriter, r *http.Request) {
	idRequest := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idRequest)
	if err != nil {
		_ = render.Render(w, r, response.NewResponse(err.Error(), http.StatusBadRequest))
		return
	}
	err = s.sv.Remove(id)
	if err != nil {
		_ = render.Render(w, r, response.NewResponse(err.Error(), http.StatusNotFound))
		return
	}
	_ = render.Render(w, r, response.NewResponse(nil, http.StatusNoContent))

}
