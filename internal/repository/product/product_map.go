// Package productRepository provides an in-memory implementation of a repository for products.
// It uses a map to store product data, making it suitable for testing or simple applications
// where data persistence is not required.

package productRepository

import (
	"errors"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
)

// ProductMap implements a product repository using an in-memory map.
// The key of the map is the product ID.
type ProductMap struct {
	// db holds the product data. It is a private field to encapsulate storage.
	db map[int]models.Product
}

// NewProductMap is a constructor that creates and returns a new instance of ProductMap.
// It can be initialized with a pre-existing map of products.

func NewProductMap(db map[int]models.Product) *ProductMap {
	// Initialize with an empty map to ensure it's not nil.
	defaultDb := make(map[int]models.Product)
	if db != nil {
		// If an initial database is provided, use it.
		defaultDb = db
	}
	return &ProductMap{db: defaultDb}
}

// FindAll returns a copy of all products currently stored in the repository.
// It returns an error if the operation fails, which is nil in this implementation.
func (r *ProductMap) FindAll() (v map[int]models.Product, err error) {
	v = make(map[int]models.Product)
	// Create a new map to return a copy, preventing external modification
	// of the internal database map.
	for key, value := range r.db {
		v[key] = value
	}
	return
}

// Create adds a new product to the repository.
// It returns an error if a product with the same ID already exists.
func (r *ProductMap) Create(P models.Product) (err error) {
	// Check if a product with the given ID already exists.
	_, exists := r.db[P.ID]
	if exists {
		// Return a descriptive error if the product is a duplicate.
		err = errors.New("1")
		return
	}
	// Add the new product to the map.
	r.db[P.ID] = P
	return
}

// FindByID searches for a product by its unique ID.
// It returns the found product or an error if the product does not exist.
func (r *ProductMap) FindByID(ID int) (P models.Product, err error) {
	// Check if the product exists in the map.
	value, exists := r.db[ID]
	if !exists {
		// Return a descriptive error if the product is not found.
		err = errors.New("1")
		return
	}
	// Return the found product.
	P = value
	return
}
