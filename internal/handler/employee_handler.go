package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/service"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/response"
	"net/http"
	"strconv"
)

type EmployeeHandler struct {
	sr *service.EmployeeService
}

func NewEmployeeHandler(service *service.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{
		sr: service,
	}
}

func (he *EmployeeHandler) GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	employees, err := he.sr.GetEmployees()
	if err != nil {
		render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}
	render.Render(w, r, response.NewResponse(employees, http.StatusOK))
}

func (he *EmployeeHandler) GetEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idRequest := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idRequest)
	if err != nil {
		render.Render(w, r, response.NewErrorResponse("Invalid ID", http.StatusBadRequest))
		return
	}

	employee, err := he.sr.GetEmployeeById(id)
	if err != nil {
		render.Render(w, r, response.NewErrorResponse("Internal error", http.StatusInternalServerError))
		return
	}

	render.Render(w, r, response.NewResponse(employee, http.StatusOK))
}
