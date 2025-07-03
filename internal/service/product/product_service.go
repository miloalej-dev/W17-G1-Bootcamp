// Package productService defines the business logic layer for product operations.
// It provides an interface that abstracts the underlying implementation details.

package productService

import "github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"

// ProductService defines the set of methods that a product service must implement.
// It acts as a contract for the business logic, decoupling the application's core
// from the transport layer (e.g., HTTP handlers) and the data access layer.

type ProductService interface {
	// FindAll returns all available products.
	// It encapsulates the business logic for retrieving the full list of products.
	// It returns a map of products keyed by their ID, or an error if the operation fails.

	FindAll() (v map[int]models.Product, err error)

	// Create handles the business logic for adding a new product.
	// This can include validation, enrichment, or other pre-processing steps
	// before passing the data to the repository. It returns an error if the
	// creation process fails.

	Create(P models.Product) (err error)
	// FindByID handles the business logic for retrieving a single product by its ID.
	// It returns the requested product or a service-specific error if the product
	// cannot be found or if another issue occurs.

	FindByID(ID int) (P models.Product, err error)
	UpdateProduct(ID int, Body models.Product) (P models.Product, err error)
	Delete(ID int) (err error)
}
