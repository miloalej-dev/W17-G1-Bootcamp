package service

import "github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"

type EmployeeService interface {
	RetrieveAll() (e []models.Employee, err error)

	Retrieve(id int) (models.Employee, error)

	Register(emp models.Employee) (models.Employee, error)
	Modify(emp models.Employee) (models.Employee, error)

	PartialModify(id int, fields map[string]interface{}) (models.Employee, error)

	Remove(id int) error

	RetrieveInboundOrdersReport() ([]models.EmployeeInboundOrdersReport, error)

	RetrieveInboundOrdersReportById(id int) (models.EmployeeInboundOrdersReport, error)
}
