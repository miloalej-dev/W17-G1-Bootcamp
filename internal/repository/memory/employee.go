package memory

import (
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
)

type EmployeeMap struct {
	db map[int]models.Employee
}

func NewEmployeeMap(db map[int]models.Employee) *EmployeeMap {
	return &EmployeeMap{db: db}
}

func (e *EmployeeMap) FindAll() ([]models.Employee, error) {
	employees := make([]models.Employee, 0)
	for _, emp := range e.db {
		employees = append(employees, emp)
	}
	if len(employees) <= 0 {
		return nil, repository.ErrEntityNotFound
	}
	return employees, nil
}

func (e *EmployeeMap) FindById(id int) (models.Employee, error) {
	employee, exists := e.db[id]
	if !exists {
		return models.Employee{}, repository.ErrEntityNotFound
	}
	return employee, nil
}

func (e *EmployeeMap) Create(emp models.Employee) (models.Employee, error) {
	id := len(e.db) + 1
	emp.Id = id
	e.db[id] = emp
	return emp, nil
}

func (e *EmployeeMap) Update(emp models.Employee) (models.Employee, error) {
	_, exists := e.db[emp.Id]
	if !exists {
		return models.Employee{}, repository.ErrEntityNotFound
	}
	e.db[emp.Id] = emp
	return emp, nil
}

func (e *EmployeeMap) PartialUpdate(id int, fields map[string]interface{}) (models.Employee, error) {
	employee, exists := e.db[id]

	if !exists {
		return models.Employee{}, repository.ErrEntityNotFound
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
		employee.WarehouseId = int(val.(float64))
	}

	e.db[id] = employee
	return employee, nil
}

func (e *EmployeeMap) Delete(id int) error {
	_, exists := e.db[id]
	if !exists {
		return repository.ErrEntityNotFound
	}
	delete(e.db, id)
	return nil
}
