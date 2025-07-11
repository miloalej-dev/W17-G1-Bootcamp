package json

import (
	"encoding/json"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
	"os"
)

// EmployeeFile is a struct that implements the LoaderVehicle interface
type EmployeeFile struct {
	// path is the path to the file that contains the vehicles in JSON format
	path string
}

// NewEmployeeFile is a function that returns a new instance of EmployeeFile
func NewEmployeeFile(path string) *EmployeeFile {
	return &EmployeeFile{
		path: path,
	}
}

// Load is a method that loads the employees
func (l *EmployeeFile) Load() (data map[int]models.Employee, err error) {
	// open file
	file, err := os.Open(l.path)
	if err != nil {
		return
	}
	defer file.Close()

	// decode file
	var employees []models.Employee
	err = json.NewDecoder(file).Decode(&employees)
	if err != nil {
		return
	}

	// serialize vehicles
	data = make(map[int]models.Employee)
	for _, employee := range employees {
		data[employee.Id] = models.Employee{
			Id:           employee.Id,
			CardNumberId: employee.CardNumberId,
			FirstName:    employee.FirstName,
			LastName:     employee.LastName,
			WarehouseId:  employee.WarehouseId,
		}
	}

	return
}
