// Package productService implements the business logic for product-related operations.
// It acts as an intermediary between the transport layer (e.g., HTTP handlers) and the
// data access layer (repository), enforcing business rules.

package _default

import (
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/service"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
)

// NewProductDefault is a constructor function that creates a new instance of ProductDefault.
// It takes a ProductRepository as a dependency, promoting loose coupling and testability.

func NewProductDefault(rp repository.ProductRepository) *ProductDefault {
	return &ProductDefault{rp: rp}
}

// ProductDefault is the default concrete implementation of the product service.
type ProductDefault struct {
	// rp is the repository dependency. By using an interface, this service
	// is decoupled from the specific database implementation (e.g., in-memory, SQL).
	rp repository.ProductRepository
}

// RetrieveAll retrieves all products by calling the repository's FindAll method.
// It directly passes through the results and any error from the repository.
func (s *ProductDefault) RetrieveAll() (v []models.Product, err error) {
	return s.rp.FindAll()
}

// Register attempts to add a new product using the repository.
// If the repository returns any error, it is replaced with the generic
// errorProduct.ErrorCreate.
func (s *ProductDefault) Register(body models.Product) (models.Product, error) {
	return s.rp.Create(body)
}

// Retrieve retrieves a single product by its ID from the repository.
// If the repository returns any error (e.g., not found), it is replaced
// with the generic errorProduct.ErrorNotFound.
func (s *ProductDefault) Retrieve(id int) (models.Product, error) {
	return s.rp.FindById(id)

}
func (s *ProductDefault) Modify(body models.Product) (models.Product, error) {
	return s.rp.Update(body)
}

func (s *ProductDefault) PartialModify(id int, fields map[string]any) (models.Product, error) {
	return s.rp.PartialUpdate(id, fields)

}
func (s *ProductDefault) Remove(id int) (err error) {
	return s.rp.Delete(id)
}

func (s *ProductDefault) RetrieveRecordsCountByProductId(id int) (models.ProductReport, error) {

	if _, err := s.Retrieve(id); err == nil {
		return s.rp.FindRecordsCountByProductId(id)
	}
	return models.ProductReport{}, service.ErrProductNotFound
}
