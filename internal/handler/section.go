package handler

import (
	"encoding/json"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/service"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
	"net/http"
	"strconv"
)

func NewSectionDefault(sv service.SectionService) *SectionDefault {
	return &SectionDefault{sv: sv}
}

type SectionDefault struct {
	sv service.SectionService
}

func (s *SectionDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sections, err := s.sv.FindAll()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}
		if len(sections) == 0 {
			response.JSON(w, http.StatusNotFound, nil)
			return
		}
		response.JSON(w, http.StatusOK, sections)
	}
}

func (s *SectionDefault) FindByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idRequest := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idRequest)
		if err != nil {
			response.JSON(w, http.StatusNotFound, nil)
			return
		}
		section, err := s.sv.FindByID(id)
		if err != nil {
			response.JSON(w, http.StatusNotFound, nil)
			return
		}
		render.Status(r, http.StatusOK)
		render.JSON(w, r, section)
	}
}

func (s *SectionDefault) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var section models.Section

		err := json.NewDecoder(r.Body).Decode(&section)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, nil)
			return
		}

		createdSecton, err := s.sv.Create(section)

		if err != nil {
			if err.Error() == "Section already exists" {
				response.JSON(w, http.StatusConflict, nil)
				return
			}
			response.JSON(w, http.StatusBadRequest, nil)
			return
		}
		response.JSON(w, http.StatusCreated, createdSecton)

	}
}

func (s *SectionDefault) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idRequest := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idRequest)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, nil)
			return
		}
		var section models.Section
		err = json.NewDecoder(r.Body).Decode(&section)
		section.Id = id
		if err != nil {
			response.JSON(w, http.StatusBadRequest, nil)
			return
		}
		updatedSecton, err := s.sv.Update(section)

		if err != nil {
			if err.Error() == "section not found" {
				response.JSON(w, http.StatusNotFound, nil)
				return
			}
			response.JSON(w, http.StatusInternalServerError, err)
			return
		}
		response.JSON(w, http.StatusOK, updatedSecton)
	}
}

func (s *SectionDefault) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idRequest := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idRequest)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, nil)
			return
		}
		err = s.sv.Delete(id)
		if err != nil {
			response.JSON(w, http.StatusNotFound, nil)
			return
		}
		response.JSON(w, http.StatusNoContent, nil)
	}
}
