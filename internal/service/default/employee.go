package _default

import (
	"errors"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
)

type EmployeeService struct {
	rp repository.EmployeeRepository
}

func NewEmployeeService(repo repository.EmployeeRepository) *EmployeeService {
	return &EmployeeService{rp: repo}
}

func (sv *EmployeeService) RetrieveAll() (e []models.Employee, err error) {
	e, err = sv.rp.FindAll()
	return
}

func (sv *EmployeeService) Retrieve(id int) (models.Employee, error) {
	return sv.rp.FindById(id)
}

func (sv *EmployeeService) Register(emp models.Employee) (models.Employee, error) {
	return sv.rp.Create(emp)
}

func (sv *EmployeeService) Modify(emp models.Employee) (models.Employee, error) {
	return sv.rp.Update(emp)
}

func (sv *EmployeeService) PartialModify(id int, fields map[string]interface{}) (models.Employee, error) {
	return sv.rp.PartialUpdate(id, fields)
}

func (sv *EmployeeService) Remove(id int) error {
	return sv.rp.Delete(id)
}

func (sv *EmployeeService) RetrieveInboundOrdersReport() ([]models.EmployeeInboundOrdersReport, error) {
	return sv.rp.InboundOrdersReport()
}

func (sv *EmployeeService) RetrieveInboundOrdersReportById(id int) (models.EmployeeInboundOrdersReport, error) {

	if id < 1 {
		return models.EmployeeInboundOrdersReport{}, errors.New("id must be greater than 0")
	}

	return sv.rp.InboundOrdersReportById(id)
}
