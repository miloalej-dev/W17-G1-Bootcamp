package json

import (
	"encoding/json"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
	"os"
)

// WarehouseFile is a struct that implements the LoaderWarehouse interface
type WarehouseFile struct {
	// path is the path to the file that contains the warehouses in JSON format
	path string
}

// NewWarehouseFile is a function that returns a new instance of EmployeeFile
func NewWarehouseFile(path string) *WarehouseFile {
	return &WarehouseFile{
		path: path,
	}
}

// Load is a method that loads the foos
func (l *WarehouseFile) Load() (warehouses map[int]models.Warehouse, err error) {
	// open file
	file, err := os.Open(l.path)
	if err != nil {
		return
	}
	defer file.Close()

	// decode file
	var warehousesDoc []models.Warehouse
	err = json.NewDecoder(file).Decode(&warehousesDoc)
	if err != nil {
		return
	}

	// serialize warehouses
	warehouses = make(map[int]models.Warehouse)
	for _, warehouse := range warehousesDoc {
		warehouses[warehouse.Id] = models.Warehouse{
			Id: 				warehouse.Id,
			Code:               warehouse.Code,
			Address:            warehouse.Address,
			Telephone:          warehouse.Telephone,
			MinimumCapacity:    warehouse.MinimumCapacity,
			MinimumTemperature: warehouse.MinimumTemperature,
		}
	}

	return
}
