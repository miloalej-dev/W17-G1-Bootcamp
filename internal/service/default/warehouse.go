package _default

import (
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository"
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

func (s *WarehouseDefault) RetrieveAll() ([]models.Warehouse, error) {
	return s.rp.FindAll()
}

func (s *WarehouseDefault) Retrieve(id int) (models.Warehouse, error) {
	return s.rp.FindById(id)
}

func (s *WarehouseDefault) Register(entity models.Warehouse) (models.Warehouse, error) {
	return s.rp.Create(entity)
}

func (s *WarehouseDefault) Modify(entity models.Warehouse) (models.Warehouse, error) {
	return s.rp.Update(entity)
}

func (s *WarehouseDefault) PartialModify(id int, fields map[string]interface{}) (models.Warehouse, error) {
	return s.rp.PartialUpdate(id, fields)
}

func (s *WarehouseDefault) Remove(id int) error {
	return s.rp.Delete(id)
}
