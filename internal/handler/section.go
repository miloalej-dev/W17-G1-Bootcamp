package handler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/service/section"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
	"net/http"
	"strconv"
)

type SectionDefault struct {
	sv section.SectionService
}

func NewSectionDefault(sv section.SectionService) *SectionDefault {
	return &SectionDefault{sv: sv}
}

func (s *SectionDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sections, err := s.sv.FindAll()
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, err)
			return
		}
		if len(sections) == 0 {
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, nil)
			return
		}
		render.Status(r, http.StatusOK)
		render.JSON(w, r, sections)
	}
}

func (s *SectionDefault) FindByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idRequest := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idRequest)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, "Invalid Id")
			return
		}
		section, err := s.sv.FindByID(id)
		if err != nil {
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, err)
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
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err)
			return
		}

		createdSecton, err := s.sv.Create(section)
		if err.Error() == "Section already exists" {
			render.Status(r, http.StatusConflict)
			render.JSON(w, r, err)
			return
		}
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, err)
			return
		}
		render.Status(r, http.StatusCreated)
		render.JSON(w, r, createdSecton)

	}
}

func (s *SectionDefault) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idRequest := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idRequest)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, "Invalid Id")
			return
		}
		var section models.Section
		err = json.NewDecoder(r.Body).Decode(&section)
		section.Id = id
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err)
			return
		}
		updatedSecton, err := s.sv.Update(section)
		if err.Error() == "section not found" {
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, err)
			return
		}
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, err)
			return
		}
		render.Status(r, http.StatusOK)
		render.JSON(w, r, updatedSecton)
	}
}

func (s *SectionDefault) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idRequest := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idRequest)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, "Invalid Id")
			return
		}
		err = s.sv.Delete(id)
		if err != nil {
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, err)
			return
		}
		render.Status(r, http.StatusNoContent)
		render.JSON(w, r, nil)

	}
}
