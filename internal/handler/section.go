package handler

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository"
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
	w.Header().Set("Content-Type", "application/json")
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
		WarehousesId:       *data.WarehousesId,
		ProductTypeId:      *data.ProductTypeId,
	}

	createdSection, err := s.sv.Register(section)

	if err != nil {
		if errors.Is(err, repository.ErrEntityAlreadyExists) {
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
		if errors.Is(err, repository.ErrEntityNotFound) {
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

func (h *SectionHandler) GetSectionReportProducts(w http.ResponseWriter, r *http.Request) {
	// Obtener el ID del query param
	idParam := r.URL.Query().Get("id")

	var data interface{}
	var err error

	if idParam != "" {
		// Si hay un ID, lo convertimos a int
		id, errConv := strconv.Atoi(idParam)
		if errConv != nil {
			// Manejar error de conversión
			return
		}
		// Llamamos al servicio para un ID específico
		data, err = h.sv.RetrieveSectionReport(&id)

	} else {
		// Si no hay ID, llamamos al servicio para obtener todos los reportes
		data, err = h.sv.RetrieveSectionReport(nil)
	}

	if err != nil {
		if errors.Is(err, repository.ErrSectionNotFound) {
			_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusNotFound))
			return
		}
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}

	_ = render.Render(w, r, response.NewResponse(data, http.StatusOK))
}
