package memory

import (
	"errors"
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
	panic("implement me")
}

func (e *EmployeeMap) Delete(id int) error {
	//TODO implement me
	panic("implement me")
}
