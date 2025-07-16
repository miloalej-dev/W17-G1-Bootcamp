// Package productService implements the business logic for product-related operations.
// It acts as an intermediary between the transport layer (e.g., HTTP handlers) and the
// data access layer (repository), enforcing business rules.

package _default

import (
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
)

// NewProductDefault is a constructor function that creates a new instance of ProductDefault.
// It takes a ProductRepository as a dependency, promoting loose coupling and testability.

func NewProductBatchDefault(rp repository.ProductBatchRepository) *ProductBatchDefault {
	return &ProductBatchDefault{rp: rp}
}

// ProductDefault is the default concrete implementation of the product service.
type ProductBatchDefault struct {
	// rp is the repository dependency. By using an interface, this service
	// is decoupled from the specific database implementation (e.g., in-memory, SQL).
	rp repository.ProductBatchRepository
}

// RetrieveAll retrieves all products by calling the repository's FindAll method.
// It directly passes through the results and any error from the repository.
func (s *ProductBatchDefault) RetrieveAll() (v []models.ProductBatch, err error) {
	return s.rp.FindAll()
}

// Register attempts to add a new product using the repository.
// If the repository returns any error, it is replaced with the generic
// errorProduct.ErrorCreate.
func (s *ProductBatchDefault) Register(body models.ProductBatch) (models.ProductBatch, error) {

	// Convert hours from string to data base format TIME hours
	body.ManufacturingHour = body.ManufacturingHour * 100
	return s.rp.Create(body)
}

// Retrieve retrieves a single product by its ID from the repository.
// If the repository returns any error (e.g., not found), it is replaced
// with the generic errorProduct.ErrorNotFound.
func (s *ProductBatchDefault) Retrieve(id int) (models.ProductBatch, error) {
	return s.rp.FindById(id)

}
func (s *ProductBatchDefault) Modify(body models.ProductBatch) (models.ProductBatch, error) {
	return s.rp.Update(body)
}

func (s *ProductBatchDefault) PartialModify(id int, fields map[string]any) (models.ProductBatch, error) {
	return s.rp.PartialUpdate(id, fields)

}
func (s *ProductBatchDefault) Remove(id int) (err error) {
	return s.rp.Delete(id)
}
