// Package productService defines the business logic layer for product operations.
// It provides an interface that abstracts the underlying implementation details.

package service

import "github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"

// ProductService defines the set of methods that a product service must implement.
// It acts as a contract for the business logic, decoupling the application's core
// from the transport layer (e.g., HTTP handlers) and the data access layer.

type ProductService interface {
	RetrieveAll() ([]models.Product, error)
	Retrieve(id int) (models.Product, error)
	Register(Product models.Product) (models.Product, error)
	Modify(Product models.Product) (models.Product, error)
	PartialModify(id int, fields map[string]any) (models.Product, error)
	Remove(id int) error
	RetrieveRecordsCountByProductId(id int) (models.ProductReport, error)
	RetrieveRecordsCount() ([]models.ProductReport, error)
}
