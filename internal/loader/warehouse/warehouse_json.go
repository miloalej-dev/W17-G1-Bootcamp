package loaderWarehouse

import (
	"encoding/json"
	"fmt"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
	"os"
)

// JsonFile is a struct that implements the LoaderWarehouse interface
type JsonFile struct {
	// path is the path to the file that contains the warehouses in JSON format
	path string
}

// NewJSONFile is a function that returns a new instance of JsonFile
func NewJSONFile(path string) *JsonFile {
	return &JsonFile{
		path: path,
	}
}

// Load is a method that loads the foos
func (l *JsonFile) Load() (warehouses map[int]models.Warehouse, err error) {
	// open file
	file, err := os.Open(l.path)
	if err != nil {
		fmt.Println("xx")
		return
	}
	defer file.Close()

	// decode file
	var warehousesDoc []models.WarehouseDoc
	err = json.NewDecoder(file).Decode(&warehousesDoc)
	if err != nil {
		return
	}

	// serialize warehouses
	warehouses = make(map[int]models.Warehouse)
	for _, warehouse := range warehousesDoc {
		warehouses[warehouse.ID] = models.Warehouse{
			ID: warehouse.ID,
			WarehouseAttributes: models.WarehouseAttributes{
				Code:               warehouse.Code,
				Address:            warehouse.Address,
				Telephone:          warehouse.Telephone,
				MinimumCapacity:    warehouse.MinimumCapacity,
				MinimumTemperature: warehouse.MinimumTemperature,
			},
		}
	}

	return
}
