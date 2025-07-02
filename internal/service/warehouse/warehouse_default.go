package service

import (
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository/warehouse"
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

func (s *WarehouseDefault) GetAll() (warehouses map[int]models.Warehouse, err error) {
	warehouses,err = s.rp.GetAll()
	return
}

func (s *WarehouseDefault) GetById(id int) (warehouse models.Warehouse, err error) {
	warehouse, err = s.rp.GetById(id)
	return
}

func (s *WarehouseDefault) Create(warehouseJson models.WarehouseDoc) (warehouse models.Warehouse, err error) {
	warehouse, err = s.rp.Create(warehouseJson)
	return
}

func (s *WarehouseDefault) Update(warehouse models.Warehouse) (err error) {
	err = s.rp.Update(warehouse)
	return
}

func (s *WarehouseDefault) Delete(id int) (err error) {
	err = s.rp.Delete(id)
	return
}
