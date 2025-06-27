package productRepository

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

// FindAll is a method that returns a map of all Products
func (r *ProductMap) FindAll() (v map[int]models.Product, err error) {
	v = make(map[int]models.Product)
	// copy db
	for key, value := range r.db {
		v[key] = value
	}
	return
}
