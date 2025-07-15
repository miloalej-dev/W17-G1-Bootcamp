package database

import (
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
	"gorm.io/gorm"
)

type EmployeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) *EmployeeRepository {
	return &EmployeeRepository{
		db: db,
	}
}

func (e EmployeeRepository) FindAll() ([]models.Employee, error) {
	var employees []models.Employee
	result := e.db.Find(&employees)
	if result.Error != nil {
		return nil, result.Error
	}
	if len(employees) == 0 {
		return nil, repository.ErrEntityNotFound
	}
	return employees, nil
}

func (e EmployeeRepository) FindById(id int) (models.Employee, error) {
	var employee models.Employee
	result := e.db.First(&employee, id)
	if result.Error != nil {
		return models.Employee{}, repository.ErrEntityNotFound
	}
	return employee, nil
}

func (e EmployeeRepository) Create(employee models.Employee) (models.Employee, error) {
	result := e.db.Create(&employee)
	if result.Error != nil {
		return models.Employee{}, result.Error
	}
	return employee, nil
}

func (e EmployeeRepository) Update(employee models.Employee) (models.Employee, error) {
	result := e.db.Save(&employee)
	if result.Error != nil {
		return models.Employee{}, result.Error
	}
	return employee, nil
}

func (e EmployeeRepository) PartialUpdate(id int, fields map[string]interface{}) (models.Employee, error) {
	var employee models.Employee
	result := e.db.First(&employee, id)
	if result.Error != nil {
		return models.Employee{}, repository.ErrEntityNotFound
	}
	result = e.db.Model(&employee).Updates(fields)
	if result.Error != nil {
		return models.Employee{}, result.Error
	}
	return employee, nil
}

func (e EmployeeRepository) Delete(id int) error {
	result := e.db.Delete(&models.Employee{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
