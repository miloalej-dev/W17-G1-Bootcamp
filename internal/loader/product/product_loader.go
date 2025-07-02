package productLoader

import "github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"

// Loader is an interface that represents the loader
type ProductLoader interface {
	Load() (v map[int]models.Product, err error)
}
