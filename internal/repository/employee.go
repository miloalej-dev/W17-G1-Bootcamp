package repository

import "github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"

type EmployeeRepository interface {
	Repository[int, models.Employee]
}
