package product

import "github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"

// ProductMap es un struct que representa el repositorio de producto
type ProductMap struct {
	// db es un mapa de products
	db map[int]models.Product
}

func NewProductMap(db map[int]models.Product) *ProductMap {
	// defaultDb es un mapa vacio
	defaultDb := make(map[int]models.Product)
	if db != nil {
		defaultDb = db
	}
	return &ProductMap{db: defaultDb}
}
