package productLoader

import "github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"

// Loader is an interface that represents the loader
<<<<<<<< HEAD:internal/loader/product/product_loader.go
type ProductLoader interface {
	Load() (v map[int]models.Product, err error)
========
type Loader interface {
	Load() (warehouses map[int]models.Warehouse, err error)
>>>>>>>> origin/develop:internal/loader/warehouse/warehouse_loader.go
}
