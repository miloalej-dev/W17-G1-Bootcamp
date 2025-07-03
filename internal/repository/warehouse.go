package repository

import (
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
)

type WarehouseRepository interface {
	Repository[int, models.Warehouse]
}
