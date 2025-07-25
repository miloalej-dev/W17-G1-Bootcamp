package database

import (
	"errors"
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

	return employees, nil
}

func (e EmployeeRepository) FindById(id int) (models.Employee, error) {
	var employee models.Employee
	result := e.db.First(&employee, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return models.Employee{}, repository.ErrEntityNotFound
	}

	if result.Error != nil {
		return models.Employee{}, result.Error
	}
	return employee, nil
}

func (e EmployeeRepository) Create(employee models.Employee) (models.Employee, error) {
	result := e.db.Create(&employee)
	switch {
	case errors.Is(result.Error, gorm.ErrForeignKeyViolated):
		return models.Employee{}, repository.ErrForeignKeyViolation
	case result.Error != nil:
		return models.Employee{}, result.Error
	}
	return employee, nil
}

func (e EmployeeRepository) Update(employee models.Employee) (models.Employee, error) {
	result := e.db.Save(&employee)
	switch {
	case errors.Is(result.Error, gorm.ErrForeignKeyViolated):
		return models.Employee{}, repository.ErrForeignKeyViolation
	case result.Error != nil:
		return models.Employee{}, result.Error
	}
	return employee, nil
}

func (e EmployeeRepository) PartialUpdate(id int, fields map[string]interface{}) (models.Employee, error) {
	var employee models.Employee
	result := e.db.First(&employee, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return models.Employee{}, repository.ErrEntityNotFound
	}
	if result.Error != nil {
		return models.Employee{}, result.Error
	}
	result = e.db.Model(&employee).Updates(fields)
	switch {
	case errors.Is(result.Error, gorm.ErrForeignKeyViolated):
		return models.Employee{}, repository.ErrForeignKeyViolation
	case result.Error != nil:
		return models.Employee{}, result.Error
	}
	return employee, nil
}

func (e EmployeeRepository) Delete(id int) error {
	result := e.db.Delete(&models.Employee{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return repository.ErrEntityNotFound
	}
	return nil
}

// InboundOrdersReport  returns inbound orders count for all employees
func (e EmployeeRepository) InboundOrdersReport() ([]models.EmployeeInboundOrdersReport, error) {
	var reports []models.EmployeeInboundOrdersReport

	result := e.db.Table("employees e").
		Select(`
            e.id,
            e.card_number_id,
            e.first_name,
            e.last_name,
			e.warehouse_id,
            COUNT(io.id) AS inbound_orders_count
        `).
		Joins("LEFT JOIN inbound_orders io ON e.id = io.employee_id").
		Joins("LEFT JOIN warehouses w ON e.warehouse_id = w.id").
		Group("e.id, e.card_number_id, e.first_name, e.last_name, e.warehouse_id").
		Order("e.id").
		Scan(&reports)

	if result.Error != nil {
		return nil, result.Error
	}

	return reports, nil
}

// InboundOrdersReportById returns inbound orders count for a specific employee
func (e EmployeeRepository) InboundOrdersReportById(id int) (models.EmployeeInboundOrdersReport, error) {
	var report models.EmployeeInboundOrdersReport

	result := e.db.Table("employees e").
		Select(`
            e.id,
            e.card_number_id,
            e.first_name,
            e.last_name,
			e.warehouse_id,
            COUNT(io.id) AS inbound_orders_count
        `).
		Joins("LEFT JOIN inbound_orders io ON e.id = io.employee_id").
		Joins("LEFT JOIN warehouses w ON e.warehouse_id = w.id").
		Where("e.id = ?", id).
		Group("e.id, e.card_number_id, e.first_name, e.last_name, e.warehouse_id").
		Scan(&report)

	switch {
	case result.RowsAffected == 0:
		return models.EmployeeInboundOrdersReport{}, repository.ErrEntityNotFound
	case result.Error != nil:
		return models.EmployeeInboundOrdersReport{}, result.Error
	}

	return report, nil
}
