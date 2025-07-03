package json

import (
	"encoding/json"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
	"os"
)

// JsonFile is a struct that implements the LoaderVehicle interface
type JsonFile struct {
	// path is the path to the file that contains the vehicles in JSON format
	path string
}

// NewJSONFile is a function that returns a new instance of JsonFile
func NewJSONFile(path string) *JsonFile {
	return &JsonFile{
		path: path,
	}
}

// Load is a method that loads the employees
func (l *JsonFile) Load() (data map[int]models.Employee, err error) {
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
			CardNumberId: employee.FirstName,
			FirstName:    employee.FirstName,
			LastName:     employee.LastName,
			WarehouseId:  employee.WarehouseId,
		}
	}

	return
}
