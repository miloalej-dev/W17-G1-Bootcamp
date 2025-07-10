package handler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/service/default"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/request"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/response"
	"net/http"
	"strconv"
)

type EmployeeHandler struct {
	service *_default.EmployeeService
}

func NewEmployeeHandler(service *_default.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{
		service: service,
	}
}

// GetEmployees handles GET requests to retrieve all employees
func (h *EmployeeHandler) GetEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	employees, err := h.service.RetrieveAll()
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusNoContent))
		return
	}
	_ = render.Render(w, r, response.NewResponse(employees, http.StatusOK))
}

// GetEmployee handles GET requests to retrieve an employee by ID
func (h *EmployeeHandler) GetEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}

	employee, err := h.service.Retrieve(id)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusNotFound))
		return
	}

	_ = render.Render(w, r, response.NewResponse(employee, http.StatusOK))
}

// CreateEmployee handles POST requests to create a new employee
func (h *EmployeeHandler) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data := &request.EmployeeRequest{}
	if err := render.Bind(r, data); err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusUnprocessableEntity))
		return
	}
	employee := models.Employee{
		CardNumberId: *data.CardNumberId,
		FirstName:    *data.FirstName,
		LastName:     *data.LastName,
		WarehouseId:  *data.WarehouseId,
	}
	employeeRes, err := h.service.Register(employee)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusInternalServerError))
		return
	}
	_ = render.Render(w, r, response.NewResponse(employeeRes, http.StatusCreated))
}

// PutEmployee handles PUT requests to update an employee
func (h *EmployeeHandler) PutEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}
	data := &request.EmployeeRequest{}
	err = render.Bind(r, data)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}
	employee := models.Employee{
		Id:           id,
		CardNumberId: *data.CardNumberId,
		FirstName:    *data.FirstName,
		LastName:     *data.LastName,
		WarehouseId:  *data.WarehouseId,
	}
	updatedEmployee, err := h.service.Modify(employee)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusNotFound))
		return
	}
	_ = render.Render(w, r, response.NewResponse(updatedEmployee, http.StatusOK))
}

// PatchEmployee handles PATCH requests to partially update an employee
func (h *EmployeeHandler) PatchEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}
	var fields map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&fields)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}

	updatedEmployee, err := h.service.PartialModify(id, fields)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusInternalServerError))
		return
	}
	_ = render.Render(w, r, response.NewResponse(updatedEmployee, http.StatusOK))
}

// DeleteEmployee handles DELETE requests to remove an employee
func (h *EmployeeHandler) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}

	err = h.service.Remove(id)
	if err != nil {
		_ = render.Render(w, r, response.NewErrorResponse(err.Error(), http.StatusInternalServerError))
		return
	}

	_ = render.Render(w, r, response.NewResponse(nil, http.StatusNoContent))
}
