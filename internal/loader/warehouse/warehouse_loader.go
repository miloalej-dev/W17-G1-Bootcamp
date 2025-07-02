package loaderWarehouse

import "github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"

type Loader interface {
	Load() (warehouses map[int]models.Warehouse, err error)
}
