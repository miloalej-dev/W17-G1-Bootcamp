package service

import (
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
)

type WarehouseService interface {
	GetAll() (warehouses map[int]models.Warehouse, err error)
	GetById(id int) (warehouse models.Warehouse, err error)
}
