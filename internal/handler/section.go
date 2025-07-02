package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/service/section"
	"net/http"
	"strconv"
)

type SectionDefault struct {
	sv section.SectionService
}

func NewSectionDefault(sv section.SectionService) *SectionDefault {
	return &SectionDefault{}
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
