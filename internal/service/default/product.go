// Package productService implements the business logic for product-related operations.
// It acts as an intermediary between the transport layer (e.g., HTTP handlers) and the
// data access layer (repository), enforcing business rules.

package _default

import (
	"errors"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
)

// ErrorCreate is returned when an attempt to create a product fails because
// a product with the same ID already exists in the repository.
var ErrorCreate = errors.New("There is already a product with that ID")

// ErrorNotFound is returned when a lookup operation fails to find a product
// with the specified ID.
var ErrorNotFound = errors.New("Product Not found with that ID")

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

// FindAll retrieves all products by calling the repository's FindAll method.
// It directly passes through the results and any error from the repository.
func (s *ProductDefault) FindAll() (v map[int]models.Product, err error) {
	v, err = s.rp.FindAll()
	if err != nil {
		err = ErrorNotFound
	}
	return
}

// Create attempts to add a new product using the repository.
// If the repository returns any error, it is replaced with the generic
// errorProduct.ErrorCreate.
func (s *ProductDefault) Create(P models.Product) (err error) {
	err = s.rp.Create(P)
	if err != nil {
		err = ErrorCreate
	}
	return
}

// FindByID retrieves a single product by its ID from the repository.
// If the repository returns any error (e.g., not found), it is replaced
// with the generic errorProduct.ErrorNotFound.
func (s *ProductDefault) FindByID(ID int) (P models.Product, err error) {
	P, err = s.rp.FindByID(ID)
	if err != nil {
		err = ErrorNotFound
	}
	return
}
func (s *ProductDefault) UpdateProduct(ID int, Body models.Product) (P models.Product, err error) {
	P, err = s.rp.UpdateProduct(ID, Body)
	if err != nil {
		err = ErrorNotFound
	}
	return
}
func (s *ProductDefault) Delete(ID int) (err error) {
	err = s.rp.Delete(ID)
	if err != nil {
		err = ErrorNotFound
	}
	return
}
