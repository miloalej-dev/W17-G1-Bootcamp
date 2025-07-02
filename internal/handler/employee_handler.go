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

func (he *EmployeeHandler) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	employeeJson := &request.EmployeeRequest{}
	if err := render.Bind(r, employeeJson); err != nil {
		render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusBadRequest))
	}
	empoyee := models.Employee{
		CardNumberId: *employeeJson.CardNumberId,
		FirstName:    *employeeJson.FirstName,
		LastName:     *employeeJson.LastName,
		WarehouseId:  *employeeJson.WarehouseId,
	}
	employeeRes, err := he.sr.CreateEmployee(empoyee)
	if err != nil {
		render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusInternalServerError))
		return
	}
	render.Render(w, r, response.NewResponse(employeeRes, http.StatusCreated))
}

func (he *EmployeeHandler) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	idRequest := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idRequest)
	if err != nil {
		render.Render(w, r, response.NewErrorResponse("Invalid ID", http.StatusBadRequest))
		return
	}
	employeeJson := &request.EmployeeRequest{}
	if err := json.NewDecoder(r.Body).Decode(&employeeJson); err != nil {
		render.Render(w, r, response.NewErrorResponse("internal error", http.StatusInternalServerError))
	}
	employee, err := he.sr.GetEmployeeById(id)
	if err != nil {
		render.Render(w, r, response.NewErrorResponse("internal error", http.StatusInternalServerError))
		return
	}
	if employeeJson.CardNumberId != nil && *employeeJson.CardNumberId != employee.CardNumberId {
		employee.CardNumberId = *employeeJson.CardNumberId
	}
	if employeeJson.FirstName != nil && *employeeJson.FirstName != employee.FirstName {
		employee.FirstName = *employeeJson.FirstName
	}
	if employeeJson.LastName != nil && *employeeJson.LastName != employee.LastName {
		employee.LastName = *employeeJson.LastName
	}
	if employeeJson.WarehouseId != nil && *employeeJson.WarehouseId != employee.WarehouseId {
		employee.WarehouseId = *employeeJson.WarehouseId
	}
	employeeRes, err := he.sr.ModifyEmployee(employee)
	if err != nil {
		render.Render(w, r, response.NewErrorResponse("Internal error", http.StatusInternalServerError))
		return
	}
	render.Render(w, r, response.NewResponse(employeeRes, http.StatusOK))
}
func (he *EmployeeHandler) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	idRequest := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idRequest)
	if err != nil {
		render.Render(w, r, response.NewErrorResponse("Invalid ID", http.StatusBadRequest))
		return
	}

	err = he.sr.DeleteEmployee(id)
	if err != nil {
		render.Render(w, r, response.NewErrorResponse("Internal error", http.StatusInternalServerError))
		return
	}

	render.Render(w, r, response.NewResponse(nil, http.StatusOK))
}
