package database

import (
	"errors"
	"gorm.io/gorm"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository"
)

// Warehouse repository
type WarehouseDB struct {
	db *gorm.DB
}

// Creates a new Warehouse repository
func NewWarehouseDB(db *gorm.DB) *WarehouseDB {
	return &WarehouseDB{db: db}
}

func (r *WarehouseDB) FindAll() ([]models.Warehouse, error) {
	warehouses := make([]models.Warehouse, 0)
	result := r.db.Find(&warehouses)
	if result.Error != nil {
		return nil, result.Error
	}
	return warehouses, nil
}

func (r *WarehouseDB) FindById(id int) (models.Warehouse, error) {
	var warehouse models.Warehouse
	result := r.db.First(&warehouse, id)
	if result.Error != nil {
		return models.Warehouse{}, result.Error
	}
	return warehouse, nil
}

func (r *WarehouseDB) Create(warehouse models.Warehouse) (models.Warehouse, error) {
	result := r.db.Create(&warehouse)

	switch {
	case errors.Is(result.Error, gorm.ErrForeignKeyViolated):
			return models.Warehouse{}, repository.ErrLocalityNotFound
	case result.Error != nil:
			return models.Warehouse{}, result.Error
	}

	return warehouse, nil
}

func (r *WarehouseDB) Update(warehouse models.Warehouse) (models.Warehouse, error) {
	result := r.db.Save(&warehouse)
	if result.Error == nil {
		return warehouse, nil
	}

	return models.Warehouse{}, result.Error
}

func (r *WarehouseDB) PartialUpdate(id int, fields map[string]interface{}) (models.Warehouse, error) {
	var warehouse models.Warehouse
	result := r.db.First(&warehouse, id)
	switch {
	case errors.Is(result.Error, gorm.ErrRecordNotFound):
			return models.Warehouse{}, repository.ErrEntityNotFound
	case result.Error != nil:
			return models.Warehouse{}, result.Error
	}

	if val, ok := fields["code"]; ok {
		warehouse.WarehouseCode = val.(string)
	}
	if val, ok := fields["address"]; ok {
		warehouse.Address = val.(string)
	}
	if val, ok := fields["telephone"]; ok {
		warehouse.Telephone = val.(string)
	}
	if val, ok := fields["minimum_capacity"]; ok {
		warehouse.MinimumCapacity = int(val.(float64))
	}
	if val, ok := fields["minimum_temperature"]; ok {
		warehouse.MinimumTemperature = int(val.(float64))
	}
	if val, ok := fields["locality_id"]; ok {
		warehouse.LocalityId = int(val.(float64))
	}

	result = r.db.Save(&warehouse)
	if result.Error != nil {
		return models.Warehouse{}, result.Error
	}
	return warehouse, nil
}

func (r *WarehouseDB) Delete(id int) error {
	var warehouse models.Warehouse
	result := r.db.Delete(&warehouse, id)
	if result.RowsAffected < 1 {
		return repository.ErrEntityNotFound
	}
	return result.Error
}
