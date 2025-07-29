package database

import (
	"errors"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
	"gorm.io/gorm"
)

// ProductMap implements a product repository using an in-memory map.
// The key of the map is the product ID.
type ProductBatchRepository struct {
	// db conection to db
	db *gorm.DB
}

// NewProductMap is a constructor that creates and returns a new instance of ProductMap.
// It can be initialized with a pre-existing map of products.
func NewProductBatchRepository(db *gorm.DB) *ProductBatchRepository {
	return &ProductBatchRepository{db: db}
}

// FindAll not specified in the User Story
func (r *ProductBatchRepository) FindAll() ([]models.ProductBatch, error) {
	panic("method FindAll not implemented for ProductBatchRepository")
}

// Create adds a new product to the repository.
// It returns an error if a product with the same ID already exists.
func (r *ProductBatchRepository) Create(body models.ProductBatch) (models.ProductBatch, error) {
	result := r.db.Create(&body)
	switch {
	case errors.Is(result.Error, gorm.ErrForeignKeyViolated):
		return models.ProductBatch{}, repository.ErrForeignKeyViolation
	case result.Error != nil:
		return models.ProductBatch{}, result.Error
	}
	return body, nil
}
func (r *ProductBatchRepository) Update(body models.ProductBatch) (models.ProductBatch, error) {
	panic("method Update not implemented for ProductBatchRepository")
}

// FindById not specified in the User Story
func (r *ProductBatchRepository) FindById(id int) (models.ProductBatch, error) {
	panic("method FindById not implemented for ProductBatchRepository")
}

// PartialUpdate not specified in the User Story
func (r *ProductBatchRepository) PartialUpdate(id int, fields map[string]interface{}) (models.ProductBatch, error) {
	panic("method PartialUpdate not implemented for ProductBatchRepository")
}

// Delete not specified in the User Story
func (r *ProductBatchRepository) Delete(id int) error {
	panic("method Delete not implemented for ProductBatchRepository")
}
