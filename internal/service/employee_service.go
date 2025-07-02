package service

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

func (sv *EmployeeService) GetEmployees() (e []models.Employee, err error) {
	e, err = sv.rp.FindAll()
	return
}
