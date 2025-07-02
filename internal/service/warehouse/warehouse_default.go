package service

import (
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository/warehouse"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
)

// WarehouseDefault is a struct that represents the default service for warehouses
type WarehouseDefault struct {
	// rp is the repository that will be used by the service
	rp repository.WarehouseRepository
}

// NewWarehouseDefault is a function that returns a new instance of WarehouseDefault
func NewWarehouseDefault(rp repository.WarehouseRepository) *WarehouseDefault {
	return &WarehouseDefault{rp: rp}
}

func (s *WarehouseDefault) FindAll() ([]models.Warehouse, error) {
	return s.rp.FindAll()
}

func (s *WarehouseDefault) FindById(id int) (models.Warehouse, error) {
	return s.rp.FindById(id)
}

func (s *WarehouseDefault) Create(entity models.Warehouse) (models.Warehouse, error) {
	return s.rp.Create(entity)
}

func (s *WarehouseDefault) Update(entity models.Warehouse) (models.Warehouse, error) {
	return s.rp.Update(entity)
}

func (s *WarehouseDefault) PartialUpdate(id int, fields map[string]interface{}) (models.Warehouse, error) {
	return s.rp.PartialUpdate(id, fields)
}

func (s *WarehouseDefault) Delete(id int) error {
	return s.rp.Delete(id)
}
