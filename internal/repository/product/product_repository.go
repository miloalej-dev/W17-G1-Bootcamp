package productRepository

import "github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"

type ProductRepository interface {
	FindAll() (v map[int]models.Product, err error)
}
