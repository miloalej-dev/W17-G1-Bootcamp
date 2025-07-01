package repository

import (
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/loader/warehouse"
)

// Warehouse repository
type WarehouseMap struct {
	db	map[int]models.Warehouse
}

// Creates a new Warehouse
func NewWarehouseMap() *WarehouseMap {
	defaultDB := make(map[int]models.Warehouse)

	ld := loader.NewWarehouseJSONFile("docs/db/warehouse.json")
	db, err := ld.Load()
	if err != nil {
		return nil
	}

	if db != nil {
		defaultDB = db
	}

	return &WarehouseMap{
		db: defaultDB,
	}
}

// Returns all the warehouses
func (r *WarehouseMap) GetAll() (warehouses map[int]models.Warehouse, err error) {
	warehouses = make(map[int]models.Warehouse)
	for id,warehouse := range r.db {
		warehouses[id] = warehouse
	}
	return
}


func (r *WarehouseMap) GetById(id int) (warehouse models.Warehouse, err error) {
	warehouse =  r.db[id]
	return
}

func (r *WarehouseMap) Create(warehouseJson models.WarehouseDoc) (warehouse models.Warehouse, err error) {
	id := len(r.db) + 1
	warehouse = models.Warehouse {
		ID:	id,
		WarehouseAttributes: models.WarehouseAttributes {
			Code: warehouseJson.Code,
			Address: warehouseJson.Address,
			Telephone: warehouseJson.Telephone,
			MinimunCapacity: warehouseJson.MinimunCapacity,
			MinimumTemperature: warehouseJson.MinimumTemperature,
		},
	}

	r.db[id] = warehouse
	return
}
