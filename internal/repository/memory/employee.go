package memory

import (
	"errors"
	"strconv"

	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
)

type EmployeeMap struct {
	db map[int]models.Employee
}

func NewEmployeeMap(db map[int]models.Employee) *EmployeeMap {
	return &EmployeeMap{db: db}
}

func (r *EmployeeMap) FindAll() ([]models.Employee, error) {
	employees := make([]models.Employee, 0)
	for _, e := range r.db {
		employees = append(employees, e)
	}
	return employees, nil
}

func (e *EmployeeMap) FindById(id int) (models.Employee, error) {
	employee, exists := e.db[id]
	if !exists {
		return models.Employee{}, errors.New("employee no found")
	}
	return employee, nil
}

func (e *EmployeeMap) Create(emp models.Employee) (models.Employee, error) {
	id := len(e.db) + 1
	newEmployee := models.Employee{
		Id:           id,
		CardNumberId: emp.CardNumberId,
		FirstName:    emp.FirstName,
		LastName:     emp.LastName,
		WarehouseId:  emp.WarehouseId,
	}
	e.db[id] = newEmployee
	return newEmployee, nil
}

func (e *EmployeeMap) Update(emp models.Employee) (models.Employee, error) {
	_, exists := e.db[emp.Id]
	if !exists {
		return models.Employee{}, errors.New("employee does not exists")
	}
	e.db[emp.Id] = emp
	return emp, nil
}

func (e *EmployeeMap) PartialUpdate(id int, fields map[string]interface{}) (models.Employee, error) {
	employee, exists := e.db[id]

	if !exists {
		return models.Employee{}, errors.New("employee Not Found")
	}

	if val, ok := fields["card_number_id"]; ok {
		employee.CardNumberId = val.(string)
	}
	if val, ok := fields["first_name"]; ok {
		employee.FirstName = val.(string)
	}
	if val, ok := fields["last_name"]; ok {
		employee.LastName = val.(string)
	}
	if val, ok := fields["warehouse_id"]; ok {
		idWarehouse, _ := strconv.Atoi(val.(string))
		employee.WarehouseId = idWarehouse
	}

	e.db[id] = employee
	return employee, nil
}

func (e *EmployeeMap) Delete(id int) error {
	_, exists := e.db[id]
	if !exists {
		return errors.New("employee not fount")
	}
	delete(e.db, id)
	return nil
}
