package service

import (
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
)

type WarehouseService interface {
	RetrieveAll() ([]models.Warehouse, error)
	Retrieve(id int) (models.Warehouse, error)
	Register(seller models.Warehouse) (models.Warehouse, error)
	Modify(seller models.Warehouse) (models.Warehouse, error)
	PartialModify(id int, fields map[string]any) (models.Warehouse, error)
	Remove(id int) error
}
