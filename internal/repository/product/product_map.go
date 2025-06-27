package product

import "github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"

// ProductMap is a struct that represents  repositorio of producto
type ProductMap struct {
	// db is a map of product
	db map[int]models.Product
}

func NewProductMap(db map[int]models.Product) *ProductMap {
	// defaultDb is an empty map
	defaultDb := make(map[int]models.Product)
	if db != nil {
		defaultDb = db
	}
	return &ProductMap{db: defaultDb}
}
