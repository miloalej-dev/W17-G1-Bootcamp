package _default

import (
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
)

type EmployeeService struct {
	rp repository.Repository[int, models.Employee]
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
