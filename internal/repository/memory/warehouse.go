package memory

import (
	"errors"
	loader "github.com/miloalej-dev/W17-G1-Bootcamp/internal/loader/json"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
)

// Warehouse repository
type WarehouseMap struct {
	db map[int]models.Warehouse
}

// Creates a new Warehouse repository
func NewWarehouseMap() *WarehouseMap {
	// defaultDB is an empty map
	defaultDB := make(map[int]models.Warehouse)

	ld := loader.NewWarehouseFile("docs/db/warehouse.json")
	db, err := ld.Load()

	if err != nil {
		return nil
	}
	if db != nil {
		defaultDB = db
	}
	return &WarehouseMap{db: defaultDB}
}

func (r *WarehouseMap) FindAll() ([]models.Warehouse, error) {
	warehouses := make([]models.Warehouse, 0)
	for _, warehouse := range r.db {
		warehouses = append(warehouses, warehouse)
	}
	return warehouses, nil
}

func (r *WarehouseMap) FindById(id int) (models.Warehouse, error) {
	warehouse, found := r.db[id]
	if !found {
		return models.Warehouse{}, errors.New("warehouse not found")
	}
	return warehouse, nil
}

func (r *WarehouseMap) Create(warehouse models.Warehouse) (models.Warehouse, error) {
	id := len(r.db) + 1
	warehouse.Id = id
	r.db[id] = warehouse
	return warehouse, nil
}

func (r *WarehouseMap) Update(warehouse models.Warehouse) (models.Warehouse, error) {
	_, found := r.db[warehouse.Id]
	if !found {
		return models.Warehouse{}, errors.New("warehouse not found")
	}

	r.db[warehouse.Id] = warehouse
	return warehouse, nil
}

func (r *WarehouseMap) PartialUpdate(id int, fields map[string]interface{}) (models.Warehouse, error) {
	warehouse, found := r.db[id]

	if !found {
		return models.Warehouse{}, errors.New("warehouse not found")
	}

	if val, ok := fields["code"]; ok {
		warehouse.Code = val.(string)
	}
	if val, ok := fields["address"]; ok {
		warehouse.Address = val.(string)
	}
	if val, ok := fields["telephone"]; ok {
		warehouse.Telephone = val.(string)
	}
	if val, ok := fields["minimum_capacity"]; ok {
		warehouse.MinimumCapacity = val.(int)
	}
	if val, ok := fields["minimum_temperature"]; ok {
		warehouse.MinimumTemperature = val.(int)
	}

	r.db[id] = warehouse
	return warehouse, nil
}

func (r *WarehouseMap) Delete(id int) error {
	_, found := r.db[id]
	if !found {
		return errors.New("warehouse not found")
	}
	delete(r.db, id)
	return nil
}
