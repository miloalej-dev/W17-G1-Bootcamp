package database

import (
	"errors"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
	"gorm.io/gorm"
)

// ProductMap implements a product repository using an in-memory map.
// The key of the map is the product ID.
type ProductRepository struct {
	// db conection to db
	db *gorm.DB
}

// NewProductMap is a constructor that creates and returns a new instance of ProductMap.
// It can be initialized with a pre-existing map of products.
func NewProductDB(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

// FindAll returns a copy of all products currently stored in the repository.
// It returns an error if the operation fails, which is nil in this implementation.
func (r *ProductRepository) FindAll() ([]models.Product, error) {
	var products []models.Product
	// GORM genera: SELECT * FROM products;
	result := r.db.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

// Create adds a new product to the repository.
// It returns an error if a product with the same ID already exists.
func (r *ProductRepository) Create(body models.Product) (models.Product, error) {
	result := r.db.Create(&body)
	if result.Error != nil {
		return models.Product{}, errors.New(result.Error.Error())
	}
	return body, nil
}
func (r *ProductRepository) Update(body models.Product) (models.Product, error) {
	result := r.db.Save(body)
	if result.Error != nil {
		return models.Product{}, result.Error
	}
	return body, nil
}

// FindById searches for a product by its unique ID.
// It returns the found product or an error if the product does not exist.
func (r *ProductRepository) FindById(id int) (models.Product, error) {
	var product models.Product
	result := r.db.First(&product, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return models.Product{}, gorm.ErrRecordNotFound
		}
		return models.Product{}, result.Error
	}
	return product, nil
}

// PartialUpdate updates specific fields of an existing product.
func (r *ProductRepository) PartialUpdate(id int, fields map[string]interface{}) (models.Product, error) {
	var product models.Product
	// Primero, busca el producto para asegurarte de que existe.
	if err := r.db.First(&product, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Product{}, gorm.ErrRecordNotFound
		}
		return models.Product{}, err
	}

	// Aplica las actualizaciones.
	if err := r.db.Model(&product).Updates(fields).Error; err != nil {
		return models.Product{}, err
	}

	return product, nil
}

// Delete elimina un producto por su ID.
func (r *ProductRepository) Delete(id int) error {
	result := r.db.Delete(&models.Product{}, id)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
