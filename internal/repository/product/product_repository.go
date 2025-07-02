// Package productRepository defines the interfaces and implementations for product data access.
package productRepository

import "github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"

// ProductRepository defines the standard operations to be performed on a product collection.
// It acts as a contract that any concrete repository (like in-memory, database, etc.)
// must implement to be used by the application's services.
type ProductRepository interface {
	// FindAll retrieves all products from the repository.
	// It returns a map of products, with the product ID as the key,
	// and an error if the operation fails.
	FindAll() (v map[int]models.Product, err error)

	// Create adds a new product to the repository.
	// It takes a Product model and returns an error if the creation fails,
	// for instance, if a product with the same ID already exists.
	Create(P models.Product) (err error)

	// FindByID looks for a single product by its unique integer ID.
	// It returns the found product or an error if no product with that ID exists.
	FindByID(ID int) (P models.Product, err error)
	UpdateProduct(ID int, Body models.Product) (P models.Product, err error)
	Delete(ID int) (err error)
}
