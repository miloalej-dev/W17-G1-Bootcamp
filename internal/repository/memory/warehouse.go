package memory

import (
	"errors"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
)

// Warehouse repository
type WarehouseMap struct {
	db map[int]models.Warehouse
}

// Creates a new Warehouse repository
func NewWarehouseMap(defaultDB map[int]models.Warehouse) *WarehouseMap {
	return &WarehouseMap{
		db: defaultDB,
	}
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
		return models.Warehouse{}, errors.New("Warehouse Not Found")
	}
	return warehouse, nil
}

func (r *WarehouseMap) Create(warehouse models.Warehouse) (models.Warehouse, error) {
	id := len(r.db) + 1
	warehouse = models.Warehouse{
		ID: id,
		WarehouseAttributes: models.WarehouseAttributes{
			Code:               warehouse.Code,
			Address:            warehouse.Address,
			Telephone:          warehouse.Telephone,
			MinimumCapacity:    warehouse.MinimumCapacity,
			MinimumTemperature: warehouse.MinimumTemperature,
		},
	}

	r.db[id] = warehouse
	return warehouse, nil
}

func (r *WarehouseMap) Update(warehouse models.Warehouse) (models.Warehouse, error) {
	_, found := r.db[warehouse.ID]
	if !found {
		return models.Warehouse{}, errors.New("Warehouse Not Found")
	}

	r.db[warehouse.ID] = warehouse
	return warehouse, nil
}

func (r *WarehouseMap) PartialUpdate(id int, fields map[string]interface{}) (models.Warehouse, error) {
	warehouse, found := r.db[id]

	if !found {
		return models.Warehouse{}, errors.New("Warehouse Not Found")
	}

	if val, ok := fields["code"]; ok {
		warehouse.Code = val.(string)
	}
	if val, ok := fields["address"]; ok {
		warehouse.Code = val.(string)
	}
	if val, ok := fields["telephone"]; ok {
		warehouse.Code = val.(string)
	}
	if val, ok := fields["minimum_capacity"]; ok {
		warehouse.Code = val.(string)
	}
	if val, ok := fields["minimum_temperature"]; ok {
		warehouse.Code = val.(string)
	}

	r.db[id] = warehouse
	return warehouse, nil
}

func (r *WarehouseMap) Delete(id int) error {
	_, found := r.db[id]
	if !found {
		return errors.New("Warehouse Not Found")
	}
	delete(r.db, id)
	return nil
}
